package food

import (
	"github.com/artback/edamamWrapper/pkg/edamam"
	"net/url"
)

const parserUrl = "https://api.edamam.com/api/food-database/parser"

type Query struct {
	Key           string
	Id            string
	Upc           string
	Ingredients   []string
	Health        []string
	Calories      edamam.Range
	Category      string
	CategoryLabel string
}

func (q Query) GetURL() (string, error) {
	u, err := url.Parse(parserUrl)
	query := url.Values{}
	if q.Id != "" {
		query.Add("app_id", q.Id)
	}
	if q.Key != "" {
		query.Add("app_key", q.Key)
	}
	if q.Upc != "" {
		query.Add("upc", q.Upc)
	}
	if q.Category != "" {
		query.Add("category", q.Category)
	}
	if q.CategoryLabel != "" {
		query.Add("categoryLabel", q.CategoryLabel)
	}
	if len(q.Ingredients) > 0 {
		for _, ing := range q.Ingredients {
			query.Add("ingredients", ing)
		}
	}
	if len(q.Health) > 0 {
		for _, h := range q.Health {
			query.Add("health", h)
		}
	}
	caloriesRange := q.Calories.String()
	if caloriesRange != "" {
		query.Add("calories", caloriesRange)
	}

	u.RawQuery = query.Encode()
	return u.String(), err
}
