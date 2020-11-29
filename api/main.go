package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func home(c *gin.Context) {

}

func main() {
	db, err := sql.Open("mysql", "user:1234@/test")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	var txt string
	err = db.QueryRow("SELECT txt FROM text WHERE id = 1").Scan(&txt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(txt)

	r := gin.Default()
	r.GET("/", home)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// localhost:5000
	r.Run(":5000")
}
