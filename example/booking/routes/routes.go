package routes

import "github.com/gin-gonic/gin"

func RegisteringRoutes(server *gin.Engine) {

	server.GET("/events", getEvents)    // GET, PUT, POST, DELETE, PATCH
	server.GET("/events/:id", getEvent) // /events/1 , /events/4
	server.POST("/events", createEvent)
	server.PUT("/event/:id", updateEvent)
	server.DELETE("/event/:id", deleteEvent)

}
