package food

import (
	"fmt"
	"github.com/artback/edamamWrapper/internal/network"
)

func GetOnUPC(c Credentials, client network.GetClient, upc string) (*Articles, error) {
	url := c.GetApiURL()
	url = fmt.Sprintf("&upc=%s", upc)
	return getArticles(url, client)
}
func GetOnIngredients(c Credentials, client network.GetClient, ingredients []string) (*Articles, error) {
	url := c.GetApiURL()
	for _, ing := range ingredients {
		url += fmt.Sprintf("&=%s", ing)
	}
	return getArticles(url, client)
}
