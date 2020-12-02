package model

import (
	"database/sql"
	"fmt"
)

type User struct {
	Id   int    `json:"id" example:"1"`
	Name string `json:"name" example:"user name"`
}

// takes db as parameter not by gin context
// because it's executed before creating a context
func CreateUserTable(db *sql.DB) {
	_, err := db.Exec("CREATE TABLE if not exists user ( id integer auto_increment, name text, primary key (id) )")
	if err != nil {
		panic(err)
	}
	fmt.Println("--> Created user table")
}
