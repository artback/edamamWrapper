package food

import (
	"github.com/artback/edamamWrapper/edamam"
	"github.com/artback/edamamWrapper/internal/query"
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

func (q Query) getURL() (string, error) {
	u, err := url.Parse(parserUrl)
	v := query.Values{}
	v.Add("app_id", q.Id)
	v.Add("app_key", q.Key)
	v.Add("upc", q.Upc)
	v.Add("category", q.Category)
	v.Add("categoryLabel", q.CategoryLabel)
	v.Add("ingredients", q.Ingredients...)
	v.Add("health", q.Health...)
	v.Add("calories", q.Calories.String())
	u.RawQuery = v.Encode()
	return u.String(), err
}
