package Structs

type Coffee struct {
	Roasts      []string `json:"Roasts"`
	Modifiers   []string `json:"Modifiers"`
	Types       []string `json:"Types"`
	TypesPlural []string `json:"TypesPlural"`
	Condiments  []string `json:"Condiments"`
	Sizes       []string `json:"Sizes"`
	SizesPlural []string `json:"SizesPlural"`
	Quantities  []string `json:"Quantities"`
}
