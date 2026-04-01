package http

import (
	"errors"
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

		ctx.JSON(handler.getError(exceptions.NewAuthError(errUnauthorized)))
		return
	}

	var user models.User

	err = handler.db.Client.First(&user, userID).Error
	if err != nil {
		handler.logger.Errorf("failed to find user by id: %v", err)

		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(handler.getError(exceptions.NewAuthError(errUnauthorized)))
			return
		}

		ctx.JSON(handler.getError(exceptions.NewInternalServerError(errSomethingWentWrong)))
		return
	}

	var req SendFeedbackRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		handler.logger.Errorf("failed to bind `send feedback` request body: %v", err)

		ctx.JSON(handler.getError(exceptions.NewBadRequestError(errInvalidRequestBody)))
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

		ctx.JSON(handler.getError(exceptions.NewInternalServerError(errSomethingWentWrong)))
		return
	}

	ctx.JSON(
		http.StatusOK,
		SuccessDataResponse{
			Data: []struct{}{},
		},
	)
}
