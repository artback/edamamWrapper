package food

import (
	"fmt"
	"github.com/artback/edamamWrapper/internal/network"
	"github.com/artback/edamamWrapper/pkg/edamam"
	"net/http"
)

func GetOnUPC(c edamam.Credentials, upc string) (*Articles, error) {
	return getOnUPC(c, http.DefaultClient, upc)
}
func GetOnIngredients(c edamam.Credentials, ingredients []string) (*Articles, error) {
	return getOnIngredients(c, http.DefaultClient, ingredients)
}

func getOnUPC(c edamam.Credentials, client network.GetClient, upc string) (*Articles, error) {
	url := c.GetApiURL()
	url = fmt.Sprintf("&upc=%s", upc)
	return getArticles(url, client)
}
func getOnIngredients(c edamam.Credentials, client network.GetClient, ingredients []string) (*Articles, error) {
	url := c.GetApiURL()
	for _, ing := range ingredients {
		url += fmt.Sprintf("&=%s", ing)
	}
	return getArticles(url, client)
}
