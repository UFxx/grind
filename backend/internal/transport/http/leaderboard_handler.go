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
		leaderboards.GET("/seasons/:season_id", handler.getSeasonDetailAction)
	}
}

func (handler *LeaderboardHandler) getSeasonsAction(ctx *gin.Context) {

	userID, err := handler.getUserID(ctx)
	if err != nil {
		handler.logger.Errorf("failed to get user from context: %v", err)

		ctx.JSON(handler.getError(exceptions.NewAuthError(errUnauthorized)))
		return
	}

	var user models.User

	err = handler.db.Client.First(&user, userID).Error
	if err != nil {
		handler.logger.Errorf("failed to find user by id: %v", err)

		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(handler.getError(exceptions.NewAuthError(errUnauthorized)))
			return
		}

		ctx.JSON(handler.getError(exceptions.NewInternalServerError(errSomethingWentWrong)))
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

func (handler *LeaderboardHandler) getSeasonDetailAction(ctx *gin.Context) {

	userID, err := handler.getUserID(ctx)
	if err != nil {
		handler.logger.Errorf("failed to get user from context: %v", err)

		ctx.JSON(handler.getError(exceptions.NewAuthError(errUnauthorized)))
		return
	}

	var user models.User

	err = handler.db.Client.First(&user, userID).Error
	if err != nil {
		handler.logger.Errorf("failed to find user by id: %v", err)

		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(handler.getError(exceptions.NewAuthError(errUnauthorized)))
			return
		}

		ctx.JSON(handler.getError(exceptions.NewInternalServerError(errSomethingWentWrong)))
		return
	}

	rawSeasonID := ctx.Param("season_id")

	seasonID, err := uuid.Parse(rawSeasonID)
	if err != nil {
		handler.logger.Errorf("failed to parse season id: %v", err)

		ctx.JSON(handler.getError(exceptions.NewBadRequestError(fmt.Errorf("invalid season_id path param"))))
		return
	}

	var season models.LeaderboardSeason

	err = handler.db.Client.First(&season, "id = ?", seasonID).Error

	if err != nil {
		handler.logger.Errorf("failed to get current season: %v", err)

		if ok := errors.Is(err, gorm.ErrRecordNotFound); ok {
			ctx.JSON(handler.getError(exceptions.NewNotFoundError(fmt.Errorf("season not found"))))
			return
		}

		ctx.JSON(handler.getError(exceptions.NewInternalServerError(errSomethingWentWrong)))
		return
	}

	var entries []models.LeaderboardEntry

	err = handler.db.Client.
		Where("season_id = ?", season.ID).
		Order("score DESC").
		Order("id ASC").
		Preload("User").
		Find(&entries).
		Error

	if err != nil {
		handler.logger.Errorf("failed to get season entries: %v", err)

		ctx.JSON(handler.getError(exceptions.NewInternalServerError(errSomethingWentWrong)))
		return
	}

	responseEntries := make([]LeaderboardEntry, 0, len(entries))

	cta := "CREATE ONE ACTIVITY TO GET ON THE LEADERBOARD!"

	myEntry := MyLeaderboardEntry{
		Name:      user.Name,
		AvatarURL: user.AvatarURL,
	}

	for i, entry := range entries {
		position := i + 1

		responseEntries = append(responseEntries, LeaderboardEntry{
			Position:  position,
			Name:      entry.User.Name,
			AvatarURL: entry.User.AvatarURL,
			Score:     entry.Score,
		})

		if entry.UserID == user.ID {
			if position == 1 {
				cta = "YOU ARE THE LEADER! CREATE MORE ACTIVITIES TO STAY ON TOP!"
			} else {
				cta = fmt.Sprintf("EARN %d MORE POINTS TO GET TO #%d!", entries[position-2].Score-entry.Score+1, position-1)
			}

			myEntry = MyLeaderboardEntry{
				Position:        &position,
				Name:            entry.User.Name,
				AvatarURL:       entry.User.AvatarURL,
				Score:           &entry.Score,
				IsInLeaderboard: true,
			}
		}
	}

	ctx.JSON(
		http.StatusOK,
		SuccessDataResponse{
			Data: LeaderboardSeasonDetailResponse{
				Season: LeaderboardSeasonResponseItem{
					ID:          season.ID,
					Name:        season.Name,
					PeriodStart: season.PeriodStart,
					PeriodEnd:   season.PeriodEnd,
				},
				CTA:     cta,
				Entries: responseEntries,
				Me:      myEntry,
			},
		},
	)
}
