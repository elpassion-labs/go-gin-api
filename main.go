package main

import (
	"github.com/gin-gonic/gin"
	"go-gin-api/db"
	"go-gin-api/routes"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	db.Init()

	router := routes.SetupRouter()
	router.Run()

	defer db.CloseDB()
}
