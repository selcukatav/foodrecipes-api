package routes

import (
	"database/sql"
	"fmt"
	"net/http"

	"foodrecipes.com/m/v2/api/handlers"
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
	recipes, err := handlerDB.FilterRecipes(criteria)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error at filtering ingredients")

	}
	return c.JSON(http.StatusOK, recipes)
}
