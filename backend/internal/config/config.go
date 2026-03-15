package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type SkillConfig struct {
	StartLevelCost        float64
	BaseXP                float64
	AdditionalCoefficient float64
	ImpactMultiplier      float64
	NewMultiplier         float64
	HardMultiplier        float64
}

type Config struct {
	Database struct {
		ConnectionString string
	}
	App struct {
		Port         string
		Debug        bool
		AllowOrigins []string
	}
	Telegram struct {
		BotToken string
	}
	AI struct {
		AccessKey string
	}
	JWT struct {
		Secret          string
		LifetimeSeconds int64
	}
	Skill SkillConfig
}

func New() *Config {

	return &Config{}
}

func (config *Config) Load() error {

	viper.SetConfigFile("config.yml")

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read config: %v", err)
	}

	if err := viper.Unmarshal(config); err != nil {
		return fmt.Errorf("failed to unmarshal config: %v", err)
	}

	return nil
}
