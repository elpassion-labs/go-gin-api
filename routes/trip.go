package routes

import (
	"github.com/gin-gonic/gin"
	"go-gin-api/db"
	"go-gin-api/models"
)

func CreateTrip(c *gin.Context) {
	var trip models.Trip
	var dbConn = db.GetDB()

	c.BindJSON(&trip)
	var result = dbConn.Create(&trip)
	if result.Error == nil {
		c.JSON(200, trip)
	} else {
		c.JSON(422, result.GetErrors())
	}
}
