package food

import (
	"github.com/artback/edamamWrapper/internal/network"
	"github.com/artback/edamamWrapper/pkg/edamam"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

func TestGetArticles(t *testing.T) {
	const URL = ""
	expected := Food{
		Label:         "cheese",
		Category:      "Generic foods",
		CategoryLabel: "food",
		Nutrients: edamam.Nutrients{
			Kcal:    406,
			Protein: 24.04,
			Fat:     33.82,
			Carbs:   1.33,
		},
	}
	body, err := os.Open("../../testdata/articles_test.json")
	assert.Nil(t, err)
	response := http.Response{Body: body}
	client := network.GetMock{Response: response}
	articles, err := getFoods(URL, client)
	if assert.Nil(t, err) {
		assert.Equal(t, expected, articles.Articles[0])
	}
}
