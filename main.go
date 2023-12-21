package main

import (
	"trucode2/apitaskmanager/db"
	"trucode2/apitaskmanager/models"
	"trucode2/apitaskmanager/routers"
)

func main() {
	// establish a database connection using the DBconnection function from the "db" package
	db.DBconnetion()

	// automatically migrate data models to the database
	db.DB.AutoMigrate(&models.User{}, &models.Project{}, &models.Task{})

	// routers
	router := routers.SetupRouter()

	// specify the port on which the server will listen
	port := ":5000"

	// run the HTTP server on the specified port
	router.Run(port)

}
