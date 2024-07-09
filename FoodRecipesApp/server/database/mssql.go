package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

type Database struct {
	DB *sql.DB
}

func ConnectMssql() (*Database, error) {
	var err error
	db := &Database{}

	// Conn string
	connString := "server=DESKTOP-HVQPCRP;integrated security=true;database=FoodRecipes"
	// SQL connection
	db.DB, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Couldn't connect to the db: ", err)
		return nil, err
	}

	err = db.DB.Ping()
	if err != nil {
		log.Fatal("DB connection couldn't be verified: ", err)
		return nil, err
	}

	fmt.Println("Connected to DB!!!")

	return db, nil
}

func (db *Database) Close() error {
	return db.DB.Close()
}