package http

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sunsetsavorer/grind/internal/ai"
	"github.com/sunsetsavorer/grind/internal/config"
	"github.com/sunsetsavorer/grind/internal/db"
	"github.com/sunsetsavorer/grind/internal/exceptions"
	"github.com/sunsetsavorer/grind/internal/jwt"
	"github.com/sunsetsavorer/grind/internal/logger"
	"github.com/sunsetsavorer/grind/internal/skill"
	"github.com/sunsetsavorer/grind/internal/validator"
)

type BaseHandler struct {
	config       *config.Config
	db           *db.DB
	jwt          *jwt.JWT
	validator    *validator.Validator
	logger       *logger.Logger
	skillService *skill.SkillService
	aiService    *ai.AIService
}

func NewBaseHandler(
	config *config.Config,
	db *db.DB,
	jwt *jwt.JWT,
	validator *validator.Validator,
	logger *logger.Logger,
	skillService *skill.SkillService,
	aiService *ai.AIService,
) *BaseHandler {

	return &BaseHandler{
		config:       config,
		db:           db,
		jwt:          jwt,
		validator:    validator,
		logger:       logger,
		skillService: skillService,
		aiService:    aiService,
	}
}

func (h *BaseHandler) getUserID(ctx *gin.Context) (uuid.UUID, error) {

	userIDStr, exist := ctx.Get("user_id")
	if !exist {
		return uuid.Nil, fmt.Errorf("failed to get user id from context")
	}

	userID, err := uuid.Parse(fmt.Sprint(userIDStr))
	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to parse uuid from user context: %v", err)
	}

	return userID, nil
}

func (h *BaseHandler) getError(err error) (int, any) {

	if _, ok := err.(*exceptions.AuthError); ok {
		return http.StatusUnauthorized, ErrorResp[OtherError]{
			Errors: OtherError{err.Error()},
		}
	}

	if _, ok := err.(*exceptions.ValidationError); ok {
		return http.StatusUnprocessableEntity, ErrorResp[map[string]string]{
			Errors: err.(*exceptions.ValidationError).Errors(),
		}
	}

	if _, ok := err.(*exceptions.ManyRequestsError); ok {
		return http.StatusTooManyRequests, ErrorResp[OtherError]{
			Errors: OtherError{err.Error()},
		}
	}

	if _, ok := err.(*exceptions.NotFoundError); ok {
		return http.StatusNotFound, ErrorResp[OtherError]{
			Errors: OtherError{err.Error()},
		}
	}

	if _, ok := err.(*exceptions.BadRequestError); ok {
		return http.StatusBadRequest, ErrorResp[OtherError]{
			Errors: OtherError{err.Error()},
		}
	}

	if _, ok := err.(*exceptions.InternalServerError); ok {
		return http.StatusInternalServerError, ErrorResp[OtherError]{
			Errors: OtherError{err.Error()},
		}
	}

	return http.StatusInternalServerError, ErrorResp[OtherError]{
		Errors: OtherError{errUnknownError.Error()},
	}
}
