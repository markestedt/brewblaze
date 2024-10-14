package recipe

import (
	"math"
	"net/http"
	"strconv"
)

type UserInput struct {
	Measurements string
	BatchSize    float64
	Style        string
	Description  string
	Yeast        string
}

func ParseRequest(r *http.Request) UserInput {
	r.ParseForm()

	batchSize, err := strconv.ParseFloat(r.FormValue("batch-size"), 64)
	if err != nil {
		batchSize = 10
	}

	return UserInput{
		Measurements: r.FormValue("measurements"),
		BatchSize:    math.Abs(batchSize),
		Style:        r.FormValue("style"),
		Description:  r.FormValue("description"),
		Yeast:        r.FormValue("yeast"),
	}
}

type Document struct {
	RecipeJson string `json:"recipe-data"`
}

type Json struct {
	Name            string  `json:"name"`
	Description     string  `json:"description"`
	BatchSize       float64 `json:"batchsize"`
	VolumeUnit      string  `json:"volumeunit"`
	OriginalGravity float64 `json:"originalgravity"`
	FinalGravity    float64 `json:"finalgravity"`
	Abv             string  `json:"abv"`
	Ibu             string  `json:"ibu"`
	Srm             string  `json:"srm"`
	Style           string  `json:"style"`
	Fermentables    []struct {
		Weight float64 `json:"weight"`
		Unit   string  `json:"unit"`
		Name   string  `json:"name"`
	} `json:"fermentables"`
	Hops []struct {
		Weight      float64 `json:"weight"`
		Unit        string  `json:"unit"`
		Name        string  `json:"name"`
		TimingValue string  `json:"timingvalue"`
		TimingUnit  string  `json:"timingunit"`
		Use         string  `json:"use"`
	} `json:"hops"`
	Yeast struct {
		Name   string `json:"name"`
		Amount string `json:"amount"`
	} `json:"yeast"`
	OtherAdditions []struct {
		Weight float64 `json:"weight"`
		Unit   string  `json:"unit"`
		Name   string  `json:"name"`
	} `json:"otheradditions"`
	Instructions []struct {
		Step int32  `json:"step"`
		Text string `json:"text"`
	} `json:"instructions"`
}
