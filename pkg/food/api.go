package food

import (
	"fmt"
	"github.com/artback/edamamWrapper/pkg/edamam"
)

const parserUrl = "https://api.edamam.com/api/food-database/parser"

type Credentials edamam.Credentials

func (c Credentials) GetURL() string {
	return fmt.Sprintf("%s?app_id=%s&app_key=%s", parserUrl, c.Id, c.Key)
}
