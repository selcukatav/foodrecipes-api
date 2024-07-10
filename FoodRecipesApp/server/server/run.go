package server

import (
	"log"

	"foodrecipes.com/m/v2/api/middlewares"
	"foodrecipes.com/m/v2/api/routes"
	"foodrecipes.com/m/v2/database"
)

func Run() {
	db, err := database.ConnectMssql()

	e := routes.New()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	e.Use(middlewares.DatabaseMiddleware(db.DB))
	
	//handlerDB := handlers.NewDatabase(db.DB)

	// handlerDB.Select()

	e.Logger.Fatal(e.Start(":3001"))

}
