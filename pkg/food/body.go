package food

import (
	"encoding/json"
	"io"
)

type body struct {
	Hints []struct {
		Food Article `json:"food"`
	} `json:"hints"`
	Parsed []struct {
		Food Article `json:"food"`
	} `json:"parsed"`
	Links Links `json:"_links"`
}

func (b *body) unmarshalReader(reader io.Reader) error {
	return json.NewDecoder(reader).Decode(b)
}
