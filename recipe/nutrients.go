package recipe

type Nutrient struct {
	Quantity float64 `json:"quantity"`
	Unit     string  `json:"unit"`
}
type Nutrients struct {
	Kcal               Nutrient `json:"ENERC_KCAL"`
	Protein            Nutrient `json:"PROCNT"`
	Fat                Nutrient `json:"FAT"`
	Carbs              Nutrient `json:"CHOCDF"`
	SaturatedFat       Nutrient `json:"FASAT"`
	MonoSaturatedFat   Nutrient `json:"FAMS"`
	PolyUnsaturatedFat Nutrient `json:"FAPU"`
	Fiber              Nutrient `json:"FIBTG"`
	Sugar              Nutrient `json:"SUGAR"`
	Cholesterol        Nutrient `json:"CHOLE"`
	Sodium             Nutrient `json:"NA"`
	Calcium            Nutrient `json:"CA"`
	Magnesium          Nutrient `json:"MG"`
	Potassium          Nutrient `json:"K"`
	Iron               Nutrient `json:"FE"`
	Zinc               Nutrient `json:"ZN"`
	Phosphorus         Nutrient `json:"P"`
	VitaminA           Nutrient `json:"VITA_RAE"`
	VitaminC           Nutrient `json:"VITC"`
	VitaminB1          Nutrient `json:"THIA"`
	VitaminB2          Nutrient `json:"RIBF"`
	VitaminB3          Nutrient `json:"NIA"`
	VitaminB6          Nutrient `json:"VITB6A"`
	VitaminB9          Nutrient `json:"FOLDFE"`
	VitaminB12         Nutrient `json:"VITB12"`
	VitaminD           Nutrient `json:"VITD"`
	VitaminE           Nutrient `json:"TOCPHA"`
	VitaminK           Nutrient `json:"VITK1"`
	Water              Nutrient `json:WATER`
}
