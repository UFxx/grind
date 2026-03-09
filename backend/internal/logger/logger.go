package logger

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	*logrus.Logger
}

func New() (*Logger, error) {

	log := logrus.New()

	file, err := os.OpenFile("logs/app.jsonl", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, fmt.Errorf("Failed to open log file: %v", err)
	}
	defer file.Close()

	log.SetOutput(file)

	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetLevel(logrus.DebugLevel)

	return &Logger{log}, nil
}
