package logger

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	*logrus.Logger
	logFile *os.File
}

func New() (*Logger, error) {

	log := logrus.New()

	err := os.MkdirAll("logs", 0o755)
	if err != nil {
		return nil, fmt.Errorf("failed to create logs directory: %v", err)
	}

	file, err := os.OpenFile("logs/app.jsonl", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, fmt.Errorf("Failed to open log file: %v", err)
	}

	log.SetOutput(file)

	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetLevel(logrus.DebugLevel)

	return &Logger{
		Logger:  log,
		logFile: file,
	}, nil
}

func (logger *Logger) Close() {
	logger.logFile.Close()
}
