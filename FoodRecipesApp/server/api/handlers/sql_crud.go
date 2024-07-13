package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strings"

	"foodrecipes.com/m/v2/domains"
	"github.com/labstack/echo/v4"
)

type Database struct {
	DB *sql.DB
}

func NewDatabase(db *sql.DB) *Database {
	return &Database{DB: db}
}

func Update() {

}

// db test query
func (d *Database) Select() {
	query := "SELECT name FROM ingredients"
	rows, err := d.DB.Query(query)
	if err != nil {
		log.Fatal("Data couldn't add: ", err)
	}
	defer rows.Close()

	for rows.Next() {
		var name string

		err := rows.Scan(&name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("names: %s", name)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}

// normal signup
func (d *Database) CreateUser(email, username, password string) error {
	var count int
	query := "SELECT COUNT(*) FROM users WHERE email = @p1 OR username = @p2"
	err := d.DB.QueryRow(query, email, username).Scan(&count)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error checking existing email or username")

	}

	if count > 0 {
		return echo.NewHTTPError(http.StatusInternalServerError, "User exists!")
	}

	queryInsert := "INSERT INTO users (email, username, password) VALUES (@p1,@p2,@p3)"
	_, err = d.DB.Exec(queryInsert, email, username, password)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error creating user")
	}

	return nil
}

// signing up with oauth2, provides signup without any password
func (d *Database) CreateUserWithoutPassword(email, username string) error {
	var count int
	query := "SELECT COUNT(*) FROM oauth_users WHERE email = @p1 OR username = @p2"
	err := d.DB.QueryRow(query, email, username).Scan(&count)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error checking existing email or username")
	}

	if count > 0 {
		return echo.NewHTTPError(http.StatusInternalServerError, "User exists!")
	}

	queryInsert := "INSERT INTO oauth_users (email, username) VALUES (@p1,@p2)"
	_, err = d.DB.Exec(queryInsert, email, username)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error creating user")
	}

	return nil
}

// FilterRecipes function to filter recipes based on ingredient criteria
// TODO: Should grouped by recipe_ids
func (d *Database) FilterRecipes(criteria []int) ([]domains.RecipeIngredients, error) {
	query := `SELECT recipe_id, ingredient_id, quantity FROM recipe_ingredients WHERE `

	whereClauses := []string{}
	for _, criterion := range criteria {
		whereClauses = append(whereClauses, fmt.Sprintf("ingredient_id = %d", criterion))
	}
	query += strings.Join(whereClauses, " OR ")
	query += " ORDER BY "
	for i, criterion := range criteria {
		if i != 0 {
			query += " + "
		}
		query += fmt.Sprintf("CHARINDEX('%d', CAST(ingredient_id AS VARCHAR))", criterion)
	}
	query += " DESC"

	rows, err := d.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error while executing query: %w", err)
	}
	defer rows.Close()

	var recipes []domains.RecipeIngredients
	for rows.Next() {
		var recipe domains.RecipeIngredients
		if err := rows.Scan(&recipe.RecipeId, &recipe.IngredientId, &recipe.Quantity); err != nil {
			return nil, fmt.Errorf("error while scanning row: %w", err)
		}
		recipes = append(recipes, recipe)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error after scanning rows: %w", err)
	}

	return recipes, nil
}
