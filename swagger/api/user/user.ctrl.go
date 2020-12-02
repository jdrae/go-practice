package user

import (
	"database/sql"
	"fmt"
	model "test/model"

	"github.com/gin-gonic/gin"
)

type User = model.User

func add(c *gin.Context) {
	// db := c.MustGet("db").(*sql.DB)
}

func getAll(c *gin.Context) {
	fmt.Println("--> controller worked ")

	db := c.MustGet("db").(*sql.DB)
	rows, err := db.Query("SELECT * FROM user")
	if err != nil {
		c.AbortWithStatus(500)
		return
	}

	var users []User
	for rows.Next() {
		var user User
		rows.Scan(&user.Id, &user.Name)
		users = append(users, user)
	}
	defer rows.Close()
	c.JSON(200, users)
}

func get(c *gin.Context) {
	// db := c.MustGet("db").(*sql.DB)

}
