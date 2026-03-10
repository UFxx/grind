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

type UserHandler struct {
	*BaseHandler
}

func NewUserHandler(baseHandler *BaseHandler) *UserHandler {

	return &UserHandler{
		BaseHandler: baseHandler,
	}
}

func (handler *UserHandler) RegisterRoutes(router *gin.RouterGroup) {

	userGroup := router.Group("/users/me", handler.jwt.GinJWTAuthMiddleware())
	{
		userGroup.GET("/profile", handler.getProfileAction)

		userGroup.GET("/invite-codes", handler.getInviteCodesAction)
		userGroup.POST("/invite-codes", handler.createInviteCodeAction)
	}
}

func (handler *UserHandler) getProfileAction(c *gin.Context) {

	userID, err := handler.getUserID(c)
	if err != nil {
		handler.logger.Errorf("failed to get user id from context: %v", err)
		c.JSON(handler.getError(exceptions.NewAuthError(fmt.Errorf("unauthorized"))))
		return
	}

	var user models.User

	err = handler.db.Client.
		Preload("Inviter").
		First(&user, userID).
		Error

	if err != nil {
		handler.logger.Errorf("failed to get user profile: %v", err)

		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(handler.getError(exceptions.NewAuthError(fmt.Errorf("unauthorized"))))
			return
		}

		c.JSON(handler.getError(exceptions.NewBadRequestError(fmt.Errorf("something went wrong"))))
		return
	}

	var inviter *Profile

	if user.Inviter != nil {
		inviter = &Profile{
			ID:         user.Inviter.ID,
			TelegramID: user.Inviter.TelegramID,
			Nickname:   user.Inviter.Nickname,
			AvatarURL:  user.Inviter.AvatarURL,
			CreatedAt:  user.Inviter.CreatedAt,
		}
	}

	c.JSON(
		http.StatusOK,
		SuccessDataResponse{
			Data: GetProfileResponse{
				Profile: Profile{
					ID:         user.ID,
					TelegramID: user.TelegramID,
					Nickname:   user.Nickname,
					AvatarURL:  user.AvatarURL,
					CreatedAt:  user.CreatedAt,
				},
				Inviter: inviter,
			},
		},
	)
}

func (handler *UserHandler) getInviteCodesAction(c *gin.Context) {

	userID, err := handler.getUserID(c)
	if err != nil {
		handler.logger.Errorf("failed to get user id from context: %v", err)
		c.JSON(handler.getError(exceptions.NewAuthError(fmt.Errorf("unauthorized"))))
		return
	}

	var user models.User

	err = handler.db.Client.
		Preload("InviteCodes").
		First(&user, userID).
		Error

	if err != nil {
		handler.logger.Errorf("failed to get user with invite codes: %v", err)

		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(handler.getError(exceptions.NewAuthError(fmt.Errorf("unauthorized"))))
			return
		}

		c.JSON(handler.getError(exceptions.NewBadRequestError(fmt.Errorf("something went wrong"))))
		return
	}

	inviteCodes := make([]UserInviteCode, 0, len(user.InviteCodes))

	for _, inviteCode := range user.InviteCodes {
		inviteCodes = append(inviteCodes, UserInviteCode{
			ID:        inviteCode.ID,
			Code:      inviteCode.Code,
			Uses:      inviteCode.Uses,
			MaxUses:   inviteCode.MaxUses,
			CreatedAt: inviteCode.CreatedAt,
		})
	}

	c.JSON(
		http.StatusOK,
		SuccessDataResponse{
			Data: inviteCodes,
		},
	)
}

func (handler *UserHandler) createInviteCodeAction(c *gin.Context) {

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
