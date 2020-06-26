package query

import "net/url"

type Values url.Values

func (v Values) Add(key string, values ...string) {
	uv := url.Values(v)
	for _, value := range values {
		if value != "" && value != "0" {
			uv.Add(key, value)
		}
	}
}
func (v Values) Encode() string {
	uv := url.Values(v)
	return uv.Encode()
}
