package recipe

import (
	"fmt"
	"math"
	"net/http"
	"os"
	"strconv"
)

type UserInput struct {
	Measurements string
	BatchSize    float64
	Style        string
	Description  string
	Yeast        string
	UseSrm       bool
	Srm          float64
}

func ParseRequest(r *http.Request) UserInput {
	r.ParseForm()

	batchSize, err := strconv.ParseFloat(r.FormValue("batch-size"), 64)
	if err != nil {
		batchSize = 10
	}

	useSrm, err := strconv.ParseBool(r.FormValue("useSrm"))
	if err != nil {
		useSrm = false
	}

	srm, err := strconv.ParseFloat(r.FormValue("srm"), 64)
	if err != nil {
		srm = 20
	}

	return UserInput{
		Measurements: r.FormValue("measurements"),
		BatchSize:    math.Abs(batchSize),
		Style:        r.FormValue("style"),
		Description:  r.FormValue("description"),
		Yeast:        r.FormValue("yeast"),
		UseSrm:       useSrm,
		Srm:          srm,
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

func GetLink(recipeId string) string {
	return fmt.Sprintf("%s/recipe/%s", os.Getenv("BASE_URL"), recipeId)
}
