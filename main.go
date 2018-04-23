package main

import (
	"github.com/gin-gonic/gin"
	"go-gin-api/db"
	"go-gin-api/routes"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", routes.Ping)
	router.POST("/api/trips", routes.CreateTrip)

	return router
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	db.Init()

	router := setupRouter()
	router.Run()

	defer db.CloseDB()
}
