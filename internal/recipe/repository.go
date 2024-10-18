package recipe

import (
	"context"
	"fmt"
	"os"

	"github.com/appwrite/sdk-for-go/databases"
	"github.com/appwrite/sdk-for-go/id"
	"github.com/sashabaranov/go-openai"
	"github.com/sashabaranov/go-openai/jsonschema"
)

type Repository struct {
	Db *databases.Databases
	Ai *openai.Client
}

func (r Repository) Get(recipeId string) (string, error) {
	docResponse, err := r.Db.GetDocument(
		os.Getenv("APPWRITE_DB_ID"),
		os.Getenv("APPWRITE_RECIPE_COL_ID"),
		recipeId,
	)

	if err != nil {
		return "", err
	}

	var recipeDocumentData Document
	err = docResponse.Decode(&recipeDocumentData)

	if err != nil {
		return "", err
	}

	return recipeDocumentData.RecipeJson, nil
}

func (r Repository) Create(recipe string, userPrompt string) (string, error) {
	doc, err := r.Db.CreateDocument(
		os.Getenv("APPWRITE_DB_ID"),
		os.Getenv("APPWRITE_RECIPE_COL_ID"),
		id.Unique(),
		map[string]interface{}{
			"recipe-data": recipe,
			"user-prompt": userPrompt,
		},
	)
	return doc.Id, err
}

func (r Repository) Generate(input UserInput) (string, string, error) {
	var output Json
	schema, err := jsonschema.GenerateSchemaForType(output)

	if err != nil {
		return "", "", err
	}

	systemPrompt := `You are BrewBlaze, an AI-powered expert homebrewing assistant with extensive knowledge of brewing techniques, 
	beer styles, ingredients, and recipe development. Your goal is to help users create accurate and customized beer recipes tailored to 
	their preferences and brewing equipment. Always provide detailed and clear instructions, breaking down the brewing process step-by-step. 
	When generating recipes, ensure you accurately list all key ingredients,		
	including fermentables (malts, sugars), hops, yeast, and any other additions (e.g., spices, fruit, adjuncts). 
	Each ingredient should be clearly specified, with quantities appropriate for the given batch size. 
	Include instructions for the use of each ingredient, detailing when and how they should be added 
	during the brewing process (e.g., mash, boil, fermentation, dry hopping). Dont forget the other additions (e.g., spices, fruit, adjuncts) in the instructions.
	Always specify brewing parameters like target ABV, IBU, SRM, mash temperature, boil time, and fermentation conditions. 
	Emphasize safety and accuracy, ensuring all ingredient measurements, brewing steps, and fermentation guidelines are suitable for homebrewers. 
	Respond with clear, structured recipes, including sections for Ingredients, Instructions, and any special notes or tips. 
	Be open to refining the recipe based on user feedback, suggest improvements if applicable, and avoid complex jargon unless specifically requested by the user.
	Always respect the units of measurements selected by the user.
	Awlays respect the users selected type of yeast, if any.
	Use a grist-to-water ratio of 2.6 to 3.2 liters of water per kilogram of grist, unless specifically told to deviate from that ratio.
	Always give the beer a funny name.`

	volumeUnit := "litres"

	if input.Measurements == "imperial" {
		volumeUnit = "gallons"
	}

	userPrompt := fmt.Sprintf("Generate a recipe for %f %s of %s with a tasting profile that matches this desciption: %s. All measurements should be using the %s system.", input.BatchSize, volumeUnit, input.Style, input.Description, input.Measurements)

	if input.Yeast != "any" {
		userPrompt = fmt.Sprintf("%s Only suggest recipes using %s yeast.", userPrompt, input.Yeast)
	}

	if input.UseSrm {
		userPrompt = fmt.Sprintf("%s Try to hit a target SRM value of %f", userPrompt, input.Srm)
	}

	resp, err := r.Ai.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4o20240806,
			ResponseFormat: &openai.ChatCompletionResponseFormat{
				Type: openai.ChatCompletionResponseFormatTypeJSONSchema,
				JSONSchema: &openai.ChatCompletionResponseFormatJSONSchema{
					Name:   "recipe",
					Schema: schema,
					Strict: true,
				},
			},
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: systemPrompt,
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: userPrompt,
				},
			},
		},
	)

	if err != nil {
		return "", "", err
	}

	return resp.Choices[0].Message.Content, userPrompt, nil
}
