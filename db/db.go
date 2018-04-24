package db

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/qor/validations"

	"github.com/elpassion-labs/go-gin-api/models"
)

var db *gorm.DB
var err error

// Init creates a connection to database and
// automigrates any models
func Init() {
	os.Remove(dbName())

	db, _ = gorm.Open("sqlite3", dbName())
	if err != nil {
		fmt.Println(err)
	}

	db.AutoMigrate(&models.Trip{})

	validations.RegisterCallbacks(db)
}

func GetDB() *gorm.DB {
	return db
}

func CloseDB() {
	db.Close()
}

func dbName() string {
	return fmt.Sprintf("%s.db", gin.Mode())
}
