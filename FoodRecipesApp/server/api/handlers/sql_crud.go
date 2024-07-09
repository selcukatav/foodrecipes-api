package handlers

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

func (d *Database) Insert(value1, value2 string) {
	query := "INSERT INTO your_table (column1, column2) VALUES (@p1, @p2)"
	_, err := d.DB.Exec(query, value1, value2)
	if err != nil {
		log.Fatal("Veri eklenemedi: ", err)
	}

	fmt.Println("Veri başarıyla eklendi!")
}

func Update() {

}

func (d *Database) Select() {
	query := "SELECT name FROM ingredients"
	rows, err := d.DB.Query(query)
	if err != nil {
		log.Fatal("Veri çekilemedi: ", err)
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
func Delete() {

}
