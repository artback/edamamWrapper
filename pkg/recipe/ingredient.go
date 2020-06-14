package recipe

type Ingredient struct {
	Text         string  `json:"text"`
	Quantity     float64 `json:"quantity"`
	Measure      string  `json:"measure"`
	Food         string  `json:"food"`
	Weight       float64 `json:"weight"`
	FoodCategory string  `json:"foodCategory"`
}
