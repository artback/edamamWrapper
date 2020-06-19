package recipe

import (
	"fmt"
	"github.com/artback/edamamWrapper/edamam"
	"github.com/artback/edamamWrapper/internal/network"
	"io"
	"net/http"
)

type Recipe struct {
	Source            string       `json:"source"`
	Url               string       `json:"url"`
	Portions          float64      `json:"yield"`
	DietLabels        []string     `json:"dietLabels"`
	HealthLabels      []string     `json:"healthLabels"`
	Cautions          []string     `json:"cautions"`
	Ingredients       []Ingredient `json:"ingredients"`
	Label             string       `json:"label"`
	FoodContentsLabel string       `json:"foodContentsLabel"`
	Calories          float64      `json:"calories"`
	TotalWeight       float64      `json:"totalWeight"`
	TotalTime         float64      `json:"totalTime"`
	CuisineType       []string     `json:"cuisineType"`
	MealType          []string     `json:"mealType"`
	DishType          []string     `json:"dishType"`
	Nutrients         Nutrients    `json:"totalNutrients"`
	TotalDaily        Nutrients    `json:"totalDaily"`
}

type Response struct {
	Recipes []Recipe
	Next    func() (*Response, error)
}

func GetRecipes(query Query, client network.GetClient) (*Response, error) {
	url, err := query.getUrl()
	if err != nil {
		return nil, edamam.InternalError{Err: err}
	}
	return getRecipes(url, client)
}
func getRecipes(url string, client network.GetClient) (*Response, error) {
	resp, err := client.Get(url)
	if err != nil {
		return nil, edamam.HttpError{Err: err}
	}
	var r Response
	if err := r.unmarshalReader(resp.Body, url); err != nil {
		return &r, edamam.InternalError{Err: err}
	}
	return &r, nil
}
func (r *Response) unmarshalReader(reader io.Reader, url string) error {
	var body body
	if err := body.unmarshalReader(reader); err != nil {
		return edamam.InternalError{Err: err}
	}
	for _, h := range body.Hits {
		r.Recipes = append(r.Recipes, h.Recipe)
	}
	if body.More == true {
		from := body.To
		to := body.To + (body.To - body.From)
		url := fmt.Sprintf("%s&from=%d&to=%d", url, from, to)
		r.Next = createNext(url)
	}
	return nil
}

func createNext(url string) func() (*Response, error) {
	return func() (*Response, error) {
		return getRecipes(url, http.DefaultClient)
	}
}
