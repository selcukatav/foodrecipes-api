package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	e.GET("/", MainPage)
	e.POST("/login", Login)
	e.POST("/register", Register)
	return e
}
func MainPage(c echo.Context) error {

	return c.String(http.StatusOK, "Welcome To Main Page")
}
