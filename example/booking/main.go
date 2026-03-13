package main

import (
	"booking/db"
	models "booking/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/events", getEvents) // GET, PUT, POST, DELETE, PATCH
	server.POST("/events", createEvent)

	server.Run(":8080") // localhost:8080

}

func getEvents(context *gin.Context) {
	event, err := models.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error!"})
		return
	}
	context.JSON(http.StatusOK, event)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindBodyWithJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data. error"})
		fmt.Println(err)
		return
	}

	event.ID = 1
	event.UserID = 1

	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not create event. Try again later..."})
		fmt.Println(err)
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created !", "events": event})

}
