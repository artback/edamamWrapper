package edamam

type Nutrients struct {
	Kcal    float64 `json:"ENERC_KCAL"`
	Protein float64 `json:"PROCNT"`
	Fat     float64 `json:"FAT"`
	Carbs   float64 `json:"CHOCDF"`
}
type Nutrient struct {
	Quantity float64 `json:"quantity"`
	Unit     string  `json:"unit"`
}
type TotalNutrients struct {
	Kcal    Nutrient `json:"ENERC_KCAL"`
	Protein Nutrient `json:"PROCNT"`
	Fat     Nutrient `json:"FAT"`
	Carbs   Nutrient `json:"CHOCDF"`
}
