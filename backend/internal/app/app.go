package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sunsetsavorer/grind/internal/config"
	"github.com/sunsetsavorer/grind/internal/db"
	"github.com/sunsetsavorer/grind/internal/jwt"
	"github.com/sunsetsavorer/grind/internal/transport/http"
	"github.com/sunsetsavorer/grind/internal/validator"
)

type App struct {
}

func New() *App {

	return &App{}
}

func (a *App) Run() error {

	config := config.New()

	if err := config.Load(); err != nil {
		return fmt.Errorf("failed to load config: %v", err)
	}

	db, err := db.New(config.DSN)
	if err != nil {
		return fmt.Errorf("failed to open db connection: %v", err)
	}

	jwt := jwt.New(config.JWTSecret, config.JwtLifetimeSeconds)

	validator := validator.New()

	baseHandler := http.NewBaseHandler(
		config,
		db,
		jwt,
		validator,
	)

	router := gin.Default()

	apiGroup := router.Group("/api")
	{
		authHandler := http.NewAuthHandler(baseHandler)
		authHandler.RegisterRoutes(apiGroup)
	}

	if err := router.Run(config.AppAddress); err != nil {
		return fmt.Errorf("failed to run server: %v", err)
	}

	return nil
}
