package http

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sunsetsavorer/grind/internal/exceptions"
	"github.com/sunsetsavorer/grind/internal/models"
)

type EventHandler struct {
	*BaseHandler
}

func NewEventHandler(baseHandler *BaseHandler) *EventHandler {

	return &EventHandler{
		BaseHandler: baseHandler,
	}
}

func (handler *EventHandler) RegisterRoutes(router *gin.RouterGroup) {

	eventTypeGroup := router.Group("/event-types", handler.jwt.GinJWTAuthMiddleware())
	{
		eventTypeGroup.GET("", handler.getEventTypesAction)
	}
}

func (handler *EventHandler) getEventTypesAction(ctx *gin.Context) {

	_, err := handler.getUserID(ctx)
	if err != nil {
		handler.logger.Errorf("failed to get user id from context: %v", err)

		ctx.JSON(handler.getError(exceptions.NewAuthError(fmt.Errorf("unauthorized"))))
		return
	}

	var eventTypes []models.EventType

	err = handler.db.Client.
		Order("title ASC").
		Find(&eventTypes).
		Error

	if err != nil {
		handler.logger.Errorf("failed to get event types: %v", err)

		ctx.JSON(handler.getError(exceptions.NewBadRequestError(fmt.Errorf("something went wrong"))))
		return
	}

	responseItems := make([]GetEventTypesResponseItem, 0, len(eventTypes))

	for _, eventType := range eventTypes {
		responseItems = append(responseItems, GetEventTypesResponseItem{
			ID:    eventType.ID,
			Title: eventType.Title,
		})
	}

	ctx.JSON(
		http.StatusOK,
		SuccessDataResponse{
			Data: responseItems,
		},
	)
}
