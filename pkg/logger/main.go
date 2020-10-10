package logger

import (
	log "github.com/sirupsen/logrus"
)

func Error(message string) {
	log.Error(message)
}

func Info(message string) {
	log.Info(message)
}
