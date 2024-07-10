package routes

import (
	"database/sql"
	"net/http"

	"foodrecipes.com/m/v2/api/handlers"
	"foodrecipes.com/m/v2/domains"
	"github.com/labstack/echo/v4"
)

func Register(c echo.Context) error {
	var user *domains.Users
	db := c.Get("db").(*sql.DB)

	handlerDB := handlers.NewDatabase(db)

	if err := c.Bind(&user); err != nil {
		return err
	}
	email := user.Email
	username := user.Username
	password := user.Password

	//Calls create user function
	err := handlerDB.CreateUser(email, username, password)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "There is an error at registering a user.")
	}

	return c.JSON(http.StatusOK, "User Created")
}
