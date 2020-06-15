package food

import (
	"github.com/artback/edamamWrapper/internal/network"
	"github.com/artback/edamamWrapper/pkg/edamam"
)

type Food struct {
	Nutrients         Nutrients `json:"nutrients"`
	Category          string    `json:"category"`
	CategoryLabel     string    `json:"categoryLabel"`
	Label             string    `json:"label"`
	FoodContentsLabel string    `json:"foodContentsLabel"`
}
type Response struct {
	Articles []Food
	Next     func() (*Response, error)
}

func GetFoods(query Query, client network.GetClient) (*Response, error) {
	q, err := query.GetURL()
	if err != nil {
		return nil, edamam.InternalError{Err: err}
	}
	return getFoods(q, client)
}
