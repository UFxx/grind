package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sunsetsavorer/grind/internal/enums"
	"github.com/sunsetsavorer/grind/internal/exceptions"
	"github.com/sunsetsavorer/grind/internal/models"
	"github.com/sunsetsavorer/grind/internal/skill"
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
		userGroup.DELETE("", handler.deleteMyProfileAction)

		userGroup.GET("/invite-codes", handler.getMyInviteCodesAction)
		userGroup.POST("/invite-codes", handler.createInviteCodeAction)
		userGroup.DELETE("/invite-codes/:code_id", handler.deleteMyInviteCodeAction)

		userGroup.GET("/skills", handler.getMySkillsAction)
		userGroup.GET("/skills-progress", handler.getMySkillsProgressAction)

		userGroup.GET("/activities", handler.getMyActivitiesAction)
		userGroup.POST("/activities", handler.createActivityAction)
		userGroup.DELETE("/activities/:activity_id", handler.deleteMyActivityAction)
	}
}

func (handler *UserHandler) getMyProfileAction(ctx *gin.Context) {

	userID, err := handler.getUserID(ctx)
	if err != nil {
		handler.logger.Errorf("failed to get user id from context: %v", err)

		ctx.JSON(handler.getError(exceptions.NewAuthError(errUnauthorized)))
		return
	}

	var user models.User

	err = handler.db.Client.
		Preload("Skills").
		Preload("Inviter").
		Preload("Inviter.Skills").
		Preload("Inviter.Rank").
		Preload("Rank").
		First(&user, userID).
		Error

	if err != nil {
		handler.logger.Errorf("failed to get user profile: %v", err)

		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(handler.getError(exceptions.NewAuthError(errUnauthorized)))
			return
		}

		ctx.JSON(handler.getError(exceptions.NewInternalServerError(errSomethingWentWrong)))
		return
	}

	var inviter *Profile

	if user.InvitedBy.Valid {
		inviter = &Profile{
			ID:         user.Inviter.ID,
			TelegramID: user.Inviter.TelegramID,
			Name:       user.Inviter.Name,
			Rank: Rank{
				ID:   user.Inviter.Rank.ID,
				Name: user.Inviter.Rank.DisplayName,
			},
			AvatarURL: user.Inviter.AvatarURL,
			Level:     handler.calcUserLevel(user.Inviter.Skills),
			CreatedAt: user.Inviter.CreatedAt,
		}
	}

	ctx.JSON(
		http.StatusOK,
		SuccessDataResponse{
			Data: GetProfileResponse{
				Profile: Profile{
					ID:         user.ID,
					TelegramID: user.TelegramID,
					Name:       user.Name,
					Rank: Rank{
						ID:   user.Rank.ID,
						Name: user.Rank.DisplayName,
					},
					AvatarURL: user.AvatarURL,
					Level:     handler.calcUserLevel(user.Skills),
					CreatedAt: user.CreatedAt,
				},
				Inviter: inviter,
			},
		},
	)
}

