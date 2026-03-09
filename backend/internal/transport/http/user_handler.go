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

	userGroup := router.Group("/users", handler.jwt.GinJWTAuthMiddleware())
	{
		userGroup.GET("/me/profile", handler.getProfileAction)
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

	var inviter *UserProfile

	if user.Inviter != nil {
		inviter = &UserProfile{
			ID:         user.Inviter.ID,
			TelegramID: user.Inviter.TelegramID,
			Nickname:   user.Inviter.Nickname,
			CreatedAt:  user.Inviter.CreatedAt,
		}
	}

	c.JSON(
		http.StatusOK,
		SuccessDataResponse{
			Data: GetUserProfileResponse{
				UserProfile: UserProfile{
					ID:         user.ID,
					TelegramID: user.TelegramID,
					Nickname:   user.Nickname,
					CreatedAt:  user.CreatedAt,
				},
				Inviter: inviter,
			},
		},
	)
}
