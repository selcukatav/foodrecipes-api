package routes

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"foodrecipes.com/m/v2/api/handlers"
	"foodrecipes.com/m/v2/domains"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

var (
	clientID     = ""
	clientSecret = ""
	redirectURL  = "http://localhost:3001/callback"
	oauthConf    = &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Scopes:       []string{"user:email"},
		Endpoint:     github.Endpoint,
	}
	oauthStateString = "randomstatestring"
	db               *sql.DB
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	clientID := os.Getenv("GITHUB_CLIENT_ID")
	clientSecret := os.Getenv("GITHUB_CLIENT_SECRET")

	if clientID == "" || clientSecret == "" {
		log.Fatal("GITHUB_CLIENT_ID and GITHUB_CLIENT_SECRET must be set")
	}

}

func HandleGithubLogin(c echo.Context) error {
	url := oauthConf.AuthCodeURL(oauthStateString, oauth2.AccessTypeOffline)
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func HandleGithubCallback(c echo.Context) error {
	state := c.QueryParam("state")
	if state != oauthStateString {
		log.Println("invalid oauth state")
		return c.Redirect(http.StatusTemporaryRedirect, "/")
	}

	code := c.QueryParam("code")
	token, err := oauthConf.Exchange(context.Background(), code)
	if err != nil {
		log.Printf("oauthConf.Exchange() failed with '%s'\n", err)
		return c.Redirect(http.StatusTemporaryRedirect, "/")
	}

	client := oauthConf.Client(context.Background(), token)
	resp, err := client.Get("https://api.github.com/user")
	if err != nil {
		log.Printf("client.Get() failed with '%s'\n", err)
		return c.Redirect(http.StatusTemporaryRedirect, "/")
	}
	defer resp.Body.Close()

	var user domains.Users

	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		log.Printf("json.NewDecoder() failed with '%s'\n", err)
		return c.Redirect(http.StatusTemporaryRedirect, "/")
	}
	handlerDB := handlers.NewDatabase(db)

	handlerDB.CreateUserWithoutPassword(user.Email, user.Username)

	return c.JSON(http.StatusOK, user)
}
