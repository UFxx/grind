package http

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

		userGroup.GET("/skills", handler.getSkillsAction)
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

func (handler *UserHandler) getSkillsAction(c *gin.Context) {

	userID, err := handler.getUserID(c)
	if err != nil {
		handler.logger.Errorf("failed to get user id from context: %v", err)
		c.JSON(handler.getError(exceptions.NewAuthError(fmt.Errorf("unauthorized"))))
		return
	}

	// Find user with skills
	var user models.User

	err = handler.db.Client.
		Preload("Skills").
		Preload("Skills.BaseSkill").
		Preload("Skills.ParentSkill").
		First(&user, userID).
		Error

	if err != nil {
		handler.logger.Errorf("failed to get user with skills: %v", err)

		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(handler.getError(exceptions.NewAuthError(fmt.Errorf("unauthorized"))))
			return
		}

		c.JSON(handler.getError(exceptions.NewBadRequestError(fmt.Errorf("something went wrong"))))
		return
	}

	// Create subskills map
	subskillsMap := make(map[uuid.UUID][]models.UserSkill)

	for _, userSkill := range user.Skills {
		// Skip root skills
		if userSkill.ParentSkill == nil {
			continue
		}

		subskillsMap[userSkill.ParentSkill.ID] = append(subskillsMap[userSkill.ParentSkill.ID], userSkill)
	}

	// Build skills tree
	rootSkills := make([]RootSkill, 0, len(subskillsMap))

	for _, userSkill := range user.Skills {

		isRoot := userSkill.ParentSkill == nil

		// Skip subskills, they will be processed with their parent skill
		if !isRoot {
			continue
		}

		subskills, hasSubskills := subskillsMap[userSkill.ID]

		if !hasSubskills {
			subskills = make([]models.UserSkill, 0)
		}

		subskillItems := make([]Subskill, 0, len(subskills))

		for _, subskill := range subskills {
			subskillProgress := handler.skillService.CalcProgress(subskill.TotalXP)

			subskillItems = append(subskillItems, Subskill{
				ID:    subskill.ID,
				Title: subskill.Title,
				SkillProgress: SkillProgress{
					CurrentLevel:        subskillProgress.CurrentLevel,
					NextLevel:           subskillProgress.NextLevel,
					TotalXP:             subskillProgress.TotalXP,
					CurrentLevelStartXP: subskillProgress.CurrentLevelStartXP,
					NextLevelStartXP:    subskillProgress.NextLevelStartXP,
					XPToNextLevel:       subskillProgress.XPToNextLevel,
				},
			})
		}

		rootSkillProgress := handler.skillService.CalcProgress(userSkill.TotalXP)

		rootSkills = append(rootSkills, RootSkill{
			ID:    userSkill.ID,
			Title: userSkill.Title,
			SkillProgress: SkillProgress{
				CurrentLevel:        rootSkillProgress.CurrentLevel,
				NextLevel:           rootSkillProgress.NextLevel,
				TotalXP:             rootSkillProgress.TotalXP,
				CurrentLevelStartXP: rootSkillProgress.CurrentLevelStartXP,
				NextLevelStartXP:    rootSkillProgress.NextLevelStartXP,
				XPToNextLevel:       rootSkillProgress.XPToNextLevel,
			},
			Items: subskillItems,
		})
	}

	c.JSON(
		http.StatusOK,
		SuccessDataResponse{
			Data: rootSkills,
		},
	)
}
