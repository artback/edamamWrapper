package network

import (
	"net/http"
)

type GetMock struct {
	http.Response
	Err error
}

func (m GetMock) Get(_ string) (*http.Response, error) {
	if m.Err != nil {
		return nil, m.Err
	}
	return &m.Response, nil
}
