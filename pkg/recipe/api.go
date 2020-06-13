package recipe

import (
	"fmt"
	"github.com/artback/edamamWrapper/pkg/edamam"
)

const url = "https://api.edamam.com/search"

type Credentials edamam.Credentials

func (c Credentials) GetApiURL() string {
	return fmt.Sprintf("%s?app_id=%s&app_key=%s", url, c.Id, c.Key)
}
