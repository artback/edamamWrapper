# Golang wrapper around the Edamam API
[EDAMAM API](https://developer.edamam.com/)

## There are 3 different apis in the edamam suite therefor this project will consist of 3 packages.
They all have the same structure of api key and api id but returns different types.

TODO:
The nutrition analyse package is not implemented.
The Nutrition list is not fully implemented only consist of a few properties.
--Should have a way of getting a list of available categories,cautions,meal types etc.

### Query

First you need a Query object for the api consisting of a key and id you get from Edamam and
the additional query parameters
``` 
query := food.Query{Key: "121321fda","1fsafsaf",}
query := food.Query{Key: "121321fda","1fsafsaf"}

```

### Food

The structure of the query parameters:
```go
type Query struct {
	Key           string
	Id            string
	Upc           string
	Ingredients   []string
	Health        []string
	Calories      edamam.Range
	Category      string
	CategoryLabel string
}
type Range struct {
	To   int
	From int
}
```
Get food articles limited by query
```go
func GetFoods(query Query, client network.GetClient) (*Response, error)
```

Response contains of a list of food articles and a function which resolves the next page of result if there is a next page url in the api response:
```go
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
```go
type Query struct {
	SearchText     string
	Key            string
	Id             string
	MaxIngredients int
	DietLabel      string
	HealthLabel    string
	CuisineType    []string
	MealType       string
	DishType       []string
	Calories       edamam.Range
	Time           edamam.Range
	Excluded       []string
}
type Range struct {
	To   int
	From int
}
```
Get recipes on limited by query.
```go
func GetRecipes(query Query, client network.GetClient) (*Response, error)
```

Response contains of a list of recipes and a function which resolves the next page of result if there is a next page url in the api response:
```go
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
