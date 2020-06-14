package recipe

import (
	"encoding/json"
	"io"
)

type body struct {
	Hits []struct {
		Recipe Recipe `json:"Recipe"`
	} `json:"hits"`
	To    int  `json:"to"`
	From  int  `json:"from"`
	More  bool `json:"more"`
	Count int  `json:"count"`
}

func (b *body) unmarshalReader(reader io.Reader) error {
	return json.NewDecoder(reader).Decode(b)
}
