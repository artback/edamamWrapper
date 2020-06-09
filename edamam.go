package edamam

import (
	"fmt"
	"github.com/artback/edamamWrapper/internal/api"
	"github.com/artback/edamamWrapper/internal/read"
	"io"
	"net/http"
)

func GetOnUPC(c api.Credentials, reader io.Reader) (*http.Response, error) {
	url := c.GetApiURL()
	upc, err := read.ToString(reader)
	if err != nil {
		return nil, err
	}
	url += fmt.Sprintf("&upc=%s", *upc)
	return http.Get(url)
}
func GetOnIngredients(c api.Credentials, readers ...io.Reader) (*http.Response, error) {
	url := c.GetApiURL()
	for _, reader := range readers {
		ing, err := read.ToString(reader)
		if err != nil {
			return nil, err
		}
		url += fmt.Sprintf("&=%s", *ing)
	}
	return http.Get(url)
}
