package food

import (
	"net/url"
)

const parserUrl = "https://api.edamam.com/api/food-database/parser"

type Query struct {
	Upc         string
	Ingredients []string
	Key         string
	Id          string
}

func (q Query) GetURL() (string, error) {
	u, err := url.Parse(parserUrl)
	query := url.Values{
		"app_id":      []string{q.Id},
		"app_key":     []string{q.Key},
		"upc":         []string{q.Upc},
		"ingredients": q.Ingredients,
	}
	u.RawQuery = query.Encode()
	return u.String(), err
}
