package database

import (
	"database/sql"
	"fmt"
	model "test/model"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// Initialize database setting
func Initialize(DBCONFIG string) (*sql.DB, error) {
	db, err := sql.Open("mysql", DBCONFIG)
	if err != nil {
		panic(err)
	}
	fmt.Println("--> Connected to database")
	Migrate(db)
	return db, err
}

// Migrate needed tables for app (executed in Initialize func)
func Migrate(db *sql.DB) {
	model.CreateUserTable(db)
}

// Inject database to gin context
func Inject(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}
