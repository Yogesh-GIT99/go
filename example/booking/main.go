package main

import (
	models "booking/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/events", getEvents) // GET, PUT, POST, DELETE, PATCH
	server.POST("/events", createEvent)

	server.Run(":8080") // localhost:8080

}

func getEvents(context *gin.Context) {
	event := models.GetAllEvents()
	context.JSON(http.StatusOK, event)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindBodyWithJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data."})
		return
	}

	event.ID = 1
	event.UserID = 1
	event.Save()
	context.JSON(http.StatusCreated, gin.H{"message": "Event created !", "events": event})

}
