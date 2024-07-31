package routes

import (
	"database/sql"
	"fmt"
	"net/http"

	"foodrecipes.com/m/v2/api/handlers"
	"foodrecipes.com/m/v2/domains"
	"github.com/labstack/echo/v4"
)

func FilterIngredients(c echo.Context) error {
	criteriaParams := c.QueryParams()["criteria"]
	criteria := make([]int, len(criteriaParams))

	db := c.Get("db").(*sql.DB)

	for i, param := range criteriaParams {
		fmt.Sscanf(param, "%d", &criteria[i])
	}
	handlerDB := handlers.NewDatabase(db)
	recipes, err := handlerDB.FilterRecipesByIngredients(criteria)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error at filtering ingredients")

	}
	return c.JSON(http.StatusOK, recipes)
}

func FilterByCategory(c echo.Context) error {

	var recipe domains.Recipe
	if err := c.Bind(&recipe); err != nil {
		return err
	}
	recipeCategory := recipe.CategoryId

	db := c.Get("db").(*sql.DB)
	handlerDB := handlers.NewDatabase(db)
	recipes, err := handlerDB.FilterRecipesByCategory(recipeCategory)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, recipes)
}
