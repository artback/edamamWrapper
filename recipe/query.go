package recipe

import (
	"github.com/artback/edamamWrapper/edamam"
	"github.com/artback/edamamWrapper/internal/query"
	"net/url"
	"strconv"
)

const searchUrl = "https://api.edamam.com/search"

type Query struct {
	SearchText     string
	Key            string
	Id             string
	MaxIngredients int
	DietLabel      string
	HealthLabel    string
	CuisineType    []string
	MealType       string
	DishType       []string
	Calories       edamam.Range
	Time           edamam.Range
	Excluded       []string
}

func (q Query) getUrl() (string, error) {
	u, err := url.Parse(searchUrl)
	if err != nil {
		return "", err
	}
	v := query.Values{}
	v.Add("app_id", q.Id)
	v.Add("app_key", q.Key)
	v.Add("q", q.SearchText)
	v.Add("ingr", strconv.Itoa(q.MaxIngredients))
	v.Add("diet", q.DietLabel)
	v.Add("health", q.DietLabel)
	v.Add("cuisineType", q.CuisineType...)
	v.Add("dishType", q.DishType...)
	v.Add("mealType", q.MealType)
	v.Add("calories", q.Calories.String())
	v.Add("time", q.Time.String())
	v.Add("excluded", q.Excluded...)
	u.RawQuery = v.Encode()
	return u.String(), nil
}
