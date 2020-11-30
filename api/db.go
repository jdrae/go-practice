// +build ignore

package main

import (
	"api/model"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

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
	stmt, err := db.Prepare("insert into posts(data) values (?)")
	ErrorCheck(err)
	res, err := stmt.Exec("Post test")
	ErrorCheck(err)
	id, err := res.LastInsertId()
	ErrorCheck(err)
	fmt.Println("Insert id: ", id)

	// Query
	var post = model.Post{}
	rows, err := db.Query("select * from posts")
	ErrorCheck(err)
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&post.Id, &post.Data)
		ErrorCheck(err)
		fmt.Println(post)
	}
}
