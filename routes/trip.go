package routes

import (
	"github.com/elpassion-labs/go-gin-api/db"
	"github.com/elpassion-labs/go-gin-api/models"
	"github.com/gin-gonic/gin"
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
