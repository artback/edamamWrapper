package recipe

type Nutrient struct {
	Quantity float64 `json:"quantity"`
	Unit     string  `json:"unit"`
}
type Nutrients struct {
	Kcal    Nutrient `json:"ENERC_KCAL"`
	Protein Nutrient `json:"PROCNT"`
	Fat     Nutrient `json:"FAT"`
	Carbs   Nutrient `json:"CHOCDF"`
}
