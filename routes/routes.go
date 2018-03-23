package routes

import (
	"bikeramp-go/db"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	fmt.Printf("%v\n", db.Connection)
	c.String(200, "pong")
}
