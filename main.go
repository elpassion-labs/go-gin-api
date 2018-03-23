package main

import (
	"bikeramp-go/routes"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", routes.Ping)

	return router
}

func main() {
	router := setupRouter()
	router.Run()
}