func (handler *UserHandler) deleteMyProfileAction(ctx *gin.Context) {

	userID, err := handler.getUserID(ctx)
	if err != nil {
		handler.logger.Errorf("failed to get user id from context: %v", err)

		ctx.JSON(handler.getError(exceptions.NewAuthError(errUnauthorized)))
		return
	}

	var user models.User

	err = handler.db.Client.
		First(&user, userID).
		Error

	if err != nil {
		handler.logger.Errorf("failed to get user: %v", err)

		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(handler.getError(exceptions.NewAuthError(errUnauthorized)))
			return
		}

		ctx.JSON(handler.getError(exceptions.NewInternalServerError(errSomethingWentWrong)))
		return
	}

	err = handler.db.Client.
		Delete(&models.User{}, userID).
		Error

	if err != nil {
		handler.logger.Errorf("failed to delete user: %v", err)

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

func (handler *UserHandler) calcUserLevel(userSkills []models.UserSkill) SkillProgress {

	var totalXP int

	for _, userSkill := range userSkills {
		if userSkill.BaseSkillID.Valid {
			totalXP += userSkill.TotalXP
		}
	}

	progress := handler.skillService.CalcProgress(totalXP)

	return SkillProgress{
		CurrentLevel:        progress.CurrentLevel,
		NextLevel:           progress.NextLevel,
		TotalXP:             progress.TotalXP,
		CurrentLevelStartXP: progress.CurrentLevelStartXP,
		NextLevelStartXP:    progress.NextLevelStartXP,
		XPToNextLevel:       progress.XPToNextLevel,
	}
}

func (handler *UserHandler) getMyInviteCodesAction(ctx *gin.Context) {

	userID, err := handler.getUserID(ctx)
	if err != nil {
		handler.logger.Errorf("failed to get user id from context: %v", err)

		ctx.JSON(handler.getError(exceptions.NewAuthError(errUnauthorized)))
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
			ctx.JSON(handler.getError(exceptions.NewAuthError(errUnauthorized)))
			return
		}

		ctx.JSON(handler.getError(exceptions.NewInternalServerError(errSomethingWentWrong)))
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

		ctx.JSON(handler.getError(exceptions.NewAuthError(errUnauthorized)))
		return
	}

	var req CreateInviteCodeRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		handler.logger.Errorf("failed to bind request body: %v", err)

		ctx.JSON(handler.getError(exceptions.NewBadRequestError(errInvalidRequestBody)))
		return
	}

	if err := handler.validator.Struct(&req); err != nil {
		ctx.JSON(handler.getError(err))
		return
	}

	var inviteCode models.InviteCode

	err = handler.db.Client.
		Where("code = ?", req.Code).
		First(&inviteCode).
		Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		handler.logger.Errorf("failed to get invite code by code: %v err: %v", req.Code, err)

		ctx.JSON(handler.getError(exceptions.NewInternalServerError(errSomethingWentWrong)))
		return
	}

	if err == nil {
		ctx.JSON(handler.getError(exceptions.NewValidationError(
			[]exceptions.ValidationField{
				{Name: "code", Err: fmt.Errorf("invite code already exists")},
			},
		)))
		return
	}

	inviteCode = models.InviteCode{
		Code:      req.Code,
		MaxUses:   req.MaxUses,
		CreatedBy: userID,
	}

	err = handler.db.Client.
		Create(&inviteCode).
		Error

	if err != nil {
		handler.logger.Errorf("failed to create invite code: %v", err)

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

func (handler *UserHandler) deleteMyInviteCodeAction(ctx *gin.Context) {

	userID, err := handler.getUserID(ctx)
	if err != nil {
		handler.logger.Errorf("failed to get user id from context: %v", err)

		ctx.JSON(handler.getError(exceptions.NewAuthError(errUnauthorized)))
		return
	}

	var user models.User

	err = handler.db.Client.
		First(&user, userID).
		Error

	if err != nil {
		handler.logger.Errorf("failed to get user: %v", err)

		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(handler.getError(exceptions.NewAuthError(errUnauthorized)))
			return
		}

		ctx.JSON(handler.getError(exceptions.NewInternalServerError(errSomethingWentWrong)))
		return
	}

	rawCodeID := ctx.Param("code_id")

	codeID, err := uuid.Parse(rawCodeID)
	if err != nil {
		handler.logger.Errorf("failed to parse code id: %v", err)

		ctx.JSON(handler.getError(exceptions.NewBadRequestError(fmt.Errorf("invalid code_id path param"))))
		return
	}

	err = handler.db.Client.
		Where("id = ? ", codeID).
		Where("created_by = ?", userID).
		Delete(&models.InviteCode{}).
		Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		handler.logger.Errorf("failed to delete invite code: %v", err)

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

func (handler *UserHandler) getMySkillsProgressAction(ctx *gin.Context) {

	userID, err := handler.getUserID(ctx)
	if err != nil {
		handler.logger.Errorf("failed to get user id from context: %v", err)

		ctx.JSON(handler.getError(exceptions.NewAuthError(errUnauthorized)))
		return
	}

	// Find user with skills
	var user models.User

	err = handler.db.Client.
		First(&user, userID).
		Error

	if err != nil {
		handler.logger.Errorf("failed to get user: %v", err)

		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(handler.getError(exceptions.NewAuthError(errUnauthorized)))
			return
		}

		ctx.JSON(handler.getError(exceptions.NewInternalServerError(errSomethingWentWrong)))
		return
	}

	var userSkills []models.UserSkill

	err = handler.db.Client.
		Where("user_id = ?", userID).
		Order("total_xp DESC").
		Preload("BaseSkill").
		Preload("ParentSkill").
		Find(&userSkills).
		Error

	if err != nil {
		handler.logger.Errorf("failed to get user skills: %v", err)

		ctx.JSON(handler.getError(exceptions.NewInternalServerError(errSomethingWentWrong)))
		return
	}

	// Create subskills map
	subskillsMap := make(map[uuid.UUID][]models.UserSkill)

	for _, userSkill := range userSkills {
		// Skip root skills
		if userSkill.BaseSkillID.Valid {
			continue
		}

		subskillsMap[userSkill.ParentSkill.ID] = append(subskillsMap[userSkill.ParentSkill.ID], userSkill)
	}

	// Build skills tree
	rootSkills := make([]RootSkill, 0, len(subskillsMap))

	for _, userSkill := range userSkills {

		isRoot := userSkill.BaseSkillID.Valid

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
				ID:   subskill.ID,
				Name: subskill.DisplayName,
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
			ID:   userSkill.ID,
			Name: userSkill.DisplayName,
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

func (handler *UserHandler) createActivityAction(ctx *gin.Context) {

	// trying to get user id from context
	userID, err := handler.getUserID(ctx)
	if err != nil {
		handler.logger.Errorf("failed to get user id from context: %v", err)

		ctx.JSON(handler.getError(exceptions.NewAuthError(errUnauthorized)))
		return
	}

	// trying to find user in database
	var user models.User

	err = handler.db.Client.
		First(&user, userID).
		Error

	if err != nil {
		handler.logger.Errorf("failed to get user: %v", err)

		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(handler.getError(exceptions.NewAuthError(errUnauthorized)))
			return
		}

		ctx.JSON(handler.getError(exceptions.NewInternalServerError(errSomethingWentWrong)))
		return
	}

	// trying to bind request body to get mode
	var baseReq BaseCreateActivityRequest

	if err := ctx.ShouldBindBodyWithJSON(&baseReq); err != nil {
		handler.logger.Errorf("failed to bind request body: %v", err)

		ctx.JSON(handler.getError(exceptions.NewBadRequestError(errInvalidRequestBody)))
		return
	}

	// trying to validate mode field
	var dto CreateActivityDTO

	switch baseReq.Mode {
	case enums.ActivityCreateModes.Manually:
		dto, err = handler.getManualCreateActivityDTO(ctx)

	case enums.ActivityCreateModes.AI:
		dto, err = handler.getAiCreateActivityDTO(ctx, userID)

	default:
		ctx.JSON(handler.getError(exceptions.NewValidationError([]exceptions.ValidationField{
			{Name: "mode", Err: fmt.Errorf("invalid value")},
		})))
		return
	}

	if err != nil {
		handler.logger.Errorf("failed to get create activity dto: %v", err)

		ctx.JSON(handler.getError(err))
		return
	}

	err = handler.createActivityAndUpdateLeaderboardEntry(userID, dto)
	if err != nil {
		handler.logger.Errorf("failed to create activity: %v", err)

		ctx.JSON(handler.getError(err))
		return
	}

	ctx.JSON(
		http.StatusOK,
		SuccessDataResponse{
			Data: []struct{}{},
		},
	)
}

func (handler *UserHandler) getManualCreateActivityDTO(ctx *gin.Context) (CreateActivityDTO, error) {

	var req ManuallyCreatedActivityRequest

	if err := ctx.ShouldBindBodyWithJSON(&req); err != nil {
		handler.logger.Errorf("failed to bind request body: %v", err)

		return CreateActivityDTO{}, exceptions.NewBadRequestError(errInvalidRequestBody)
	}

	if err := handler.validator.Struct(&req); err != nil {
		return CreateActivityDTO{}, err
	}

	return CreateActivityDTO{
		Mode:               req.Mode,
		Description:        req.Description,
		ActivityCategoryID: req.ActivityCategoryID,
		HasImpact:          req.HasImpact,
		IsNew:              req.IsNew,
		IsHard:             req.IsHard,
		SkillWeights:       req.SkillWeights,
	}, nil
}

func (handler *UserHandler) getAiCreateActivityDTO(ctx *gin.Context, userID uuid.UUID) (CreateActivityDTO, error) {

	var req AICreatedActivityRequest

	if err := ctx.ShouldBindBodyWithJSON(&req); err != nil {
		handler.logger.Errorf("failed to bind request body: %v", err)

		return CreateActivityDTO{}, exceptions.NewBadRequestError(errInvalidRequestBody)
	}

	if err := handler.validator.Struct(&req); err != nil {
		return CreateActivityDTO{}, err
	}

	promptsMap, err := handler.getAiPromptsMap()
	if err != nil {
		handler.logger.Errorf("failed to get ai prompt map: %v", err)

		return CreateActivityDTO{}, err
	}

	var wg sync.WaitGroup

	wg.Add(2)

	var activityCategory models.ActivityCategory
	var activityEvaluation EvaluateActivityResponse
	var skillWeights []skill.SkillWeight

	go func() {
		defer wg.Done()

		// get activity category
		activityCategory, err = handler.getAiSelectedActivityCategory(req.Description, promptsMap[enums.AIPromptCodes.SelectActivityCategory])
		if err != nil {
			handler.logger.Errorf("failed to get ai selected activity category: %v", err)
			return
		}

		// get outstanding activities in this category
		activities, err := handler.getOutstandingActivitiesByCategory(userID, activityCategory.ID)
		if err != nil {
			handler.logger.Errorf("failed to get outstanding activities: %v", err)
			return
		}

		// get ai activity evaluation
		activityEvaluation, err = handler.getAiActivityEvaluation(req.Description, activities, promptsMap[enums.AIPromptCodes.EvaluateActivity])
		if err != nil {
			handler.logger.Errorf("failed to get ai evaluated activity: %v", err)
		}
	}()

	go func() {
		defer wg.Done()

		// get ai skill weights distribution
		skillWeights, err = handler.getAiSkillWeightsDistribution(userID, req.Description, promptsMap[enums.AIPromptCodes.SkillWeightsDistribution])

		if err != nil {
			handler.logger.Errorf("failed to get ai skill weights distribution: %v", err)
		}
	}()

	wg.Wait()
	if err != nil {
		return CreateActivityDTO{}, exceptions.NewInternalServerError(errSomethingWentWrong)
	}

	return CreateActivityDTO{
		Mode:               req.Mode,
		Description:        req.Description,
		ActivityCategoryID: activityCategory.ID,
		HasImpact:          activityEvaluation.HasImpact,
		IsNew:              activityEvaluation.IsNew,
		IsHard:             activityEvaluation.IsHard,
		SkillWeights:       skillWeights,
	}, nil
}

func (handler *UserHandler) getAiPromptsMap() (map[string]models.AiPrompt, error) {

	var prompts []models.AiPrompt

	promptCodes := []string{
		enums.AIPromptCodes.SelectActivityCategory,
		enums.AIPromptCodes.EvaluateActivity,
		enums.AIPromptCodes.SkillWeightsDistribution,
	}

	err := handler.db.Client.
		Where("code in (?)", promptCodes).
		Find(&prompts).
		Error

	if err != nil {
		handler.logger.Errorf("failed to get ai prompts: %v", err)

		return nil, exceptions.NewInternalServerError(errSomethingWentWrong)
	}

	if len(prompts) != len(promptCodes) {
		handler.logger.Errorf("some ai prompts are missing in database")

		return nil, exceptions.NewInternalServerError(errSomethingWentWrong)
	}

	promptsMap := make(map[string]models.AiPrompt)

	for _, prompt := range prompts {
		promptsMap[prompt.Code] = prompt
	}

	return promptsMap, nil
}

func (handler *UserHandler) getAiSelectedActivityCategory(description string, prompt models.AiPrompt) (models.ActivityCategory, error) {

	var activityCategories []models.ActivityCategory

	err := handler.db.Client.
		Find(&activityCategories).
		Error

	if err != nil {
		handler.logger.Errorf("failed to get activity categories: %v", err)

		return models.ActivityCategory{}, exceptions.NewInternalServerError(errSomethingWentWrong)
	}

	activityCategoryCodes := make([]string, 0, len(activityCategories))
	activityCategoryMap := make(map[string]models.ActivityCategory)

	for _, category := range activityCategories {
		activityCategoryCodes = append(activityCategoryCodes, category.Code)
		activityCategoryMap[category.Code] = category
	}

	jsonActivityCategories, err := json.Marshal(activityCategoryCodes)
	if err != nil {
		handler.logger.Errorf("failed to marshal activity categories: %v", err)

		return models.ActivityCategory{}, exceptions.NewInternalServerError(errSomethingWentWrong)
	}

	userMessage := fmt.Sprintf(
		"options: %s\ninput: %s",
		string(jsonActivityCategories), description,
	)

	response, err := handler.aiService.GetChatCompletion(
		prompt.SystemPrompt, userMessage,
	)

	var selectActivityCategoryResponse SelectActivityCategoryResponse

	err = json.Unmarshal([]byte(response), &selectActivityCategoryResponse)
	if err != nil {
		handler.logger.Errorf("failed to unmarshal ai response: %v", err)

		return models.ActivityCategory{}, exceptions.NewInternalServerError(errSomethingWentWrong)
	}

	activityCategory, exists := activityCategoryMap[selectActivityCategoryResponse.Code]
	if !exists {
		handler.logger.Errorf("ai selected invalid activity category code: %s", selectActivityCategoryResponse.Code)
		return models.ActivityCategory{}, exceptions.NewInternalServerError(errSomethingWentWrong)
	}

	return activityCategory, nil
}

func (handler *UserHandler) getAiActivityEvaluation(
	description string,
	activities []models.Activity,
	prompt models.AiPrompt,
) (EvaluateActivityResponse, error) {

	descriptions := make([]string, 0, len(activities))

	for _, activity := range activities {
		descriptions = append(descriptions, activity.Description)
	}

	jsonDescriptions, err := json.Marshal(descriptions)
	if err != nil {
		handler.logger.Errorf("failed to marshal activity descriptions: %v", err)

		return EvaluateActivityResponse{}, exceptions.NewInternalServerError(errSomethingWentWrong)
	}

	userMessage := fmt.Sprintf(
		"past activities: %s\ninput: %s",
		string(jsonDescriptions), description,
	)

	response, err := handler.aiService.GetChatCompletion(
		prompt.SystemPrompt, userMessage,
	)

	var evaluateActivityResponse EvaluateActivityResponse

	err = json.Unmarshal([]byte(response), &evaluateActivityResponse)
	if err != nil {
		handler.logger.Errorf("failed to unmarshal ai response: %v", err)

		return EvaluateActivityResponse{}, exceptions.NewInternalServerError(errSomethingWentWrong)
	}

	return evaluateActivityResponse, nil
}

func (handler *UserHandler) getOutstandingActivitiesByCategory(userID uuid.UUID, categoryID uuid.UUID) ([]models.Activity, error) {

	var activities []models.Activity

	err := handler.db.Client.
		Where("user_id = ?", userID).
		Where("activity_category_id = ?", categoryID).
		Order("has_impact DESC").
		Order("is_new DESC").
		Order("is_hard DESC").
		Order("created_at DESC").
		Order("id DESC").
		Limit(10).
		Find(&activities).
		Error

	return activities, err
}

func (handler *UserHandler) getAiSkillWeightsDistribution(userID uuid.UUID, description string, prompt models.AiPrompt) ([]skill.SkillWeight, error) {

	var userSkills []models.UserSkill

	err := handler.db.Client.
		Where("user_id = ?", userID).
		Order("display_name ASC").
		Find(&userSkills).
		Error

	if err != nil {
		handler.logger.Errorf("failed to get user skills: %v", err)

		return []skill.SkillWeight{}, exceptions.NewInternalServerError(errSomethingWentWrong)
	}

	skillCodes := make([]string, 0, len(userSkills))
	skillMap := make(map[string]models.UserSkill)

	for _, skill := range userSkills {
		skillCodes = append(skillCodes, skill.Code)
		skillMap[skill.Code] = skill
	}

	jsonSkillCodes, err := json.Marshal(skillCodes)
	if err != nil {
		handler.logger.Errorf("failed to marshal skill codes: %v", err)

		return []skill.SkillWeight{}, exceptions.NewInternalServerError(errSomethingWentWrong)
	}

	userMessage := fmt.Sprintf(
		"options: %s\ninput: %s",
		string(jsonSkillCodes), description,
	)

	response, err := handler.aiService.GetChatCompletion(
		prompt.SystemPrompt, userMessage,
	)

	var skillWeightsDistributionResponse SkillWeightsDistributionResponse

	err = json.Unmarshal([]byte(response), &skillWeightsDistributionResponse)
	if err != nil {
		handler.logger.Errorf("failed to unmarshal ai response: %v", err)

		return []skill.SkillWeight{}, exceptions.NewInternalServerError(errSomethingWentWrong)
	}

	skillWeights := make([]skill.SkillWeight, 0, len(skillWeightsDistributionResponse.Weights))

	for _, weight := range skillWeightsDistributionResponse.Weights {
		if _, exists := skillMap[weight.Code]; !exists {
			handler.logger.Errorf("ai selected invalid skill code: %s", weight.Code)

			return []skill.SkillWeight{}, exceptions.NewInternalServerError(errSomethingWentWrong)
		}

		skillWeights = append(skillWeights, skill.SkillWeight{
			SkillID: skillMap[weight.Code].ID,
			Weight:  weight.Weight,
		})
	}

	return skillWeights, nil
}

func (handler *UserHandler) createActivityAndUpdateLeaderboardEntry(userID uuid.UUID, dto CreateActivityDTO) error {

	activityReward, err := handler.skillService.CalcActivityReward(
		dto.SkillWeights,
		dto.HasImpact,
		dto.IsHard,
		dto.IsNew,
	)

	if err != nil {
		handler.logger.Errorf("failed to calc activity reward: %v", err)

		return exceptions.NewBadRequestError(err)
	}

	var skillIDs []uuid.UUID

	for _, skillWeight := range dto.SkillWeights {
		skillIDs = append(skillIDs, skillWeight.SkillID)
	}

	var userSkills []models.UserSkill

	err = handler.db.Client.
		Where("user_id = ? AND id IN (?)", userID, skillIDs).
		Find(&userSkills).
		Error

	if err != nil {
		handler.logger.Errorf("failed to get user skills: %v", err)

		return exceptions.NewInternalServerError(errSomethingWentWrong)
	}

	if len(userSkills) != len(skillIDs) {
		return exceptions.NewValidationError([]exceptions.ValidationField{
			{Name: "skill_weights", Err: fmt.Errorf("some skill_ids do not exist")},
		})
	}

	// Create activity and rewards in transaction
	tx := handler.db.Client.Begin()

	activity := models.Activity{
		Description:        dto.Description,
		UserID:             userID,
		Source:             enums.ActivitySources.Manual,
		ActivityCategoryID: dto.ActivityCategoryID,
		HasImpact:          dto.HasImpact,
		IsHard:             dto.IsHard,
		IsNew:              dto.IsNew,
	}

	err = tx.Create(&activity).Error
	if err != nil {
		handler.logger.Errorf("failed to create activity: %v", err)

		tx.Rollback()
		return exceptions.NewInternalServerError(errSomethingWentWrong)
	}

	rewards := make([]models.ActivityReward, 0, len(activityReward.SkillRewards))

	for _, skillReward := range activityReward.SkillRewards {
		rewards = append(rewards, models.ActivityReward{
			ActivityID:  activity.ID,
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
			return exceptions.NewInternalServerError(errSomethingWentWrong)
		}
	}

	err = tx.Create(&rewards).Error
	if err != nil {
		handler.logger.Errorf("failed to create activity rewards: %v", err)

		tx.Rollback()
		return exceptions.NewInternalServerError(errSomethingWentWrong)
	}

	// getting current leaderboard season
	var season models.LeaderboardSeason

	now := time.Now().UTC()

	err = tx.
		Where("period_start <= ?", now).
		Where("period_end >= ?", now).
		First(&season).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			tx.Commit()
			return nil
		}

		handler.logger.Errorf("failed to get current leaderboard season: %v", err)

		tx.Rollback()
		return exceptions.NewInternalServerError(errSomethingWentWrong)
	}

	// get or create leaderboard entry for user and season

	var entry models.LeaderboardEntry

	err = tx.
		Where("user_id = ?", userID).
		Where("season_id = ?", season.ID).
		First(&entry).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			entry = models.LeaderboardEntry{
				UserID:   userID,
				SeasonID: season.ID,
			}
		} else {
			handler.logger.Errorf("failed to get leaderboard entry: %v", err)

			tx.Rollback()
			return exceptions.NewInternalServerError(errSomethingWentWrong)
		}
	}

	// update leaderboard entry score
	entry.Score += activityReward.TotalXP

	err = tx.Save(&entry).Error
	if err != nil {
		handler.logger.Errorf("failed to update leaderboard entry: %v", err)

		tx.Rollback()
		return exceptions.NewInternalServerError(errSomethingWentWrong)
	}

	tx.Commit()

	return nil
}

