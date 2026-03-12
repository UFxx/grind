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
		userGroup.GET("/profile", handler.getMyProfileAction)

		userGroup.GET("/invite-codes", handler.getMyInviteCodesAction)
		userGroup.POST("/invite-codes", handler.createInviteCodeAction)

		userGroup.GET("/skills", handler.getMySkillsAction)
		userGroup.GET("/skills-progress", handler.getMySkillsProgressAction)

		userGroup.GET("/events", handler.getMyEventsAction)
		userGroup.POST("/events", handler.createEventAction)
	}
}

func (handler *UserHandler) getMyProfileAction(ctx *gin.Context) {

	userID, err := handler.getUserID(ctx)
	if err != nil {
		handler.logger.Errorf("failed to get user id from context: %v", err)
		ctx.JSON(handler.getError(exceptions.NewAuthError(fmt.Errorf("unauthorized"))))
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
			ctx.JSON(handler.getError(exceptions.NewAuthError(fmt.Errorf("unauthorized"))))
			return
		}

		ctx.JSON(handler.getError(exceptions.NewBadRequestError(fmt.Errorf("something went wrong"))))
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

	ctx.JSON(
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

func (handler *UserHandler) getMyInviteCodesAction(ctx *gin.Context) {

	userID, err := handler.getUserID(ctx)
	if err != nil {
		handler.logger.Errorf("failed to get user id from context: %v", err)
		ctx.JSON(handler.getError(exceptions.NewAuthError(fmt.Errorf("unauthorized"))))
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
			ctx.JSON(handler.getError(exceptions.NewAuthError(fmt.Errorf("unauthorized"))))
			return
		}

		ctx.JSON(handler.getError(exceptions.NewBadRequestError(fmt.Errorf("something went wrong"))))
		return
	}

	inviteCodes := make([]InviteCode, 0, len(user.InviteCodes))

	for _, inviteCode := range user.InviteCodes {
		inviteCodes = append(inviteCodes, InviteCode{
			ID:        inviteCode.ID,
			Code:      inviteCode.Code,
			Uses:      inviteCode.Uses,
			MaxUses:   inviteCode.MaxUses,
			CreatedAt: inviteCode.CreatedAt,
		})
	}

	ctx.JSON(
		http.StatusOK,
		SuccessDataResponse{
			Data: inviteCodes,
		},
	)
}

func (handler *UserHandler) createInviteCodeAction(ctx *gin.Context) {

	userID, err := handler.getUserID(ctx)
	if err != nil {
		handler.logger.Errorf("failed to get user id from context: %v", err)
		ctx.JSON(handler.getError(exceptions.NewAuthError(fmt.Errorf("unauthorized"))))
		return
	}

	var req CreateInviteCodeRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		handler.logger.Errorf("failed to bind request body: %v", err)
		ctx.JSON(handler.getError(exceptions.NewBadRequestError(fmt.Errorf("invalid request body"))))
		return
	}

	if err := handler.validator.Struct(&req); err != nil {
		ctx.JSON(handler.getError(err))
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

func (handler *UserHandler) getMySkillsProgressAction(ctx *gin.Context) {

	userID, err := handler.getUserID(ctx)
	if err != nil {
		handler.logger.Errorf("failed to get user id from context: %v", err)
		ctx.JSON(handler.getError(exceptions.NewAuthError(fmt.Errorf("unauthorized"))))
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
			ctx.JSON(handler.getError(exceptions.NewAuthError(fmt.Errorf("unauthorized"))))
			return
		}

		ctx.JSON(handler.getError(exceptions.NewBadRequestError(fmt.Errorf("something went wrong"))))
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

	ctx.JSON(
		http.StatusOK,
		SuccessDataResponse{
			Data: rootSkills,
		},
	)
}

func (handler *UserHandler) createEventAction(ctx *gin.Context) {

	userID, err := handler.getUserID(ctx)
	if err != nil {
		handler.logger.Errorf("failed to get user id from context: %v", err)
		ctx.JSON(handler.getError(exceptions.NewAuthError(fmt.Errorf("unauthorized"))))
		return
	}

	var req ManuallyCreatedEventRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		handler.logger.Errorf("failed to bind request body: %v", err)
		ctx.JSON(handler.getError(exceptions.NewBadRequestError(fmt.Errorf("invalid request body"))))
		return
	}

	if err := handler.validator.Struct(&req); err != nil {
		ctx.JSON(handler.getError(err))
		return
	}

	eventReward, err := handler.skillService.CalcEventReward(
		req.SkillWeights,
		req.HasImpact,
		req.IsHard,
		req.IsNew,
	)

	if err != nil {
		handler.logger.Errorf("failed to calc event reward: %v", err)

		ctx.JSON(handler.getError(exceptions.NewBadRequestError(err)))
		return
	}

	var user models.User

	err = handler.db.Client.
		First(&user, userID).
		Error

	if err != nil {
		handler.logger.Errorf("failed to get user: %v", err)

		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(handler.getError(exceptions.NewAuthError(fmt.Errorf("unauthorized"))))
			return
		}

		ctx.JSON(handler.getError(exceptions.NewBadRequestError(fmt.Errorf("something went wrong"))))
		return
	}

	var skillIDs []uuid.UUID

	for _, skillWeight := range req.SkillWeights {
		skillIDs = append(skillIDs, skillWeight.SkillID)
	}

	var userSkills []models.UserSkill

	err = handler.db.Client.
		Where("user_id = ? AND id IN ?", userID, skillIDs).
		Find(&userSkills).
		Error

	if err != nil {
		handler.logger.Errorf("failed to get user skills: %v", err)
		ctx.JSON(handler.getError(exceptions.NewBadRequestError(fmt.Errorf("something went wrong"))))
		return
	}

	if len(userSkills) != len(skillIDs) {
		ctx.JSON(handler.getError(exceptions.NewValidationError([]exceptions.ValidationField{
			{Name: "skill_weights", Err: fmt.Errorf("some skill_ids do not exist or do not belong to user")},
		})))
		return
	}

	// Create event and rewards in transaction

	tx := handler.db.Client.Begin()

	event := models.Event{
		Title:       req.Title,
		UserID:      userID,
		Source:      "manual",
		EventTypeID: req.EventTypeID,
		HasImpact:   req.HasImpact,
		IsHard:      req.IsHard,
		IsNew:       req.IsNew,
	}

	err = tx.Create(&event).Error
	if err != nil {
		handler.logger.Errorf("failed to create event: %v", err)
		tx.Rollback()
		ctx.JSON(handler.getError(exceptions.NewBadRequestError(fmt.Errorf("something went wrong"))))
		return
	}

	rewards := make([]models.EventReward, 0, len(eventReward.SkillRewards))

	for _, skillReward := range eventReward.SkillRewards {
		rewards = append(rewards, models.EventReward{
			EventID:     event.ID,
			UserSkillID: skillReward.SkillID,
			XPAmount:    skillReward.XPAmount,
		})

		err = tx.Model(&models.UserSkill{}).
			Where("id = ?", skillReward.SkillID).
			UpdateColumn("total_xp", gorm.Expr("total_xp + ?", skillReward.XPAmount)).
			Error

		if err != nil {
			handler.logger.Errorf("failed to update user skill xp: %v", err)
			tx.Rollback()
			ctx.JSON(handler.getError(exceptions.NewBadRequestError(fmt.Errorf("something went wrong"))))
			return
		}
	}

	err = tx.Create(&rewards).Error
	if err != nil {
		handler.logger.Errorf("failed to create event rewards: %v", err)
		tx.Rollback()
		ctx.JSON(handler.getError(exceptions.NewBadRequestError(fmt.Errorf("something went wrong"))))
		return
	}

	tx.Commit()

	ctx.JSON(
		http.StatusOK,
		SuccessDataResponse{
			Data: struct{}{},
		},
	)
}

