package main

import (
	"fmt"
	"go-rest/db"
	"go-rest/route"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func test(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"hello": "world"})
}

func main() {
	err := godotenv.Load("dev.env")
	if err != nil {
		panic("could not load env file")
	}
	port := os.Getenv("PORT")
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	dbName := os.Getenv("MYSQL_DB")
	dbHost := os.Getenv("DB_HOST")
	db.InitDb(user, password, dbHost, dbName)
	server := gin.Default()
	route.RegisterServer(server)
	server.GET("/test", test)

	server.Run(fmt.Sprintf(":%s", port))
}
