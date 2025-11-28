package logger

import (
	"voicekit-mock/config"

	"go.uber.org/zap"
)


type Logger interface {

}

type logger struct {
	logger zap.SugaredLogger	
}

func NewLogger(cfg config.LoggerConfig) *Logger {
	return nil
}