func (handler *UserHandler) getMySkillsAction(ctx *gin.Context) {

	userID, err := handler.getUserID(ctx)
	if err != nil {
		handler.logger.Errorf("failed to get user id from context: %v", err)
		ctx.JSON(handler.getError(exceptions.NewAuthError(fmt.Errorf("unauthorized"))))
		return
	}

	var user models.User

	err = handler.db.Client.
		First(&user, userID).
		Error

	if err != nil {
		handler.logger.Errorf("failed to get user: %v", err)

		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(handler.getError(exceptions.NewAuthError(fmt.Errorf("unauthorized"))))
			return
		}

		ctx.JSON(handler.getError(exceptions.NewBadRequestError(fmt.Errorf("something went wrong"))))
	}

	var userSkills []models.UserSkill

	err = handler.db.Client.
		Where("user_id = ?", userID).
		Order("title ASC").
		Find(&userSkills).
		Error

	if err != nil {
		handler.logger.Errorf("failed to get user skills: %v", err)
		ctx.JSON(handler.getError(exceptions.NewBadRequestError(fmt.Errorf("something went wrong"))))
		return
	}

	response := make([]GetSkillsResponseItem, 0, len(userSkills))

	for _, userSkill := range userSkills {
		response = append(response, GetSkillsResponseItem{
			ID:    userSkill.ID,
			Title: userSkill.Title,
		})
	}

	ctx.JSON(
		http.StatusOK,
		SuccessDataResponse{
			Data: response,
		},
	)
}

