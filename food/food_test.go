package food

import (
	"github.com/artback/edamamWrapper/internal/test"
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
	body, err := os.Open("../testdata/articles_test.json")
	assert.Nil(t, err)
	response := http.Response{Body: body}
	client := test.GetMock{Response: response}
	c := Query{Key: "key", Id: "id", Ingredients: []string{"cheese", "carrot"}}
	articles, err := GetFoods(c, client)
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
	body, err := os.Open("../testdata/articles_test.json")
	assert.Nil(t, err)
	response := http.Response{Body: body}
	client := test.GetMock{Response: response}
	c := Query{Key: "key", Id: "id", Upc: "1231212141231"}
	articles, err := GetFoods(c, client)
	if assert.Nil(t, err) {
		assert.Equal(t, expected, articles.Articles[0])
	}
}

func TestGetArticles(t *testing.T) {
	const URL = ""
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
	body, err := os.Open("../testdata/articles_test.json")
	assert.Nil(t, err)
	response := http.Response{Body: body}
	client := test.GetMock{Response: response}
	articles, err := getFoods(URL, client)
	if assert.Nil(t, err) {
		assert.Equal(t, expected, articles.Articles[0])
	}
}