func (handler *UserHandler) getMySkillsAction(ctx *gin.Context) {

	userID, err := handler.getUserID(ctx)
	if err != nil {
		handler.logger.Errorf("failed to get user id from context: %v", err)

		ctx.JSON(handler.getError(exceptions.NewAuthError(errUnauthorized)))
		return
	}

	var user models.User

	err = handler.db.Client.
		First(&user, userID).
		Error

	if err != nil {
		handler.logger.Errorf("failed to get user: %v", err)

		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(handler.getError(exceptions.NewAuthError(errUnauthorized)))
			return
		}

		ctx.JSON(handler.getError(exceptions.NewInternalServerError(errSomethingWentWrong)))
		return
	}

	var userSkills []models.UserSkill

	err = handler.db.Client.
		Where("user_id = ?", userID).
		Order("display_name ASC").
		Find(&userSkills).
		Error

	if err != nil {
		handler.logger.Errorf("failed to get user skills: %v", err)

		ctx.JSON(handler.getError(exceptions.NewInternalServerError(errSomethingWentWrong)))
		return
	}

	response := make([]GetSkillsResponseItem, 0, len(userSkills))

	for _, userSkill := range userSkills {
		response = append(response, GetSkillsResponseItem{
			ID:   userSkill.ID,
			Name: userSkill.DisplayName,
		})
	}

	ctx.JSON(
		http.StatusOK,
		SuccessDataResponse{
			Data: response,
		},
	)
}

