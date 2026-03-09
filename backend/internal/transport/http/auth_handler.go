package http

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sunsetsavorer/grind/internal/exceptions"
	"github.com/sunsetsavorer/grind/internal/models"
	initdata "github.com/telegram-mini-apps/init-data-golang"
	"gorm.io/gorm"
)

type AuthHandler struct {
	*BaseHandler
}

func NewAuthHandler(baseHandler *BaseHandler) *AuthHandler {

	return &AuthHandler{
		BaseHandler: baseHandler,
	}
}

func (handler *AuthHandler) RegisterRoutes(router *gin.RouterGroup) {

	authGroup := router.Group("/auth")
	{
		authGroup.POST("/telegram", handler.telegramAuthAction)
	}
}

func (handler *AuthHandler) telegramAuthAction(c *gin.Context) {

	var req TelegramAuthRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(handler.getError(exceptions.NewBadRequestError(fmt.Errorf("invalid request body"))))
		return
	}

	initData, err := handler.getTelegramInitData(c)
	if err != nil {
		handler.logger.Errorf("failed to get telegram init data: %v", err)
		c.JSON(handler.getError(err))
		return
	}

	user, err := handler.getOrCreateUser(initData, req)
	if err != nil {
		handler.logger.Errorf("failed to get or create user: %v", err)
		c.JSON(handler.getError(err))
		return
	}

	token, err := handler.jwt.CreateToken(user.ID)
	if err != nil {
		handler.logger.Errorf("failed to create JWT token: %v", err)
		c.JSON(handler.getError(exceptions.NewBadRequestError(fmt.Errorf("something went wrong"))))
		return
	}

	c.JSON(http.StatusOK, SuccessDataResponse{
		Data: TokenResponse{
			Token: token,
		},
	})
}

func (handler *AuthHandler) getTelegramInitData(c *gin.Context) (initdata.InitData, error) {

	authParts := strings.Split(c.GetHeader("Authorization"), " ")

	if len(authParts) != 2 {
		return initdata.InitData{}, exceptions.NewBadRequestError(fmt.Errorf("invalid authorization header format"))
	}

	authType := authParts[0]
	authData := authParts[1]

	if authType != "tma" {
		return initdata.InitData{}, exceptions.NewBadRequestError(fmt.Errorf("unsupported authorization type"))
	}

	if err := initdata.Validate(
		authData,
		handler.config.TelegramBotToken,
		time.Hour,
	); err != nil {
		return initdata.InitData{}, exceptions.NewBadRequestError(fmt.Errorf("invalid Telegram auth data: %v", err))
	}

	initData, err := initdata.Parse(authData)
	if err != nil {
		return initdata.InitData{}, exceptions.NewBadRequestError(fmt.Errorf("failed to parse Telegram auth data: %v", err))
	}

	return initData, nil
}

func (handler *AuthHandler) getOrCreateUser(initData initdata.InitData, req TelegramAuthRequest) (models.User, error) {

	// Try to find existing user by Telegram ID
	var user models.User

	err := handler.db.Client.
		Where("telegram_id = ?", initData.User.ID).
		First(&user).
		Error

	if err == nil {
		return user, nil
	}

	// If error is not "record not found", return error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		handler.logger.Errorf("failed to query user: %v", err)
		return models.User{}, exceptions.NewBadRequestError(fmt.Errorf("something went wrong"))
	}

	// Validate invite code for new user registration
	if req.InviteCode == "" {
		return models.User{}, exceptions.NewValidationError([]exceptions.ValidationField{
			{
				Name: "invite_code",
				Err:  fmt.Errorf("field is required"),
			},
		})
	}

	var inviteCode models.InviteCode

	err = handler.db.Client.
		Where("code = ?", req.InviteCode).
		First(&inviteCode).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.User{}, exceptions.NewNotFoundError(fmt.Errorf("invalid code"))
		}

		handler.logger.Errorf("failed to query invite code: %v", err)
		return models.User{}, exceptions.NewBadRequestError(fmt.Errorf("something went wrong"))
	}

	if inviteCode.Uses >= inviteCode.MaxUses {
		return models.User{}, exceptions.NewBadRequestError(fmt.Errorf("invite code has reached its maximum uses"))
	}

	// Begin transaction to create new user and associate with all skills
	tx := handler.db.Client.Begin()

	// Increment invite code usage count
	inviteCode.Uses++

	err = tx.Updates(&inviteCode).Error
	if err != nil {
		handler.logger.Errorf("failed to update invite code: %v", err)
		tx.Rollback()
		return models.User{}, exceptions.NewBadRequestError(fmt.Errorf("something went wrong"))
	}

	nickname := initData.User.Username
	if nickname == "" {
		nickname = fmt.Sprintf("%s %s", initData.User.FirstName, initData.User.LastName)
	}

	user = models.User{
		TelegramID: &initData.User.ID,
		Nickname:   nickname,
		InvitedBy:  uuid.NullUUID{UUID: inviteCode.CreatedBy, Valid: true},
	}

	err = tx.Create(&user).Error
	if err != nil {
		handler.logger.Errorf("failed to create user: %v", err)
		tx.Rollback()
		return models.User{}, exceptions.NewBadRequestError(fmt.Errorf("something went wrong"))
	}

	var skills []models.Skill

	err = tx.Find(&skills).Error
	if err != nil {
		handler.logger.Errorf("failed to query skills: %v", err)
		tx.Rollback()
		return models.User{}, exceptions.NewBadRequestError(fmt.Errorf("something went wrong"))
	}

	userSkills := make([]models.UserSkill, 0, len(skills))

	for _, skill := range skills {
		userSkills = append(userSkills, models.UserSkill{
			UserID:  user.ID,
			SkillID: skill.ID,
		})
	}

	err = tx.Create(&userSkills).Error
	if err != nil {
		handler.logger.Errorf("failed to create user skills: %v", err)
		tx.Rollback()
		return models.User{}, exceptions.NewBadRequestError(fmt.Errorf("something went wrong"))
	}

	tx.Commit()

	return user, nil
}
