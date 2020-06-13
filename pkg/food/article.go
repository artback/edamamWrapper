package food

import (
	"github.com/artback/edamamWrapper/internal/network"
	"github.com/artback/edamamWrapper/pkg/edamam"
	"io"
	"net/http"
)

type Food struct {
	Nutrients         edamam.Nutrients `json:"nutrients"`
	Category          string           `json:"category"`
	CategoryLabel     string           `json:"categoryLabel"`
	Label             string           `json:"label"`
	FoodContentsLabel string           `json:"foodContentsLabel"`
}
type Foods struct {
	Articles []Food
	Next     func() (*Foods, error)
}

func getFoods(url string, client network.GetClient) (*Foods, error) {
	resp, err := client.Get(url)
	if err != nil {
		return nil, edamam.HttpError{Err: err}
	}
	var f Foods
	if err := f.unmarshalReader(resp.Body); err != nil {
		return &f, edamam.InternalError{Err: err}
	}
	return &f, nil
}
func (f *Foods) unmarshalReader(response io.Reader) error {
	var body body
	if err := body.unmarshalReader(response); err != nil {
		return edamam.InternalError{Err: err}
	}
	for _, h := range body.Hints {
		f.Articles = append(f.Articles, h.Food)
	}
	for _, p := range body.Parsed {
		f.Articles = append(f.Articles, p.Food)
	}
	if body.Links.Next != nil {
		f.Next = createNext(body.Links.Next.Url)
	}
	return nil
}

func createNext(url string) func() (*Foods, error) {
	return func() (*Foods, error) {
		return getFoods(url, http.DefaultClient)
	}
}
