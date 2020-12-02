package user

import (
	"database/sql"
	model "test/model"

	"github.com/gin-gonic/gin"
)

type User = model.User

func errorCheck(err error) {
	if err != nil {
		panic(err)
		return
	}
}

// curl -d "{""name"":""value1""}" -H "Content-Type: application/json" -X POST http://localhost:5000/v1/user/
func add(c *gin.Context) {
	var u User
	err := c.Bind(&u)
	errorCheck(err)
	db := c.MustGet("db").(*sql.DB)
	stmt, err := db.Prepare("INSERT INTO user(name) VALUES (?)")
	errorCheck(err)
	res, err := stmt.Exec(u.Name)
	errorCheck(err)
	id, err := res.LastInsertId()
	errorCheck(err)
	defer stmt.Close()
	c.JSON(200, int(id))
}

func getAll(c *gin.Context) {
	db := c.MustGet("db").(*sql.DB)
	rows, err := db.Query("SELECT * FROM user")
	errorCheck(err)

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
	id := c.Param("id")
	db := c.MustGet("db").(*sql.DB)
	row := db.QueryRow("SELECT * FROM user WHERE id=?", id)
	var u User
	err := row.Scan(&u.Id, &u.Name)
	errorCheck(err)
	c.JSON(200, u)
}
