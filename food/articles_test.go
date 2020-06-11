package food

import (
	"github.com/artback/edamamWrapper/internal/network"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

func TestGetArticles(t *testing.T) {
	const url = ""
	body, _ := os.Open("articles_test.json")
	response := http.Response{Body: body}
	client := network.GetMock{Response: response}
	articles, _ := getArticles(url, client)
	assert.Equal(t, Articles{}, articles)
}
