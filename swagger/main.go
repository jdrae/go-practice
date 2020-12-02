package main

import (
	"fmt"
	database "test/database"
	docs "test/docs"
	router "test/router"
)

func main() {
	var PORT = "5000"

	// programatically set swagger info
	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server for Swagger."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:" + PORT
	docs.SwaggerInfo.BasePath = ""

	app := router.Router()

	db, _ := database.Initialize()
	app.Use(database.Inject(db))

	fmt.Println("--> Server is listening on http://localhost:" + PORT)
	fmt.Println("--> Swagger docs created in http://localhost:" + PORT + "/swagger/index.html")

	app.Run(":" + PORT)
}
