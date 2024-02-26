package main

import (
	"danielokyere/RESTCRUD/db"
	"danielokyere/RESTCRUD/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvents)

	server.Run(":8080")

}
func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		 context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events. Try again"})
	}
	context.JSON(http.StatusOK, events)
}

func createEvents(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}
	event.ID = 1
	event.UserID = 1
	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create events. Try again"})
		return 
	}
	
	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}
