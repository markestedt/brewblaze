package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
	"text/template"

	"github.com/markestedt/brewmind/internal/beerstyles"
	"github.com/markestedt/brewmind/internal/jsonhelpers"
	recipe "github.com/markestedt/brewmind/internal/recipe"
)

type recipeViewModel struct {
	Recipe     recipe.Json
	RecipeLink string
}

type indexViewModel struct {
	Styles   []beerstyles.Style
	Examples []string
}

// type ContextValues struct {
// 	m map[string]interface{}
// }

// func (v ContextValues) Get(key string) interface{} {
// 	return v.m[key]
// }

// var sessionCookieName = "brew-blaze-session"
// var contextValues = "contextValues"

// func authenticate(f http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		session, err := r.Cookie(sessionCookieName)

// 		if err != nil {
// 			http.Error(w, "Unauthorized", http.StatusUnauthorized)
// 			return
// 		}

// 		sessionClient := appwrite.NewClient(
// 			appwrite.WithProject(os.Getenv("APPWRITE_PROJECT_ID")),
// 			appwrite.WithSession(session.Value),
// 		)

// 		account := appwrite.NewAccount(sessionClient)
// 		user, err := account.Get()

// 		if err != nil {
// 			http.Error(w, "Unauthorized", http.StatusUnauthorized)
// 			return
// 		}

// 		v := ContextValues{map[string]interface{}{
// 			"sessionClient": sessionClient,
// 			"sessionUser":   user,
// 		}}

// 		ctx := context.WithValue(r.Context(), contextValues, v)
// 		f(w, r.WithContext(ctx))
// 	}
// }

func (app *application) createRecipeHandler(w http.ResponseWriter, r *http.Request) {
	input := recipe.ParseRequest(r)

	generatedRecipe, userPrompt, err := app.recipeRepository.Generate(input)
	if err != nil {
		handleError(w, "Error generating recipe", err, http.StatusInternalServerError)
		return
	}

	recipeId, err := app.recipeRepository.Create(generatedRecipe, userPrompt)
	if err != nil {
		handleError(w, "Error saving recipe to db", err, http.StatusInternalServerError)
		return
	}

	output, err := jsonhelpers.Parse[recipe.Json](generatedRecipe)
	if err != nil {
		handleError(w, "Error parsing recipe json", err, http.StatusInternalServerError)
		return
	}

	parsedTemplate, err := template.ParseFS(templates, "templates/components/recipe.html")
	if err != nil {
		handleError(w, "Error parsing recipe component template", err, http.StatusInternalServerError)
		return
	}

	recipeLink := fmt.Sprintf("%s/recipe/%s", os.Getenv("BASE_URL"), recipeId)
	parsedTemplate.Execute(w, recipeViewModel{RecipeLink: recipeLink, Recipe: output})
}

func (app *application) getIndexHandler(w http.ResponseWriter, r *http.Request) {
	beerStyles := app.beerData.Beerjson.Styles
	sort.Slice(beerStyles, func(i, j int) bool {
		// if beerStyles[i].CategoryID != beerStyles[j].CategoryID {
		// 	return beerStyles[i].CategoryID < beerStyles[j].CategoryID
		// }

		return beerStyles[i].Name < beerStyles[j].Name
	})

	parsedTemplate, err := template.ParseFS(templates, "templates/pages/index.html")
	if err != nil {
		handleError(w, "Error parsing index page template", err, http.StatusInternalServerError)
		return
	}

	examples := [4]string{
		"I want a bright, juicy beer with flavors and aromas of tropical fruit. It should have a soft and round bitterness, easy to drink.",
		"Give me a German style Helles. It should be as clean and crisp as possible.",
		"I want a single hop English IPA, using only Mosaic.",
		"I'd like a dark and rich Baltic porter with flavors of chocolate, coffee and roasted barley. Low on the sweetness.",
	}

	viewModel := indexViewModel{
		Styles:   beerStyles,
		Examples: examples[:],
	}

	parsedTemplate.Execute(w, viewModel)
}

// func (app *application) postLoginHandler(w http.ResponseWriter, r *http.Request) {
// 	r.ParseForm()

// 	email := r.FormValue("email")
// 	password := r.FormValue("password")

// 	log.Println(email)
// 	log.Println(password)

// 	account := account.New(app.appwriteClient)
// 	session, err := account.CreateEmailPasswordSession(email, password)

// 	if err != nil {
// 		log.Println(err)
// 		http.Error(w, "Could not login", http.StatusUnauthorized)
// 		return
// 	}

// 	expires, _ := time.Parse(time.RFC3339, session.Expire)
// 	cookie := http.Cookie{
// 		HttpOnly: true,
// 		Secure:   true,
// 		SameSite: http.SameSiteLaxMode,
// 		Expires:  expires,
// 		Name:     sessionCookieName,
// 		Value:    session.Secret,
// 	}

// 	w.Header().Set("HX-Redirect", "/")
// 	http.SetCookie(w, &cookie)
// }

// func (app *application) getLoginHandler(w http.ResponseWriter, r *http.Request) {

// 	parsedTemplate, err := template.ParseFS(templates, "templates/pages/login.html")
// 	if err != nil {
// 		handleError(w, "Error parsing login page template", err, http.StatusInternalServerError)
// 		return
// 	}
// 	parsedTemplate.Execute(w, nil)
// }

func (app *application) getRecipeHandler(w http.ResponseWriter, r *http.Request) {
	recipeId := r.PathValue("recipeId")
	if recipeId == "" {
		http.Error(w, "Missing required parameter: recipeId", http.StatusBadRequest)
		return
	}

	recipeJson, err := app.recipeRepository.Get(recipeId)
	if err != nil {
		handleError(w, "Error fetching recipe data", err, http.StatusInternalServerError)
		return
	}

	recipe, err := jsonhelpers.Parse[recipe.Json](recipeJson)
	if err != nil {
		handleError(w, "Error parsing recipe json", err, http.StatusInternalServerError)
		return
	}

	parsedTemplate, err := template.ParseFS(templates, "templates/pages/recipe.html")
	if err != nil {
		handleError(w, "Error parsing recipe page template", err, http.StatusInternalServerError)
		return
	}
	parsedTemplate.Execute(w, recipe)
}

func handleError(w http.ResponseWriter, message string, err error, statusCode int) {
	log.Println(message+":", err)
	http.Error(w, message, statusCode)
}

// func getLoggedInUser(ctx context.Context) *models.User {
// 	return ctx.Value(contextValues).(ContextValues).Get("sessionUser").(*models.User)
// }

// func getSessionClient(ctx context.Context) client.Client {
// 	return ctx.Value(contextValues).(ContextValues).Get("sessionClient").(client.Client)
// }