func (handler *UserHandler) getMyEventsAction(ctx *gin.Context) {

	userID, err := handler.getUserID(ctx)
	if err != nil {
		handler.logger.Errorf("failed to get user id from context: %v", err)

		ctx.JSON(handler.getError(exceptions.NewAuthError(fmt.Errorf("unauthorized"))))
		return
	}

	var req GetEventsRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		handler.logger.Errorf("failed to bind query params: %v", err)

		ctx.JSON(handler.getError(exceptions.NewBadRequestError(fmt.Errorf("invalid query params"))))
		return
	}

	if err := handler.validator.Struct(&req); err != nil {
		ctx.JSON(handler.getError(err))
		return
	}

	var user models.User

	err = handler.db.Client.
		First(&user, userID).
		Error

	if err != nil {
		handler.logger.Errorf("failed to get user: %v", err)

		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(handler.getError(exceptions.NewAuthError(fmt.Errorf("unauthorized"))))
			return
		}

		ctx.JSON(handler.getError(exceptions.NewBadRequestError(fmt.Errorf("something went wrong"))))
		return
	}

	var total int64
	var events []models.Event

	query := handler.db.Client.
		Where("user_id = ?", userID).
		Order("created_at DESC")

	err = query.Model(&events).Count(&total).Error
	if err != nil {
		handler.logger.Errorf("failed to count user events: %v", err)

		ctx.JSON(handler.getError(exceptions.NewBadRequestError(fmt.Errorf("something went wrong"))))
		return
	}

	offset := (req.Page - 1) * req.Limit

	err = query.
		Preload("EventType").
		Preload("Rewards").
		Preload("Rewards.UserSkill").
		Limit(req.Limit).
		Offset(offset).
		Find(&events).
		Error

	if err != nil {
		handler.logger.Errorf("failed to get user events: %v", err)

		ctx.JSON(handler.getError(exceptions.NewBadRequestError(fmt.Errorf("something went wrong"))))
		return
	}

	responseItems := make([]GetEventsResponseItem, 0, len(events))

	for _, event := range events {

		rewards := make([]EventReward, 0, len(event.Rewards))

		for _, reward := range event.Rewards {
			rewards = append(rewards, EventReward{
				SkillID:  reward.UserSkill.ID,
				Title:    reward.UserSkill.Title,
				XPAmount: reward.XPAmount,
			})
		}

		responseItems = append(responseItems, GetEventsResponseItem{
			ID:        event.ID,
			Title:     event.Title,
			EventType: EventType{ID: event.EventType.ID, Title: event.EventType.Title},
			HasImpact: event.HasImpact,
			IsNew:     event.IsNew,
			IsHard:    event.IsHard,
			Rewards:   rewards,
			CreatedAt: event.CreatedAt,
		})
	}

	response := GetEventsResponse{
		PaginationResponse: GetPaginationResponse(total, req.Page, req.Limit),
		Items:              responseItems,
	}

	ctx.JSON(
		http.StatusOK,
		SuccessDataResponse{
			Data: response,
		},
	)
}
