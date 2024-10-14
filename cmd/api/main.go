package main

import (
	"embed"
	"errors"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"

	"github.com/appwrite/sdk-for-go/appwrite"
	"github.com/appwrite/sdk-for-go/client"
	"github.com/joho/godotenv"
	"github.com/markestedt/brewmind/internal/beerstyles"
	"github.com/markestedt/brewmind/internal/recipe"
	"github.com/sashabaranov/go-openai"
)

type application struct {
	beerData         *beerstyles.Json
	appwriteClient   *client.Client
	recipeRepository *recipe.Repository
}

//go:embed static
var static embed.FS

//go:embed templates
var templates embed.FS

func main() {
	LoadEnv()

	appwriteClient := appwrite.NewClient(
		appwrite.WithProject(os.Getenv("APPWRITE_PROJECT_ID")),
		appwrite.WithKey(os.Getenv("APPWRITE_API_KEY")),
	)

	aiClient := openai.NewClient(os.Getenv("AI_KEY"))

	dbService := appwrite.NewDatabases(appwriteClient)
	beerstylesRepository := beerstyles.Repository{}

	bd, err := beerstylesRepository.Get()
	if err != nil {
		log.Fatalf("Failed to load beer data: %s", err.Error())
	}

	app := &application{
		beerData:         &bd,
		appwriteClient:   &appwriteClient,
		recipeRepository: &recipe.Repository{Db: dbService, Ai: aiClient},
	}

	fSys, err := fs.Sub(static, ".")
	if err != nil {
		log.Printf("Failed to load static files: %s", err)
	}

	mux := http.NewServeMux()
	mux.Handle("GET /static/", http.FileServer(http.FS(fSys)))

	mux.HandleFunc("GET /", app.getIndexHandler)
	mux.HandleFunc("POST /create-recipe", perClientRateLimiter(app.createRecipeHandler))
	mux.HandleFunc("GET /recipe/{recipeId}", app.getRecipeHandler)

	err = http.ListenAndServe(":9595", mux)
	if errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("server closed\n")
	} else if err != nil {
		log.Fatalf("error starting server: %s\n", err)
	}
}

func LoadEnv() {
	p, err := os.Executable()

	if err != nil {
		log.Fatal(err)
	}

	p = filepath.Dir(p)
	err = godotenv.Load(path.Join(p, ".env"))

	if err != nil {
		log.Printf("Error loading .env file from %s", p)
	} else {
		return
	}

	err = godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
