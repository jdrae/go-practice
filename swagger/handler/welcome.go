package handler

import (
	model "test/model"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	c.JSON(200, gin.H{"ping": "pong"})
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
func Welcome(c *gin.Context) {
	name := c.Param("name")
	welcomeMessage := model.User{ID: 1, Name: name}

	c.JSON(200, gin.H{"message": welcomeMessage})
}
