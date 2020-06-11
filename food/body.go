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

func (b *body) Unmarshal(data []byte) {
	json.Unmarshal(data, b)
}
func (b *body) UnmarshalReader(reader io.Reader) {
	json.NewDecoder(reader).Decode(b)
}
