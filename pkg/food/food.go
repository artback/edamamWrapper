package food

import (
	"fmt"
	"github.com/artback/edamamWrapper/internal/network"
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

func GetOnUPC(c Credentials, client network.GetClient, upc string) (*Response, error) {
	url := c.GetURL()
	url = fmt.Sprintf("&upc=%s", upc)
	return GetFoods(url, client)
}
func GetOnIngredients(c Credentials, client network.GetClient, ingredients []string) (*Response, error) {
	url := c.GetURL()
	for _, ing := range ingredients {
		url += fmt.Sprintf("&=%s", ing)
	}
	return GetFoods(url, client)
}
