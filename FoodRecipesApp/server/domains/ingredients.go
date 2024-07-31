package domains

type RecipeIngredients struct {
	RecipeId     int    `json:"recipeId"`
	IngredientId int    `json:"ingredientId"`
	Quantity     string `json:"quantity"`
}

type Recipe struct {
	RecipeId int `json:"recipeId"`
	Title string `json:"title"`

	MatchCount int `json:"matchCount"`
}