package server

import (
	"log"

	"foodrecipes.com/m/v2/api/routes"
	"foodrecipes.com/m/v2/database"
)

func Run() {
	db, err := database.ConnectMssql()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Handlers paketi için veritabanı bağlantısını oluştur
	//handlerDB := handlers.NewDatabase(db.DB)

	e := routes.New()
	e.Logger.Fatal(e.Start(":3001"))

}
