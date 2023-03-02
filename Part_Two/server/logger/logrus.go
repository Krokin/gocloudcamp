package logger

import (
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
)

type Logger struct {
	Log       *logrus.Entry
	LogConfig func(code codes.Code) logrus.Level
}

func NewLogger() *Logger {
	return &Logger{Log: logrus.NewEntry(logrus.New()),
		LogConfig: func(code codes.Code) logrus.Level {
			if code == codes.OK {
				return logrus.InfoLevel
			}
			return logrus.ErrorLevel
		}}
}
