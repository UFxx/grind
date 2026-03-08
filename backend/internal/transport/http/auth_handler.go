package http

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
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

	initData, err := handler.getTelegramInitData(c)
	if err != nil {
		c.JSON(handler.getError(err))
		return
	}

	user, err := handler.getOrCreateUser(initData)
	if err != nil {
		c.JSON(handler.getError(err))
		return
	}

	token, err := handler.jwt.CreateToken(user.ID)
	if err != nil {
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

func (handler *AuthHandler) getOrCreateUser(initData initdata.InitData) (models.User, error) {

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
		return models.User{}, exceptions.NewBadRequestError(fmt.Errorf("something went wrong"))
	}

	// Begin transaction to create new user and associate with all skills
	tx := handler.db.Client.Begin()

	nickname := initData.User.Username
	if nickname == "" {
		nickname = fmt.Sprintf("%s %s", initData.User.FirstName, initData.User.LastName)
	}

	user = models.User{
		TelegramID: &initData.User.ID,
		Nickname:   nickname,
	}

	err = handler.db.Client.
		Create(&user).
		Error

	if err != nil {
		tx.Rollback()
		return models.User{}, exceptions.NewBadRequestError(fmt.Errorf("something went wrong"))
	}

	var skills []models.Skill

	err = handler.db.Client.
		Find(&skills).
		Error

	if err != nil {
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

	err = handler.db.Client.
		Create(&userSkills).
		Error

	if err != nil {
		tx.Rollback()
		return models.User{}, exceptions.NewBadRequestError(fmt.Errorf("something went wrong"))
	}

	tx.Commit()

	return user, nil
}
