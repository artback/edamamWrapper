package food

import (
	"github.com/artback/edamamWrapper/internal/network"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

func TestGetOnIngredients(t *testing.T) {
	expected := Food{
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
	body, err := os.Open("../../testdata/articles_test.json")
	assert.Nil(t, err)
	response := http.Response{Body: body}
	client := network.GetMock{Response: response}
	c := Credentials{Key: "key", Id: "id"}
	articles, err := GetOnIngredients(c, client, []string{"cheese"})
	if assert.Nil(t, err) {
		assert.Equal(t, expected, articles.Articles[0])
	}
}

func TestGetOnUPC(t *testing.T) {
	expected := Food{
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
	body, err := os.Open("../../testdata/articles_test.json")
	assert.Nil(t, err)
	response := http.Response{Body: body}
	client := network.GetMock{Response: response}
	c := Credentials{Key: "key", Id: "id"}
	articles, err := GetOnUPC(c, client, "1214321")
	if assert.Nil(t, err) {
		assert.Equal(t, expected, articles.Articles[0])
	}
}
