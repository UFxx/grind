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

	activityCategoryGroup := router.Group("/activity-categories", handler.jwt.GinJWTAuthMiddleware())
	{
		activityCategoryGroup.GET("", handler.getActivityCategoriesAction)
	}
}

func (handler *ActivityHandler) getActivityCategoriesAction(ctx *gin.Context) {

	_, err := handler.getUserID(ctx)
	if err != nil {
		handler.logger.Errorf("failed to get user id from context: %v", err)

		ctx.JSON(handler.getError(exceptions.NewAuthError(errUnauthorized)))
		return
	}

	var activityCategories []models.ActivityCategory

	err = handler.db.Client.
		Order("display_name ASC").
		Find(&activityCategories).
		Error

	if err != nil {
		handler.logger.Errorf("failed to get activity categories: %v", err)

		ctx.JSON(handler.getError(exceptions.NewInternalServerError(errSomethingWentWrong)))
		return
	}

	responseItems := make([]GetActivityCategoriesResponseItem, 0, len(activityCategories))

	for _, activityCategory := range activityCategories {
		responseItems = append(responseItems, GetActivityCategoriesResponseItem{
			ID:   activityCategory.ID,
			Code: activityCategory.Code,
			Name: activityCategory.DisplayName,
		})
	}

	ctx.JSON(
		http.StatusOK,
		SuccessDataResponse{
			Data: responseItems,
		},
	)
}
