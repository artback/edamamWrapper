# Golang wrapper around the Edamam API
[EDAMAM API](https://developer.edamam.com/)

## There are 3 different apis in the edamam suite therefor this project will consist of 3 packages.
They all have the same structure of api key and api id but returns different types.

TODO:
The nutrition package is not yet implemented.
The Nutrition list is not fully implemented only consist of a few properties.
Should have a way of getting a list of available categories,cautions,meal types etc.
Add optional parameters support according to EDAMAM API spec. 

### Credentials

First you need a credentials object for the api consisting of a key and id you get from Edamam
``` 
credentials := food.Credentials{Key: "121321fda","1fsafsaf"}

```

### Food
Get the api url containing the credentials
```
func (c Credentials) GetURL() string 
```

Get food articles based on upc number
```
func GetOnUPC(c Credentials, client network.GetClient, upc string) (*Response, error) 
```

Get food articles based on list of ingredients
```
func GetOnIngredients(c Credentials, client network.GetClient, ingredients []string) (*Response, error)
```

Get food articles from url. Can be used in  conjunction with GetURL() for example to send custom query not yet supported by wrapper
```
func GetFoods(url string, client network.GetClient) (*Response, error) 
```

Response contains of a list of food articles and a function which resolves the next page of result if there is a next page url in the api response:
```

type Response struct {
  Articles []Food
    Next     func() (*Response, error)
}
type Food struct {
  Nutrients         Nutrients `json:"nutrients"`
    Category          string    `json:"category"`
    CategoryLabel     string    `json:"categoryLabel"`
    Label             string    `json:"label"`
    FoodContentsLabel string    `json:"foodContentsLabel"`
}
// To be expanded to support all the Nutrients in the edamam api
type Nutrients struct {
  Kcal    float64 `json:"ENERC_KCAL"`
    Protein float64 `json:"PROCNT"`
    Fat     float64 `json:"FAT"`
    Carbs   float64 `json:"CHOCDF"`
}

```



### Recipes
Get the api url containing the credentials
```
func (c Credentials) GetURL() string 
```


Get recipes from url. Can be used in  conjunction with GetURL() for example to send custom query not yet supported by wrapper.
```
func GetFoods(url string, client network.GetClient) (*Response, error) 
```

Response contains of a list of recipes and a function which resolves the next page of result if there is a next page url in the api response:
```
type Response struct {
  Recipes []Recipe
    Next    func() (*Response, error)
}
type Recipe struct {
	Source            string                `json:"source"`
	Url               string                `json:"url"`
	Portions          float64               `json:"yield"`
	DietLabels        []string              `json:"dietLabels"`
	HealthLabels      []string              `json:"healthLabels"`
	Cautions          []string              `json:"cautions"`
	Ingredients       []Ingredient          `json:"ingredients"`
	Label             string                `json:"label"`
	FoodContentsLabel string                `json:"foodContentsLabel"`
	Calories          float64               `json:"calories"`
	TotalWeight       float64               `json:"totalWeight"`
	TotalTime         float64               `json:"totalTime"`
	CuisineType       []string              `json:"cuisineType"`
	MealType          []string              `json:"mealType"`
	DishType          []string              `json:"dishType"`
	Nutrients         Nutrients `json:"totalNutrients"`
	TotalDaily        Nutrients `json:"totalDaily"`
}
type Nutrient struct {
	Quantity float64 `json:"quantity"`
	Unit     string  `json:"unit"`
}
// To be expanded to support all the Nutrients in the edamam api
type Nutrients struct {
	Kcal    Nutrient `json:"ENERC_KCAL"`
	Protein Nutrient `json:"PROCNT"`
	Fat     Nutrient `json:"FAT"`
	Carbs   Nutrient `json:"CHOCDF"`
}

type Ingredient struct {
  Text         string  `json:"text"`
    Quantity     float64 `json:"quantity"`
    Measure      string  `json:"measure"`
    Food         string  `json:"food"`
    Weight       float64 `json:"weight"`
    FoodCategory string  `json:"foodCategory"`
}
```