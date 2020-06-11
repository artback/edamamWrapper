package food

import (
	"github.com/artback/edamamWrapper/internal/network"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

func TestGetArticles(t *testing.T) {
	const URL = ""
	expected := Article{
		Label:         "cheese",
		Category:      "Generic foods",
		CategoryLabel: "food",
		Nutrients: Nutrients{
			Kcal:    406,
			Protein: 24.04,
			Fat:     33.82,
			Carbs:   1.33,
		},
	}
	body, _ := os.Open("articles_test.json")
	response := http.Response{Body: body}
	client := network.GetMock{Response: response}
	articles, _ := getArticles(URL, client)
	assert.Equal(t, expected, articles.Articles[0])
}
