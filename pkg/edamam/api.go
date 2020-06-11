package edamam

import (
	"fmt"
)

const url = "https://api.edamam.com/api/food-database/parser"

type Credentials struct {
	key string
	id  string
}

func (c Credentials) GetApiURL() string {
	return fmt.Sprintf("%s?app_id=%s&app_key=%s", url, c.id, c.key)
}
