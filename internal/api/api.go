package api

import "fmt"

type Credentials struct {
	key string
	id  string
}

func (c Credentials) GetApiURL() string {
	return fmt.Sprintf("https://api.edamam.com/api/food-database/parser?app_id=%s&app_key=%s", c.id, c.key)
}
