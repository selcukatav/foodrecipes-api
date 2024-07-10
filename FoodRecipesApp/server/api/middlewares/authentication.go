package middlewares

import (
	"database/sql"
	"fmt"
	"log"
)

type Database struct {
	DB *sql.DB
}

func NewDatabase(db *sql.DB) *Database {
	return &Database{DB: db}
}

func (d *Database) Authentication(username, password string) error {

	query := "SELECT count(*) FROM users WHERE username=@p1 AND password=@p2"

	row := d.DB.QueryRow(query, username, password)

	var count int
	err := row.Scan(&count)
	if err != nil {
		log.Fatal("Error scanning row:", err)
		return err
	}

	if count == 1 {
		fmt.Println("Logged in!")
		return nil
	} else {
		fmt.Println("Username or Password is wrong! Try Again!")
		return fmt.Errorf("authentication failed")
	}
}
