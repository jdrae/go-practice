package main

import (
	"fmt"
	"os"
	api "test/api"
	database "test/database"
	docs "test/docs"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// load .env variables
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	PORT := os.Getenv("PORT")
	DBCONFIG := os.Getenv("DBCONFIG")

	// set swagger info
	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server for Swagger."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:" + PORT
	docs.SwaggerInfo.BasePath = ""

	// set router and db
	app := gin.Default()

	// initialize & set db
	db, err := database.Initialize(DBCONFIG)
	if err != nil {
		panic(err)
	}
	app.Use(database.Inject(db))

	// add routes
	api.ApplyRoutes(app)

	// run app
	fmt.Println("--> Server is listening on http://localhost:" + PORT)
	fmt.Println("--> Swagger docs created in http://localhost:" + PORT + "/swagger/index.html")
	app.Run(":" + PORT)
}
