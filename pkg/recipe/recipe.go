package recipe

import "github.com/artback/edamamWrapper/pkg/edamam"

type Recipe struct {
	Nutrients         edamam.Nutrients `json:"nutrients"`
	Category          string           `json:"category"`
	CategoryLabel     string           `json:"categoryLabel"`
	Label             string           `json:"label"`
	FoodContentsLabel string           `json:"foodContentsLabel"`
}
