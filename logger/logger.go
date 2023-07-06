package logger

import (
	"github.com/google/wire"
	"log"
)

type Logger interface {
	Log(message string)
}

type GoLogger struct {
}

func ProvideLogger() *GoLogger {
	return &GoLogger{}
}

func (l *GoLogger) Log(message string) {
	log.Println(message)
}

var LoggerSet = wire.NewSet(
	ProvideLogger,
	wire.Bind(new(Logger), new(*GoLogger)),
)
