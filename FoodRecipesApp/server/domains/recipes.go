package domains

type RecipeIngredients struct {
	RecipeId     int    `json:"recipeId"`
	IngredientId int    `json:"ingredientId"`
	Quantity     string `json:"quantity"`
}

type Recipe struct {
	RecipeId int `json:"recipeId"`
	Title string `json:"title"`
	CategoryId int `json:"categoryId"`
	MatchCount int `json:"matchCount"`	
}

type RecipeByTime struct {
	RecipeId int `json:"recipeId"`
	Title string `json:"title"`
	TimeType string `json:"timeType"`
	PrepTime int `json:"prepTime"`
	CookTime int `json:"cookTime"`
	Time1 int `json:"time1"`
	Time2 int `json:"time2"`
}