func (handler *UserHandler) getMyActivitiesAction(ctx *gin.Context) {

	userID, err := handler.getUserID(ctx)
	if err != nil {
		handler.logger.Errorf("failed to get user id from context: %v", err)

		ctx.JSON(handler.getError(exceptions.NewAuthError(errUnauthorized)))
		return
	}

	var req GetActivitiesRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		handler.logger.Errorf("failed to bind query params: %v", err)

		ctx.JSON(handler.getError(exceptions.NewBadRequestError(errInvalidQueryParams)))
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
			ctx.JSON(handler.getError(exceptions.NewAuthError(errUnauthorized)))
			return
		}

		ctx.JSON(handler.getError(exceptions.NewInternalServerError(errSomethingWentWrong)))
		return
	}

	var total int64
	var activities []models.Activity

	query := handler.db.Client.
		Where("user_id = ?", userID).
		Order("created_at DESC")

	err = query.Model(&activities).Count(&total).Error
	if err != nil {
		handler.logger.Errorf("failed to count user activities: %v", err)

		ctx.JSON(handler.getError(exceptions.NewInternalServerError(errSomethingWentWrong)))
		return
	}

	offset := (req.Page - 1) * req.Limit

	err = query.
		Preload("ActivityCategory").
		Preload("Rewards").
		Preload("Rewards.UserSkill").
		Limit(req.Limit).
		Offset(offset).
		Find(&activities).
		Error

	if err != nil {
		handler.logger.Errorf("failed to get user activities: %v", err)

		ctx.JSON(handler.getError(exceptions.NewInternalServerError(errSomethingWentWrong)))
		return
	}

	responseItems := make([]GetActivitiesResponseItem, 0, len(activities))

	for _, activity := range activities {

		rewards := make([]ActivityReward, 0, len(activity.Rewards))

		for _, reward := range activity.Rewards {
			rewards = append(rewards, ActivityReward{
				SkillID:   reward.UserSkill.ID,
				SkillName: reward.UserSkill.DisplayName,
				XPAmount:  reward.XPAmount,
			})
		}

		responseItems = append(responseItems, GetActivitiesResponseItem{
			ID:               activity.ID,
			Description:      activity.Description,
			ActivityCategory: ActivityCategory{ID: activity.ActivityCategory.ID, Name: activity.ActivityCategory.DisplayName},
			HasImpact:        activity.HasImpact,
			IsNew:            activity.IsNew,
			IsHard:           activity.IsHard,
			Rewards:          rewards,
			CreatedAt:        activity.CreatedAt,
		})
	}

	response := GetActivitiesResponse{
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

func (handler *UserHandler) deleteMyActivityAction(ctx *gin.Context) {

	userID, err := handler.getUserID(ctx)

	if err != nil {
		handler.logger.Errorf("failed to get user id from context: %v", err)

		ctx.JSON(handler.getError(exceptions.NewAuthError(errUnauthorized)))
		return
	}

	rawActivityID := ctx.Param("activity_id")

	activityID, err := uuid.Parse(rawActivityID)
	if err != nil {
		handler.logger.Errorf("failed to parse activity id: %v", err)

		ctx.JSON(handler.getError(exceptions.NewBadRequestError(errInvalidRequestBody)))
		return
	}

	var activity models.Activity

	err = handler.db.Client.
		Where("id = ?", activityID).
		Where("user_id = ?", userID).
		Preload("Rewards").
		Preload("Rewards.UserSkill").
		First(&activity).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(
				http.StatusOK,
				SuccessDataResponse{
					Data: []struct{}{},
				},
			)
			return
		}

		handler.logger.Errorf("failed to get activity: %v", err)
		ctx.JSON(handler.getError(exceptions.NewInternalServerError(errSomethingWentWrong)))
		return
	}

	userSkillsToSave := make([]models.UserSkill, 0, len(activity.Rewards))

	var xpToSubtract int

	tx := handler.db.Client.Begin()

	for _, reward := range activity.Rewards {
		userSkill := reward.UserSkill
		userSkill.TotalXP -= reward.XPAmount

		xpToSubtract += reward.XPAmount

		userSkillsToSave = append(userSkillsToSave, userSkill)
	}

	if err := tx.Delete(&activity).Error; err != nil {
		handler.logger.Errorf("failed to delete activity: %v", err)

		tx.Rollback()
		ctx.JSON(handler.getError(exceptions.NewInternalServerError(errSomethingWentWrong)))
		return
	}

	if err := tx.Save(&userSkillsToSave).Error; err != nil {
		handler.logger.Errorf("failed to update user skills: %v", err)

		tx.Rollback()
		ctx.JSON(handler.getError(exceptions.NewInternalServerError(errSomethingWentWrong)))
		return
	}

	var season models.LeaderboardSeason

	err = tx.
		Where("period_start <= ?", activity.CreatedAt).
		Where("period_end >= ?", activity.CreatedAt).
		First(&season).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			tx.Commit()

			ctx.JSON(
				http.StatusOK,
				SuccessDataResponse{
					Data: []struct{}{},
				},
			)
			return
		}

		handler.logger.Errorf("failed to get current leaderboard season: %v", err)

		tx.Rollback()
		ctx.JSON(handler.getError(exceptions.NewInternalServerError(errSomethingWentWrong)))
		return
	}

	if season.PeriodEnd.Before(time.Now().UTC()) {
		tx.Commit()

		ctx.JSON(
			http.StatusOK,
			SuccessDataResponse{
				Data: []struct{}{},
			},
		)
		return
	}

	var entry models.LeaderboardEntry

	err = tx.
		Where("user_id = ?", userID).
		Where("season_id = ?", season.ID).
		First(&entry).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			tx.Commit()
			ctx.JSON(
				http.StatusOK,
				SuccessDataResponse{
					Data: []struct{}{},
				},
			)
			return
		}

		handler.logger.Errorf("failed to get leaderboard entry: %v", err)

		tx.Rollback()
		ctx.JSON(handler.getError(exceptions.NewInternalServerError(errSomethingWentWrong)))
		return
	}

	entry.Score -= xpToSubtract

	if err := tx.Save(&entry).Error; err != nil {
		handler.logger.Errorf("failed to update leaderboard entry: %v", err)

		tx.Rollback()
		ctx.JSON(handler.getError(exceptions.NewInternalServerError(errSomethingWentWrong)))
		return
	}

	tx.Commit()
	ctx.JSON(
		http.StatusOK,
		SuccessDataResponse{
			Data: []struct{}{},
		},
	)
}
