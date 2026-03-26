package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sunsetsavorer/grind/internal/ai"
	"github.com/sunsetsavorer/grind/internal/config"
	"github.com/sunsetsavorer/grind/internal/cron"
	"github.com/sunsetsavorer/grind/internal/db"
	"github.com/sunsetsavorer/grind/internal/jwt"
	"github.com/sunsetsavorer/grind/internal/logger"
	"github.com/sunsetsavorer/grind/internal/skill"
	"github.com/sunsetsavorer/grind/internal/transport/http"
	"github.com/sunsetsavorer/grind/internal/validator"
)

type App struct {
}

func New() *App {

	return &App{}
}

func (app *App) Run() error {

	config := config.New()

	if err := config.Load(); err != nil {
		return fmt.Errorf("failed to load config: %v", err)
	}

	db, err := db.New(config.Database.ConnectionString)
	if err != nil {
		return fmt.Errorf("failed to open db connection: %v", err)
	}

	jwt := jwt.New(
		config.JWT.Secret,
		config.JWT.LifetimeSeconds,
	)

	validator := validator.New()

	aiService := ai.NewAIService(
		config.AI.Model,
		config.AI.AccessKey,
		config.AI.BaseURL,
	)

	logger, err := logger.New()
	if err != nil {
		return fmt.Errorf("failed to initialize logger: %v", err)
	}
	defer logger.Close()

	skillService := skill.NewSkillService(config.Skill)

	baseHandler := http.NewBaseHandler(
		config,
		db,
		jwt,
		validator,
		logger,
		skillService,
		aiService,
	)

	router := gin.Default()

	apiGroup := router.Group("/api")
	{
		authHandler := http.NewAuthHandler(baseHandler)
		authHandler.RegisterRoutes(apiGroup)

		userHandler := http.NewUserHandler(baseHandler)
		userHandler.RegisterRoutes(apiGroup)

		activityHandler := http.NewActivityHandler(baseHandler)
		activityHandler.RegisterRoutes(apiGroup)

		appHandler := http.NewAppHandler(baseHandler)
		appHandler.RegisterRoutes(apiGroup)

		leaderboardHandler := http.NewLeaderboardHandler(baseHandler)
		leaderboardHandler.RegisterRoutes(apiGroup)
	}

	cron := cron.New(db, logger)

	if err := cron.Start(); err != nil {
		return fmt.Errorf("failed to start cron: %v", err)
	}
	defer cron.Stop()

	if err := router.Run(":" + config.App.Port); err != nil {
		return fmt.Errorf("failed to run server: %v", err)
	}

	return nil
}
