package db

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/qor/validations"

	"go-gin-api/models"
)

var db *gorm.DB
var err error

// Init creates a connection to database and
// automigrates any models
func Init() {
	db, _ = gorm.Open("sqlite3", fmt.Sprintf("%s.db", gin.Mode()))
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
