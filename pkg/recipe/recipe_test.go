package recipe

import (
	"github.com/artback/edamamWrapper/internal/network"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

func TestGetRecipes(t *testing.T) {
	const URL = ""
	expected := Recipe{
		Label:        "Strong Cheese",
		Source:       "Not Without Salt",
		Url:          "http://notwithoutsalt.com/strong-cheese/",
		Calories:     2201.3315221999997,
		TotalWeight:  564.8423700000001,
		TotalTime:    0.0,
		CuisineType:  []string{"french"},
		MealType:     []string{"snack"},
		DishType:     []string{"condiments and sauces"},
		DietLabels:   []string{"Low-Carb"},
		HealthLabels: []string{"Sugar-Conscious", "Low Potassium", "Keto-Friendly", "Vegetarian", "Pescatarian", "Gluten-Free", "Wheat-Free", "Egg-Free", "Peanut-Free", "Tree-Nut-Free", "Soy-Free", "Fish-Free", "Shellfish-Free", "Pork-Free", "Red-Meat-Free", "Crustacean-Free", "Celery-Free", "Mustard-Free", "Sesame-Free", "Lupine-Free", "Mollusk-Free", "Kosher"},
		Cautions:     []string{"Sulfites"},
		Portions:     8,
		Nutrients: Nutrients{
			Kcal: Nutrient{
				Quantity: 2201.3315221999997,
				Unit:     "kcal",
			},
			Protein: Nutrient{
				Quantity: 109.815685748,
				Unit:     "g",
			},
			Fat: Nutrient{
				Quantity: 188.02908953399998,
				Unit:     "g",
			},
			Carbs: Nutrient{
				Quantity: 8.812068521,
				Unit:     "g",
			},
		},
		TotalDaily: Nutrients{
			Kcal: Nutrient{
				Quantity: 110.06657610999999,
				Unit:     "%",
			},
			Protein: Nutrient{
				Quantity: 219.63137149600004,
				Unit:     "%",
			},
			Fat: Nutrient{
				Quantity: 289.27552235999997,
				Unit:     "%",
			},
			Carbs: Nutrient{
				Quantity: 2.937356173666667,
				Unit:     "%",
			},
		},
		Ingredients: []Ingredient{
			{
				Text:         "1 lb left-over cheese, at room temperature",
				Quantity:     1.0,
				Measure:      "pound",
				Food:         "cheese",
				Weight:       453.59237,
				FoodCategory: "Cheese",
			},
			{
				Text:         "1/4 cup dry white wine",
				Quantity:     0.25,
				Measure:      "cup",
				Food:         "dry white wine",
				Weight:       58.8,
				FoodCategory: "wines",
			},
			{
				Text:         "3 tbsp unsalted butter, softened",
				Quantity:     3.0,
				Measure:      "tablespoon",
				Food:         "unsalted butter",
				Weight:       42.599999999999994,
				FoodCategory: "Dairy",
			},
			{
				Text:         "2 tbsp fresh parsley leaves",
				Quantity:     2.0,
				Measure:      "tablespoon",
				Food:         "parsley",
				Weight:       7.6,
				FoodCategory: "vegetables",
			},
			{
				Text:         "1 small clove garlic",
				Quantity:     1.0,
				Measure:      "clove",
				Food:         "garlic",
				Weight:       2.25,
				FoodCategory: "vegetables",
			},
		},
	}
	file, err := os.Open("../../testdata/recipes_test.json")
	assert.Nil(t, err)
	response := http.Response{Body: file}
	client := network.GetMock{Response: response}
	r, err := getRecipes(URL, client)
	if assert.Nil(t, err) {
		assert.Equal(t, expected, r.Recipes[0])
	}
}
