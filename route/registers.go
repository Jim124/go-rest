package route

import (
	"go-rest/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId := context.GetInt64("eventId")
	event, error := models.GetEventById(eventId)
	if error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch event"})
		return
	}
	error = event.Register(userId)
	if error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not create register "})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "created register successfully"})
}

func cancelRegister(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId := context.GetInt64("eventId")
	var event models.Event
	event.ID = eventId
	error := event.Cancel(userId)
	if error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not delete register"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "cancelled"})
}
