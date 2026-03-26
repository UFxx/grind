package cron

import (
	"fmt"

	syscron "github.com/robfig/cron"
	"github.com/sunsetsavorer/grind/internal/db"
	"github.com/sunsetsavorer/grind/internal/logger"
)

type Cron struct {
	*syscron.Cron
	db     *db.DB
	logger *logger.Logger
}

func New(
	db *db.DB,
	logger *logger.Logger,
) *Cron {

	return &Cron{
		Cron:   syscron.New(),
		db:     db,
		logger: logger,
	}
}

func (cron *Cron) initJobs() error {

	err := cron.AddFunc("0 0 0 * * 1", cron.createNewLeaderboardSeason) // Every Monday at midnight
	if err != nil {
		return fmt.Errorf("failed to add `create new leaderboard season` job: %v", err)
	}

	return nil
}

func (cron *Cron) Start() error {

	err := cron.initJobs()
	if err != nil {
		return err
	}

	cron.Cron.Start()

	return nil
}

func (cron *Cron) Stop() {

	cron.Cron.Stop()
}
