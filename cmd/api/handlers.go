package main

import (
	"log"
	"net/http"
	"sort"

	"github.com/markestedt/brewblaze/internal/beerstyles"
	recipe "github.com/markestedt/brewblaze/internal/recipe"
	"github.com/markestedt/brewblaze/internal/tools"
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

	viewModel := recipeViewModel{RecipeLink: recipe.GetLink(recipeId), Recipe: output}
	app.templates.ExecuteTemplate(w, "components/recipe", viewModel)

}

func (app *application) getIndexHandler(w http.ResponseWriter, r *http.Request) {
	beerStyles := app.beerData.Beerjson.Styles
	sort.Slice(beerStyles, func(i, j int) bool {
		return beerStyles[i].Name < beerStyles[j].Name
	})

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

	app.templates.ExecuteTemplate(w, "pages/index", viewModel)
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

	recipeData, err := tools.ParseJson[recipe.Json](recipeJson)
	if err != nil {
		handleError(w, "Error parsing recipe json", err, http.StatusInternalServerError)
		return
	}

	viewModel := recipeViewModel{RecipeLink: recipe.GetLink(recipeId), Recipe: recipeData}
	app.templates.ExecuteTemplate(w, "pages/recipe", viewModel)
}

func handleError(w http.ResponseWriter, message string, err error, statusCode int) {
	log.Println(message+":", err)
	http.Error(w, message, statusCode)
}
