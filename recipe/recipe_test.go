package recipe

import (
	"github.com/artback/edamamWrapper/internal/test"
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
			SaturatedFat: Nutrient{
				Quantity: 109.7465727216,
				Unit:     "g",
			},
			MonoSaturatedFat: Nutrient{
				Quantity: 47.2063784436,
				Unit:     "g",
			},
			PolyUnsaturatedFat: Nutrient{
				Quantity: 7.811323162100002,
				Unit:     "g",
			},
			Fiber: Nutrient{
				Quantity: 0.29805,
				Unit:     "g",
			},
			Sugar: Nutrient{
				Quantity: 1.9471986360000004,
				Unit:     "g",
			},
			Cholesterol: Nutrient{
				Quantity: 554.2542174,
				Unit:     "mg",
			},
			Sodium: Nutrient{
				Quantity: 2933.3993628000007,
				Unit:     "mg",
			},
			Calcium: Nutrient{
				Quantity: 3091.8249975000003,
				Unit:     "mg",
			},
			Magnesium: Nutrient{
				Quantity: 133.56443990000002,
				Unit:     "mg",
			},
			Potassium: Nutrient{
				Quantity: 447.8287012,
				Unit:     "mg",
			},
			Iron: Nutrient{
				Quantity: 1.402477792,
				Unit:     "mg",
			},
			Zinc: Nutrient{
				Quantity: 15.774538291000002,
				Unit:     "mg",
			},
			Phosphorus: Nutrient{
				Quantity: 2174.1504101,
				Unit:     "mg",
			},
			VitaminA: Nutrient{
				Quantity: 1516.3279331000003,
				Unit:     "µg",
			},
			VitaminC: Nutrient{
				Quantity: 10.81,
				Unit:     "mg",
			},
			VitaminB1: Nutrient{
				Quantity: 0.1385759399,
				Unit:     "mg",
			},
			VitaminB2: Nutrient{
				Quantity: 2.0018178858,
				Unit:     "mg",
			},
			VitaminB3: Nutrient{
				Quantity: 0.3738350243,
				Unit:     "mg",
			},
			VitaminB6: Nutrient{
				Quantity: 0.28756576130000006,
				Unit:     "mg",
			},
			VitaminB9: Nutrient{
				Quantity: 131.4195162,
				Unit:     "µg",
			},
			VitaminB12: Nutrient{
				Quantity: 4.064032856000001,
				Unit:     "µg",
			},
			VitaminD: Nutrient{
				Quantity: 3.36055422,
				Unit:     "µg",
			},
			VitaminE: Nutrient{
				Quantity: 4.585140486000001,
				Unit:     "mg",
			},
			VitaminK: Nutrient{
				Quantity: 141.04962873,
				Unit:     "µg",
			},
			Water: Nutrient{
				Quantity: 234.98289927000002,
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
			SaturatedFat: Nutrient{
				Quantity: 548.732863608,
				Unit:     "%",
			},
			Fiber: Nutrient{
				Quantity: 1.1922,
				Unit:     "%",
			},
			Cholesterol: Nutrient{
				Quantity: 184.75140580000001,
				Unit:     "%",
			},
			Sodium: Nutrient{
				Quantity: 122.22497345000004,
				Unit:     "%",
			},
			Calcium: Nutrient{
				Quantity: 309.18249975000003,
				Unit:     "%",
			},
			Magnesium: Nutrient{
				Quantity: 31.801057119047627,
				Unit:     "%",
			},
			Potassium: Nutrient{
				Quantity: 9.528270238297873,
				Unit:     "%",
			},
			Iron: Nutrient{
				Quantity: 7.791543288888889,
				Unit:     "%",
			},
			Zinc: Nutrient{
				Quantity: 143.40489355454548,
				Unit:     "%",
			},
			Phosphorus: Nutrient{
				Quantity: 310.5929157285714,
				Unit:     "%",
			},
			VitaminA: Nutrient{
				Quantity: 168.4808814555556,
				Unit:     "%",
			},
			VitaminC: Nutrient{
				Quantity: 12.011111111111111,
				Unit:     "%",
			},
			VitaminB1: Nutrient{
				Quantity: 11.547994991666666,
				Unit:     "%",
			},
			VitaminB2: Nutrient{
				Quantity: 153.98599121538462,
				Unit:     "%",
			},
			VitaminB3: Nutrient{
				Quantity: 2.336468901875,
				Unit:     "%",
			},
			VitaminB6: Nutrient{
				Quantity: 22.120443176923082,
				Unit:     "%",
			},
			VitaminB9: Nutrient{
				Quantity: 32.85487905,
				Unit:     "%",
			},
			VitaminB12: Nutrient{
				Quantity: 169.33470233333335,
				Unit:     "%",
			},
			VitaminD: Nutrient{
				Quantity: 22.4036948,
				Unit:     "%",
			},
			VitaminE: Nutrient{
				Quantity: 30.567603240000008,
				Unit:     "%",
			},
			VitaminK: Nutrient{
				Quantity: 117.54135727500001,
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
	file, err := os.Open("../testdata/recipes_test.json")
	assert.Nil(t, err)
	response := http.Response{Body: file}
	client := test.GetMock{Response: response}
	r, err := getRecipes(URL, client)
	if assert.Nil(t, err) {
		assert.Equal(t, expected, r.Recipes[0])
	}
}
