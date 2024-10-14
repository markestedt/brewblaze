package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
	"text/template"

	"github.com/markestedt/brewmind/internal/beerstyles"
	recipe "github.com/markestedt/brewmind/internal/recipe"
	"github.com/markestedt/brewmind/internal/tools"
)

type recipeViewModel struct {
	Recipe     recipe.Json
	RecipeLink string
}

type indexViewModel struct {
	Styles   []beerstyles.Style
	Examples []string
}

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

	output, err := tools.ParseJson[recipe.Json](generatedRecipe)
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

	recipe, err := tools.ParseJson[recipe.Json](recipeJson)
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
