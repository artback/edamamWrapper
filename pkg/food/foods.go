package food

import (
	"github.com/artback/edamamWrapper/internal/network"
	"github.com/artback/edamamWrapper/pkg/edamam"
	"io"
	"net/http"
)

func GetFoods(url string, client network.GetClient) (*Response, error) {
	resp, err := client.Get(url)
	if err != nil {
		return nil, edamam.HttpError{Err: err}
	}
	var f Response
	if err := f.unmarshalReader(resp.Body); err != nil {
		return &f, edamam.InternalError{Err: err}
	}
	return &f, nil
}
func (f *Response) unmarshalReader(response io.Reader) error {
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

func createNext(url string) func() (*Response, error) {
	return func() (*Response, error) {
		return GetFoods(url, http.DefaultClient)
	}
}
