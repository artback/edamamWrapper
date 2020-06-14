package recipe

import (
	"github.com/artback/edamamWrapper/internal/network"
	"github.com/artback/edamamWrapper/pkg/edamam"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

func TestGet(t *testing.T) {
	const URL = ""
	expected := Recipe{
		Label:        "cheese",
		Source:       "Not Without Salt",
		Url:          "http://notwithoutsalt.com/strong-cheese/",
		DietLabels:   []string{"Low-Carb"},
		HealthLabels: []string{"Sugar-Conscious", "Low Potassium", "Keto-Friendly", "Vegetarian", "Pescatarian", "Gluten-Free", "Wheat-Free", "Egg-Free", "Peanut-Free", "Tree-Nut-Free", "Soy-Free", "Fish-Free", "Shellfish-Free", "Pork-Free", "Red-Meat-Free", "Crustacean-Free", "Celery-Free", "Mustard-Free", "Sesame-Free", "Lupine-Free", "Mollusk-Free", "Kosher"},
		Cautions:     []string{"Sulfites"},
		Portions:     8,
		Nutrients: edamam.TotalNutrients{
			Kcal:    edamam.Nutrient{},
			Protein: edamam.Nutrient{},
			Fat:     edamam.Nutrient{},
			Carbs:   edamam.Nutrient{},
		},
	}
	body, err := os.Open("../../testdata/articles_test.json")
	assert.Nil(t, err)
	response := http.Response{Body: body}
	client := network.GetMock{Response: response}
	r, err := getRecipes(URL, client)
	if assert.Nil(t, err) {
		assert.Equal(t, expected, r.Recipes[0])
	}
}
