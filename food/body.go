package food

import (
	"encoding/json"
	"github.com/artback/edamamWrapper/internal/edamam"
	"io"
)

type body struct {
	Hints []struct {
		Food Food `json:"food"`
	} `json:"hints"`
	Parsed []struct {
		Food Food `json:"food"`
	} `json:"parsed"`
	Links edamam.Links `json:"_links"`
}

func (b *body) unmarshalReader(reader io.Reader) error {
	return json.NewDecoder(reader).Decode(b)
}
