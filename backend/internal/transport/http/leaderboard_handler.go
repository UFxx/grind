package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sunsetsavorer/grind/internal/exceptions"
	"github.com/sunsetsavorer/grind/internal/models"
)

type LeaderboardHandler struct {
	*BaseHandler
}

func NewLeaderboardHandler(baseHandler *BaseHandler) *LeaderboardHandler {

	return &LeaderboardHandler{BaseHandler: baseHandler}
}

func (handler *LeaderboardHandler) RegisterRoutes(router *gin.RouterGroup) {

	leaderboards := router.Group("/leaderboards", handler.jwt.GinJWTAuthMiddleware())
	{
		leaderboards.GET("/seasons", handler.getSeasonsAction)
	}
}

func (handler *LeaderboardHandler) getSeasonsAction(ctx *gin.Context) {

	_, err := handler.getUserID(ctx)
	if err != nil {
		handler.logger.Errorf("failed to get user from context: %v", err)

		ctx.JSON(handler.getError(exceptions.NewAuthError(errUnauthorized)))
		return
	}

	var seasons []models.LeaderboardSeason

	err = handler.db.Client.
		Order("period_end DESC").
		Find(&seasons).
		Error

	if err != nil {
		handler.logger.Errorf("failed to get seasons: %v", err)

		ctx.JSON(handler.getError(exceptions.NewInternalServerError(errSomethingWentWrong)))
		return
	}

	response := make([]LeaderboardSeasonResponseItem, 0, len(seasons))

	for _, season := range seasons {
		response = append(response, LeaderboardSeasonResponseItem{
			ID:          season.ID,
			Name:        season.Name,
			PeriodStart: season.PeriodStart,
			PeriodEnd:   season.PeriodEnd,
		})
	}

	ctx.JSON(
		http.StatusOK,
		SuccessDataResponse{
			Data: response,
		},
	)
}
