// +build ignore

package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Post struct {
	Id   int
	Data string
}

func ErrorCheck(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func PingDB(db *sql.DB) {
	err := db.Ping()
	ErrorCheck(err)
}

func main() {
	db, err := sql.Open("mysql", "user:1234@/test")
	ErrorCheck(err)
	PingDB(db)

	/*
		// Create & Use database 'sample'
		// ONLY for root privilege
		var dbname = "sample"
		_, err = db.Exec("CREATE DATABASE " + dbname)
		ErrorCheck(err)

		_, err = db.Exec("USE " + dbname)
		ErrorCheck(err)
	*/

	// Create table 'posts'
	_, err = db.Exec("CREATE TABLE if not exists posts ( id integer auto_increment, data text, primary key (id) )")
	ErrorCheck(err)

	// Insert
	// prepare
	stmt, e := db.Prepare("insert into posts(data) values (?)")
	ErrorCheck(e)
	res, e := stmt.Exec("Post test")
	ErrorCheck(e)
	id, e := res.LastInsertId()
	ErrorCheck(e)
	fmt.Println("Insert id: ", id)

	defer db.Close()
}
