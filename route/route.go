package route

import (
	"go-rest/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterServer(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", middlewares.GetPathId, getEventById)
	server.GET("/queryEvent", getEventQuery)
	authRouter := server.Group("/auth")
	// use middleware
	authRouter.Use(middlewares.Authenticate)
	authRouter.POST("/event", createEvent)
	authRouter.PUT("/events/:id", updateEvent)
	authRouter.DELETE("/events/:id", deleteEvent)
	authRouter.POST("/events/:id/register", middlewares.GetPathId, registerForEvent)
	authRouter.DELETE("/events/:id/cancel", middlewares.GetPathId, cancelRegister)
	server.POST("/signUp", createUser)
	server.POST("/login", login)
}
