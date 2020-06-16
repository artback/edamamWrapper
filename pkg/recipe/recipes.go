package recipe

import (
	"github.com/artback/edamamWrapper/internal/network"
	"github.com/artback/edamamWrapper/pkg/edamam"
)

func GetRecipes(query Query, client network.GetClient) (*Response, error) {
	url, err := query.GetURL()
	if err != nil {
		return nil, edamam.InternalError{Err: err}
	}
	return getRecipes(url, client)
}
