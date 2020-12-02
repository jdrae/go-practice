package database

import (
	"database/sql"
	"fmt"
	model "test/model"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func Initialize() (*sql.DB, error) {
	db, err := sql.Open("mysql", "user:1234@/test")
	if err != nil {
		panic(err)
	}
	fmt.Println("--> Connected to database")
	Migrate(db)
	return db, err
}

func Migrate(db *sql.DB) {
	model.CreateUserTable(db)
}

// Inject injects database to gin context
func Inject(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}
