package food

import (
	"github.com/artback/edamamWrapper/internal/network"
	"github.com/artback/edamamWrapper/pkg/edamam"
	"io"
	"net/http"
)

type Article struct {
	Nutrients         Nutrients `json:"nutrients"`
	Category          string    `json:"category"`
	CategoryLabel     string    `json:"categoryLabel"`
	Label             string    `json:"Label"`
	FoodContentsLabel string    `json:"foodContentsLabel"`
}
type Links struct {
	Next *struct {
		url string
	} `json:"next"`
}
type Articles struct {
	Articles []Article
	Next     func() (*Articles, error)
}

type Nutrients struct {
	Kcal    float64 `json:"ENERC_KCAL"`
	Protein float64 `json:"PROCNT"`
	Fat     float64 `json:"FAT"`
	Carbs   float64 `json:"CHOCDF"`
}

func getArticles(url string, client network.GetClient) (*Articles, error) {
	resp, err := client.Get(url)
	if err != nil {
		return nil, edamam.HttpError{Err: err}
	}
	var articles Articles
	if err := articles.unmarshalReader(resp.Body); err != nil {
		return &articles, edamam.InternalError{Err: err}
	}
	return &articles, nil
}
func (articles Articles) unmarshalReader(response io.Reader) error {
	var body body
	if err := body.unmarshalReader(response); err != nil {
		return edamam.InternalError{Err: err}
	}
	for _, h := range body.Hints {
		articles.Articles = append(articles.Articles, h.Food)
	}
	for _, p := range body.Parsed {
		articles.Articles = append(articles.Articles, p.Food)
	}
	if body.Links.Next != nil {
		articles.Next = next(body.Links.Next.url)
	}
	return nil
}

func next(url string) func() (*Articles, error) {
	return func() (*Articles, error) {
		return getArticles(url, http.DefaultClient)
	}
}
