package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-gin-api/db"
	"github.com/go-gin-api/routes"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	db.Init()

	router := routes.SetupRouter()
	router.Run()

	defer db.CloseDB()
}
