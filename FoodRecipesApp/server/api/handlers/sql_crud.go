package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Database struct {
	DB *sql.DB
}

func NewDatabase(db *sql.DB) *Database {
	return &Database{DB: db}
}

func (d *Database) Insert(value1, value2 string) {
	query := "INSERT INTO your_table (column1, column2) VALUES (@p1, @p2)"
	_, err := d.DB.Exec(query, value1, value2)
	if err != nil {
		log.Fatal("Data couldn't add: ", err)
	}

	fmt.Println("Data added successfully!!")
}

func Update() {

}

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

func (d *Database) CreateUser(email, username, password string) error {
    var count int
    query := "SELECT COUNT(*) FROM users WHERE email = @p1 OR username = @p2"
    err := d.DB.QueryRow(query, email, username).Scan(&count)
    if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError,"Error checking existing email or username")

    }

    if count > 0 {
        return echo.NewHTTPError(http.StatusInternalServerError,"User exists!")
    }

    queryInsert := "INSERT INTO users (email, username, password) VALUES (@p1,@p2,@p3)"
    _, err = d.DB.Exec(queryInsert, email, username, password)
    if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError,"Error creating user")
    }

    return nil
}



