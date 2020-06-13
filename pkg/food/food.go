package food

import (
	"fmt"
	"github.com/artback/edamamWrapper/internal/network"
)

func GetOnUPC(c Credentials, client network.GetClient, upc string) (*Foods, error) {
	url := c.GetURL()
	url = fmt.Sprintf("&upc=%s", upc)
	return getFoods(url, client)
}
func GetOnIngredients(c Credentials, client network.GetClient, ingredients []string) (*Foods, error) {
	url := c.GetURL()
	for _, ing := range ingredients {
		url += fmt.Sprintf("&=%s", ing)
	}
	return getFoods(url, client)
}
