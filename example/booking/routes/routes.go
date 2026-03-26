package routes

import (
	"booking/middleware"

	"github.com/gin-gonic/gin"
)

func RegisteringRoutes(server *gin.Engine) {

	server.GET("/events", getEvents)    // GET, PUT, POST, DELETE, PATCH
	server.GET("/events/:id", getEvent) // /events/1 , /events/4
	Authenticated := server.Group("/")
	Authenticated.Use(middleware.Authenticate)
	Authenticated.POST("/events", createEvent)
	Authenticated.PUT("/event/:id", updateEvent)
	Authenticated.DELETE("/event/:id", deleteEvent)
	Authenticated.POST("/event/:id/register", registerForEvent)
	Authenticated.DELETE("/event/:id/register", cancelRegisteration)
	server.PUT("/signup", signup)
	server.PUT("/login", login)

}
