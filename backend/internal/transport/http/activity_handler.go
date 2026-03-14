package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sunsetsavorer/grind/internal/exceptions"
	"github.com/sunsetsavorer/grind/internal/models"
)

type ActivityHandler struct {
	*BaseHandler
}

func NewActivityHandler(baseHandler *BaseHandler) *ActivityHandler {

	return &ActivityHandler{
		BaseHandler: baseHandler,
	}
}

func (handler *ActivityHandler) RegisterRoutes(router *gin.RouterGroup) {

	activityTypeGroup := router.Group("/activity-types", handler.jwt.GinJWTAuthMiddleware())
	{
		activityTypeGroup.GET("", handler.getActivityTypesAction)
	}
}

func (handler *ActivityHandler) getActivityTypesAction(ctx *gin.Context) {

	_, err := handler.getUserID(ctx)
	if err != nil {
		handler.logger.Errorf("failed to get user id from context: %v", err)

		ctx.JSON(handler.getError(exceptions.NewAuthError(errUnauthorized)))
		return
	}

	var activityTypes []models.ActivityType

	err = handler.db.Client.
		Order("name ASC").
		Find(&activityTypes).
		Error

	if err != nil {
		handler.logger.Errorf("failed to get activity types: %v", err)

		ctx.JSON(handler.getError(exceptions.NewInternalServerError(errSomethingWentWrong)))
		return
	}

	responseItems := make([]GetActivityTypesResponseItem, 0, len(activityTypes))

	for _, activityType := range activityTypes {
		responseItems = append(responseItems, GetActivityTypesResponseItem{
			ID:   activityType.ID,
			Code: activityType.Code,
			Name: activityType.Name,
		})
	}

	ctx.JSON(
		http.StatusOK,
		SuccessDataResponse{
			Data: responseItems,
		},
	)
}
