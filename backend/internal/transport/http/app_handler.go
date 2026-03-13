package http

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sunsetsavorer/grind/internal/exceptions"
	"github.com/sunsetsavorer/grind/internal/models"
	"gorm.io/gorm"
)

type AppHandler struct {
	*BaseHandler
}

func NewAppHandler(baseHandler *BaseHandler) *AppHandler {

	return &AppHandler{BaseHandler: baseHandler}
}

func (handler *AppHandler) RegisterRoutes(router *gin.RouterGroup) {

	appGroup := router.Group("/app", handler.jwt.GinJWTAuthMiddleware())
	{
		appGroup.POST("/feedback", handler.sendFeedbackAction)
	}
}

func (handler *AppHandler) sendFeedbackAction(ctx *gin.Context) {

	userID, err := handler.getUserID(ctx)
	if err != nil {
		handler.logger.Errorf("%v", err)

		ctx.JSON(handler.getError(exceptions.NewAuthError(fmt.Errorf("unauthorized"))))
		return
	}

	var user models.User

	err = handler.db.Client.First(&user, userID).Error
	if err != nil {
		handler.logger.Errorf("failed to find user by id: %v", err)

		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(handler.getError(exceptions.NewAuthError(fmt.Errorf("unauthorized"))))
			return
		}

		ctx.JSON(handler.getError(exceptions.NewAuthError(fmt.Errorf("something went wrong"))))
		return
	}

	var req SendFeedbackRequestBody

	if err := ctx.ShouldBindJSON(&req); err != nil {
		handler.logger.Errorf("failed to bind `send feedback` request body: %v", err)

		ctx.JSON(handler.getError(exceptions.NewBadRequestError(fmt.Errorf("invalid request body"))))
		return
	}

	if err := handler.validator.Struct(&req); err != nil {
		ctx.JSON(handler.getError(err))
		return
	}

	feedback := models.Feedback{
		Sender:  req.Sender,
		Rating:  req.Rating,
		Message: req.Message,
	}

	err = handler.db.Client.Create(&feedback).Error
	if err != nil {
		handler.logger.Errorf("failed to save feedback: %v", err)

		ctx.JSON(handler.getError(exceptions.NewBadRequestError(fmt.Errorf("something went wrong"))))
		return
	}

	ctx.JSON(
		http.StatusOK,
		SuccessDataResponse{
			Data: []struct{}{},
		},
	)
}
