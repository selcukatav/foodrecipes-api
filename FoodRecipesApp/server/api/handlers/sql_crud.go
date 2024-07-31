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
// TODO: Doesn't work atm, needs to be fixed
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
func (d *Database) FilterRecipesByIngredients(criteria []int) ([]domains.Recipe, error) {
	if len(criteria) == 0 {
		return nil, fmt.Errorf("criteria is empty")
	}

	query := `SELECT r.recipe_id, r.title, COUNT(ri.ingredient_id) AS match_count 
			  FROM recipes r 
			  JOIN recipe_ingredients ri ON r.recipe_id = ri.recipe_id 
			  WHERE ri.ingredient_id IN (%s) 
			  GROUP BY r.recipe_id, r.title 
			  ORDER BY match_count DESC`

	criteriaStr := make([]string, len(criteria))
	for i, criterion := range criteria {
		criteriaStr[i] = fmt.Sprintf("%d", criterion)
	}
	query = fmt.Sprintf(query, strings.Join(criteriaStr, ","))

	rows, err := d.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error while executing query: %w", err)
	}
	defer rows.Close()

	var recipes []domains.Recipe
	for rows.Next() {
		var recipe domains.Recipe
		if err := rows.Scan(&recipe.RecipeId, &recipe.Title, &recipe.MatchCount); err != nil {
			return nil, fmt.Errorf("error while scanning row: %w", err)
		}
		recipes = append(recipes, recipe)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error after scanning rows: %w", err)
	}

	return recipes, nil
}

// Gets recipes by category
func (d *Database) FilterRecipesByCategory(categoryId int) ([]domains.Recipe, error) {

	query := `SELECT recipe_id, title, category_id FROM recipes WHERE category_id = $1 ORDER BY title DESC`
	rows, err := d.DB.Query(query, categoryId)
	if err != nil {
		return nil, fmt.Errorf("Error while getting recipes by category", err)
	}
	defer rows.Close()

	var recipes []domains.Recipe

	for rows.Next() {
		var recipe domains.Recipe
		if err := rows.Scan(&recipe.RecipeId, &recipe.Title, &recipe.CategoryId); err != nil {
			return nil, fmt.Errorf("Error while getting recipes by category", err)
		}
		recipes = append(recipes, recipe)
	}

	// Satırları döngü içinde işlerken oluşabilecek hataları kontrol edelim
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("Error while getting recipes by category", err)
	}

	return recipes, nil
}
