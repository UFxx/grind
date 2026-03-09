package http

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sunsetsavorer/grind/internal/exceptions"
	"github.com/sunsetsavorer/grind/internal/models"
)

type InviteCodeHandler struct {
	*BaseHandler
}

func NewInviteCodeHandler(baseHandler *BaseHandler) *InviteCodeHandler {

	return &InviteCodeHandler{
		BaseHandler: baseHandler,
	}
}

func (handler *InviteCodeHandler) RegisterRoutes(router *gin.RouterGroup) {

	inviteCodeGroup := router.Group("/invite-codes", handler.jwt.GinJWTAuthMiddleware())
	{
		inviteCodeGroup.POST("", handler.createCodeAction)
	}
}

func (handler *InviteCodeHandler) createCodeAction(c *gin.Context) {

	userID, err := handler.getUserID(c)
	if err != nil {
		handler.logger.Errorf("failed to get user id from context: %v", err)
		c.JSON(handler.getError(exceptions.NewAuthError(fmt.Errorf("unauthorized"))))
		return
	}

	var req CreateInviteCodeRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		handler.logger.Errorf("failed to bind request body: %v", err)
		c.JSON(handler.getError(exceptions.NewBadRequestError(fmt.Errorf("invalid request body"))))
		return
	}

	if err := handler.validator.Struct(&req); err != nil {
		c.JSON(handler.getError(err))
		return
	}

	inviteCode := models.InviteCode{
		Code:      req.Code,
		MaxUses:   req.MaxUses,
		CreatedBy: userID,
	}

	err = handler.db.Client.
		Create(&inviteCode).
		Error

	if err != nil {
		handler.logger.Errorf("failed to create invite code: %v", err)
		c.JSON(handler.getError(exceptions.NewBadRequestError(fmt.Errorf("something went wrong"))))
		return
	}

	c.JSON(
		http.StatusOK,
		SuccessDataResponse{
			Data: []struct{}{},
		},
	)
}
