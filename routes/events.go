package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yatharthx/events-rest-api/models"
	"github.com/yatharthx/events-rest-api/utils"
)

func getEvents(ctx *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch events. Please try again.",
		})
		return
	}
	ctx.JSON(http.StatusOK, events)
}

func getEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse event ID",
		})
		return
	}

	event, err := models.GetEventByID(eventId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch event",
		})
		return
	}

	ctx.JSON(http.StatusOK, event)
}

func createEvent(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")

	if token == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Not authorized",
		})
		return
	}

	userId, err := utils.VerifyToken(token)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Not authorized",
		})
		return
	}

	var event models.Event
	err = ctx.ShouldBindJSON(&event)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request data.",
		})
		return
	}

	event.UserID = userId
	event.Save()

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Event created successfully",
		"event":   event,
	})
}

func updateEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse event ID",
		})
		return
	}

	_, err = models.GetEventByID(eventId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch the event",
		})
		return
	}

	var updatedEvent models.Event
	err = ctx.ShouldBindJSON(&updatedEvent)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request data",
		})
		return
	}

	updatedEvent.ID = eventId
	err = updatedEvent.Update()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to update the event",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Event updated",
		"event":   updatedEvent,
	})
}

func deleteEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse event ID",
		})
		return
	}

	event, err := models.GetEventByID(eventId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch the event",
		})
		return
	}

	err = event.Delete()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not remove the event",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Event removed successfully",
	})
}
