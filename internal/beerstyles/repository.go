package beerstyles

import (
	"embed"
	"encoding/json"
)

type Repository struct {
}

//go:embed bjcp_styleguide-2021.json
var f embed.FS

func (r Repository) Get() (Json, error) {
	jsonData, _ := f.ReadFile("bjcp_styleguide-2021.json")
	beerData := Json{}

	if err := json.Unmarshal(jsonData, &beerData); err != nil {
		return Json{}, err
	}
	return beerData, nil
}
