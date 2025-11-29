package logger

import (
	"go.uber.org/zap"
)

type Interface interface {
	Info(msg string, args ...any)
	Warn(msg string, args ...any)
	Error(msg string, args ...any)
	Errorf(msg string, args ...any)
	Fatal(msg string, args ...any)
	Sync()
}

type Logger struct {
	logger *zap.SugaredLogger
}

func New() (*Logger, error) {
	logger, err := zap.NewDevelopmentConfig().Build()
	if err != nil {
		return nil, err
	}
	return &Logger{logger.Sugar()}, nil
}

func (l *Logger) Info(msg string, args ...any) {
	if len(args) > 0 {
		l.logger.Infow(msg, args...)
	} else {
		l.logger.Info(msg)
	}
}

func (l *Logger) Warn(msg string, args ...any) {
	if len(args) > 0 {
		l.logger.Warnw(msg, args...)
	} else {
		l.logger.Warn(msg)
	}
}

func (l *Logger) Error(msg string, args ...any) {
	if len(args) > 0 {
		l.logger.Errorw(msg, args...)
	} else {
		l.logger.Error(msg)
	}
}

func (l *Logger) Fatal(msg string, args ...any) {
	if len(args) > 0 {
		l.logger.Fatalw(msg, args...)
	} else {
		l.logger.Fatal(msg)
	}
}

func (l *Logger) Sync() error {
	return l.logger.Sync()
}