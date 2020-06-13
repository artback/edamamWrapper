package recipe

import (
	"encoding/json"
	"github.com/artback/edamamWrapper/internal/edamam"
	"io"
)

type body struct {
	Hints []struct {
		Recipe Recipe `json:"Recipe"`
	} `json:"hints"`
	Parsed []struct {
		Recipe Recipe `json:"Recipe"`
	} `json:"parsed"`
	Links edamam.Links `json:"_links"`
}

func (b *body) unmarshalReader(reader io.Reader) error {
	return json.NewDecoder(reader).Decode(b)
}
