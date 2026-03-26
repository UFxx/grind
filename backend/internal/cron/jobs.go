package cron

import (
	"fmt"
	"strings"
	"time"

	"github.com/sunsetsavorer/grind/internal/models"
)

func (cron *Cron) createNewLeaderboardSeason() {

	cron.logger.Info("Creating new leaderboard season...")

	now := time.Now().UTC()

	day := now.Day()
	year := now.Year()

	seasonNumber := day / 7
	monthName := strings.ToUpper(now.Month().String())
	seasonName := fmt.Sprintf("%s %d S%d", monthName, year, seasonNumber)

	endOfSeason := now.Add(time.Hour * 24 * 7) // 1 week

	season := models.LeaderboardSeason{
		Name:        seasonName,
		PeriodStart: &now,
		PeriodEnd:   &endOfSeason,
	}

	err := cron.db.Client.Create(&season).Error
	if err != nil {
		cron.logger.Errorf("failed to create new leaderboard season: %v", err)
		return
	}

	cron.logger.Infof("New leaderboard season created: %+v", season)
}
