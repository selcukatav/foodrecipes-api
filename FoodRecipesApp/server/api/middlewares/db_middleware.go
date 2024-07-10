package middlewares

import (
	"database/sql"

	"github.com/labstack/echo/v4"
)

func DatabaseMiddleware(db *sql.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("db", db)
			return next(c)
		}
	}
}
