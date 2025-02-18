package main

import (
	"go-rest/db"
	"go-rest/route"
	"net/http"

	"github.com/gin-gonic/gin"
)

func test(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"hello": "world"})
}

func main() {
	db.InitDb()
	server := gin.Default()
	route.RegisterServer(server)
	server.GET("/test", test)

	server.Run(":8080")
}
