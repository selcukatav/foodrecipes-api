package domains

type RecipeIngredients struct {
	RecipeId     int    `json:"recipeId"`
	IngredientId int    `json:"ingredientId"`
	Quantity     string `json:"quantity"`
}
