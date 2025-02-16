package main

import (
	"go-rest/db"
	"go-rest/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func test(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"hello": "world"})
}

func getEvents(context *gin.Context) {
	var events = models.GetEvents()
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	error := context.ShouldBindBodyWithJSON(&event)
	if error != nil {
		log.Fatal(error)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data"})
		return
	}
	event.ID = 0
	event.UserID = 0
	event.Save()
}
func main() {
	db.InitDb()
	server := gin.Default()
	server.GET("/test", test)
	server.GET("/events", getEvents)
	server.POST("/event", createEvent)
	server.Run(":8080")
}
