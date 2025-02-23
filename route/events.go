package route

import (
	"go-rest/models"
	"go-rest/utils"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, error := models.GetEvents()
	if error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not query event"})
		return
	}
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	bearerToken := context.Request.Header.Get("Authorization")

	if bearerToken == "" {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}
	token := strings.Split(bearerToken, " ")[1]
	userId, error := utils.ValidToken(token)
	if error != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}
	var event models.Event
	error = context.ShouldBindBodyWithJSON(&event)

	if error != nil {
		log.Fatal(error)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data"})
		return
	}
	event.UserID = userId
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

func updateEvent(context *gin.Context) {
	id, error := strconv.ParseInt(context.Param("id"), 10, 64)
	if error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse value"})
		return
	}
	_, error = models.GetEventById(id)
	if error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch event"})
		return
	}
	var event models.Event
	error = context.ShouldBindBodyWithJSON(&event)
	if error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event"})
		return
	}
	event.ID = id
	error = event.UpdateEvent()
	if error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not update event"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "update event successfully"})
}

func deleteEvent(context *gin.Context) {
	id, error := strconv.ParseInt(context.Param("id"), 10, 64)
	if error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse the id"})
		return
	}
	event, error := models.GetEventById(id)
	if error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch event"})
		return
	}
	error = event.DeleteEvent()
	if error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not delete event"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "delete event successfully"})

}
