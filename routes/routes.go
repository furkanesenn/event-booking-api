package routes

import (
	"furkanesen.com/restapi/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", unRegisterForEvent)

	server.GET("/events", getEvents)    // Creating an endpoint for the HTTP Engine (GET)
	server.GET("/events/:id", getEvent) // Creating an endpoint for the HTTP Engine (GET)

	server.POST("/signup", SignUp)
	server.POST("login", LogIn)
}
