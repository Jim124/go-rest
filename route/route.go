package route

import "github.com/gin-gonic/gin"

func RegisterServer(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEventById)
	server.GET("/queryEvent", getEventQuery)
	server.POST("/event", createEvent)
	server.PUT("/events/:id", updateEvent)
}
