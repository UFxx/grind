package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	DSN              string
	AppAddress       string
	JWTSecret        string
	TelegramBotToken string
	AIKey            string
}

func New() *Config {

	return &Config{}
}

func (config *Config) Load() error {

	viper.SetConfigFile("internal/config/.env")

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read config: %v", err)
	}

	if err := viper.Unmarshal(config); err != nil {
		return fmt.Errorf("failed to unmarshal config: %v", err)
	}

	return nil
}
