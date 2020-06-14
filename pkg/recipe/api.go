package recipe

import (
	"fmt"
	"github.com/artback/edamamWrapper/pkg/edamam"
)

const searchUrl = "https://api.edamam.com/search"

type Credentials edamam.Credentials

func (c Credentials) GetApiURL() string {
	return fmt.Sprintf("%s?app_id=%s&app_key=%s", searchUrl, c.Id, c.Key)
}
