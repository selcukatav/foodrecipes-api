package middlewares

import (
	"net/http"
	"time"

	"foodrecipes.com/m/v2/domains"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

var jwtKey = []byte("foodrecipe-secret-key")

func GenerateToken(user *domains.Users) (string, error) {
	claims := jwt.MapClaims{
		"sub":      1,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 8760).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	mobileAccessToken, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}
	return mobileAccessToken, nil

}

func GenerateRefreshToken(user *domains.Users) (string, error) {
	claims := jwt.MapClaims{
		"sub":      1,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	refreshToken, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}
	return refreshToken, nil

}
func SetCookie(c echo.Context, name, value string, expires time.Time) {
	cookie := &http.Cookie{
		Name:    name,
		Value:   value,
		Expires: expires,
		Path:    "/",
	}
	c.SetCookie(cookie)
}
