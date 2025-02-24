package route

import (
	"go-rest/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterServer(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEventById)
	server.GET("/queryEvent", getEventQuery)
	authRouter := server.Group("/auth")
	// use middleware
	authRouter.Use(middlewares.Authenticate)
	authRouter.POST("/event", createEvent)
	authRouter.PUT("/events/:id", updateEvent)
	authRouter.DELETE("/events/:id", deleteEvent)
	authRouter.POST("/events/:id/register", registerForEvent)
	authRouter.DELETE("/events/:id/cancel", cancelRegister)
	server.POST("/signUp", createUser)
	server.POST("/login", login)
}
