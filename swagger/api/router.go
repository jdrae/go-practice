package api

import (
	"test/api/user"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {

	r := gin.Default()

	r.GET("/", home)
	r.GET("/welcome/:name", welcome)

	v1 := r.Group("/v1")
	{
		user.ApplyRoutes(v1)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
