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
	e.GET("/githubLogin", HandleGithubLogin)
	e.GET("/auth/callback", HandleGithubCallback)
	e.GET("/recipes", FilterIngredients)
	
	return e
}
func MainPage(c echo.Context) error {
	html := `<html><body><a href="/githubLogin">GitHub Login</a></body></html>`
	return c.HTML(http.StatusOK, html)
}
