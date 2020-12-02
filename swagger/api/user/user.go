package user

import (
	"github.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.RouterGroup) {
	user := r.Group("/user")
	{
		user.POST("/", add)
		user.GET("/", getAll)
		user.GET("/:id", get)
	}
}
