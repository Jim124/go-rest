package main

import (
	"go-rest/db"
	"go-rest/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func test(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"hello": "world"})
}

func getEvents(context *gin.Context) {
	events, error := models.GetEvents()
	if error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not query event"})
		return
	}
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
	event.UserID = 1
	error = event.Save()
	if error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event, Try again later"})
		return
	}
	context.JSON(http.StatusOK, event)

}

func getEventById(context *gin.Context) {
	id, error := strconv.ParseInt(context.Param("id"), 10, 64)
	if error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not get id from url"})
		return
	}
	event, error := models.GetEventById(id)
	if error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not get event by id"})
		return
	}
	context.JSON(http.StatusOK, event)

}

func getEventQuery(context *gin.Context) {
	id := context.Query("id")
	context.JSON(http.StatusOK, gin.H{"message": id})
}
func main() {
	db.InitDb()
	server := gin.Default()
	server.GET("/test", test)
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEventById)
	server.GET("/queryEvent", getEventQuery)
	server.POST("/event", createEvent)
	server.Run(":8080")
}
