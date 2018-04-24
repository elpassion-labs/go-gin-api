package main

import (
	"github.com/elpassion-labs/go-gin-api/db"
	"github.com/elpassion-labs/go-gin-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	db.Init()

	router := routes.SetupRouter()
	router.Run()

	defer db.CloseDB()
}
