package edamam

import "net/http"

type GetClient interface {
	Get(url string) (*http.Response, error)
}
