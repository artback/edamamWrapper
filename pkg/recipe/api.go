package recipe

import (
	"fmt"
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
	Calories       Range
	Time           Range
	Excluded       []string
}
type Range struct {
	To   int
	From int
}

func (r Range) String() string {
	if r.From == 0 && r.To == 0 {
		return ""
	}
	if r.To == 0 {
		return fmt.Sprintf("%d+", r.From)
	} else if r.From == 0 {
		return fmt.Sprintf("%d", r.To)
	}
	return fmt.Sprintf("%d-%d", r.From, r.To)
}

func (q Query) GetURL() (string, error) {
	u, err := url.Parse(searchUrl)
	if err != nil {
		return "", err
	}
	query := url.Values{}
	if q.Id != "" {
		query.Add("app_id", q.Id)
	}
	if q.Key != "" {
		query.Add("app_key", q.Key)
	}
	if q.SearchText != "" {
		query.Add("q", q.SearchText)
	}
	if q.MaxIngredients != 0 {
		query.Add("ingr", strconv.Itoa(q.MaxIngredients))
	}
	if q.DietLabel != "" {
		query.Add("diet", q.DietLabel)
	}
	if q.HealthLabel != "" {
		query.Add("health", q.DietLabel)
	}
	if len(q.CuisineType) > 0 {
		for _, c := range q.CuisineType {
			query.Add("cuisineType", c)
		}
	}
	if len(q.DishType) > 0 {
		for _, d := range q.DishType {
			query.Add("dishType", d)
		}
	}
	if q.MealType != "" {
		query.Add("mealType", q.MealType)
	}
	caloriesRange := q.Calories.String()
	if caloriesRange != "" {
		query.Add("calories", caloriesRange)
	}
	timeRange := q.Time.String()
	if timeRange != "" {
		query.Add("time", timeRange)
	}
	if len(q.Excluded) > 0 {
		for _, e := range q.Excluded {
			query.Add("excluded", e)
		}
	}
	u.RawQuery = query.Encode()
	return u.String(), nil
}
