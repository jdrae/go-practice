package main

import (
	"test/docs"
	"test/model"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func setupRouter() *gin.Engine {

	r := gin.Default()

	r.GET("/welcome/:name", welcomePathParam)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}

// Welcome godoc
// @Summary Summary를 적어 줍니다.
// @Description 자세한 설명은 이곳에 적습니다.
// @name get-string-by-int
// @Accept  json
// @Produce  json
// @Param name path string true "User name"
// @Router /welcome/{name} [get]
// @Success 200 {object} model.User
func welcomePathParam(c *gin.Context) {
	name := c.Param("name")
	welcomeMessage := model.User{ID: 1, Name: name}

	c.JSON(200, gin.H{"message": welcomeMessage})
}

func main() {

	// programatically set swagger info
	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server for Swagger."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = ""

	r := setupRouter()

	r.Run() // Listen and Server in 0.0.0.0:8080
}
