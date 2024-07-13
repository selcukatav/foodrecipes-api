package routes

import (
	"database/sql"
	"net/http"

	"foodrecipes.com/m/v2/api/middlewares"
	"foodrecipes.com/m/v2/domains"
	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	db := c.Get("db").(*sql.DB) 

	
	var user domains.Users
	if err := c.Bind(&user); err != nil {
		return err
	}
	username := user.Username
	password := user.Password

	database := middlewares.NewDatabase(db)
	err := database.Authentication(username, password)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid username or password")
	}

	accessToken, err := middlewares.GenerateToken(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to generate access token")
	}

	refreshToken, err := middlewares.GenerateRefreshToken(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to generate refresh token")
	}

	return c.JSON(http.StatusOK, echo.Map{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}
