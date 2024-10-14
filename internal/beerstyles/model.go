package beerstyles

type Json struct {
	Beerjson struct {
		Version float64 `json:"version"`
		Styles  []Style `json:"styles"`
	} `json:"beerjson"`
}

type Style struct {
	Name                string `json:"name"`
	Category            string `json:"category"`
	CategoryID          string `json:"category_id"`
	StyleID             string `json:"style_id"`
	CategoryDescription string `json:"category_description,omitempty"`
	OverallImpression   string `json:"overall_impression,omitempty"`
	Aroma               string `json:"aroma,omitempty"`
	Appearance          string `json:"appearance,omitempty"`
	Flavor              string `json:"flavor,omitempty"`
	Mouthfeel           string `json:"mouthfeel,omitempty"`
	Comments            string `json:"comments,omitempty"`
	History             string `json:"history,omitempty"`
	StyleComparison     string `json:"style_comparison,omitempty"`
	Tags                string `json:"tags,omitempty"`
	OriginalGravity     struct {
		Minimum struct {
			Unit  string  `json:"unit"`
			Value float64 `json:"value"`
		} `json:"minimum"`
		Maximum struct {
			Unit  string  `json:"unit"`
			Value float64 `json:"value"`
		} `json:"maximum"`
	} `json:"original_gravity,omitempty"`
	InternationalBitternessUnits struct {
		Minimum struct {
			Unit  string `json:"unit"`
			Value int    `json:"value"`
		} `json:"minimum"`
		Maximum struct {
			Unit  string `json:"unit"`
			Value int    `json:"value"`
		} `json:"maximum"`
	} `json:"international_bitterness_units,omitempty"`
	FinalGravity struct {
		Minimum struct {
			Unit  string  `json:"unit"`
			Value float64 `json:"value"`
		} `json:"minimum"`
		Maximum struct {
			Unit  string  `json:"unit"`
			Value float64 `json:"value"`
		} `json:"maximum"`
	} `json:"final_gravity,omitempty"`
	AlcoholByVolume struct {
		Minimum struct {
			Unit  string  `json:"unit"`
			Value float64 `json:"value"`
		} `json:"minimum"`
		Maximum struct {
			Unit  string  `json:"unit"`
			Value float64 `json:"value"`
		} `json:"maximum"`
	} `json:"alcohol_by_volume,omitempty"`
	Color struct {
		Minimum struct {
			Unit  string  `json:"unit"`
			Value float64 `json:"value"`
		} `json:"minimum"`
		Maximum struct {
			Unit  string  `json:"unit"`
			Value float64 `json:"value"`
		} `json:"maximum"`
	} `json:"color,omitempty"`
	Ingredients             string `json:"ingredients,omitempty"`
	Examples                string `json:"examples,omitempty"`
	StyleGuide              string `json:"style_guide"`
	Type                    string `json:"type"`
	EntryInstructions       string `json:"entry_instructions,omitempty"`
	Notes                   string `json:"notes,omitempty"`
	CurrentlyDefinedTypes   string `json:"currently_defined_types,omitempty"`
	StrengthClassifications string `json:"strength_classifications,omitempty"`
	Accordingly             string `json:"accordingly,omitempty"`
	VitalStatistics         string `json:"vital_statistics,omitempty"`
	ImpresionGeneral        string `json:"impresion_general,omitempty"`
	Aspecto                 string `json:"aspecto,omitempty"`
	Sabor                   string `json:"sabor,omitempty"`
	SensacionEnBoca         string `json:"sensacion_en_boca,omitempty"`
	Comentarios             string `json:"comentarios,omitempty"`
	Historia                string `json:"historia,omitempty"`
	Ingredientes            string `json:"ingredientes,omitempty"`
	EjemplosComerciales     string `json:"ejemplos_comerciales,omitempty"`
	ImpressaoGeral          string `json:"impressao_geral,omitempty"`
	Aparencia               string `json:"aparencia,omitempty"`
	SensacaoDeBoca          string `json:"sensacao_de_boca,omitempty"`
	ComparacoesDeEstilo     string `json:"comparacoes_de_estilo,omitempty"`
	ExemplosComerciais      string `json:"exemplos_comerciais,omitempty"`
	Marcacoes               string `json:"marcacoes,omitempty"`
}
