package common

import log "github.com/sirupsen/logrus"

type Logger struct {
	*log.Entry
}

func NewLogger() Logger {
	return Logger{log.WithField("type", "grpc")}
}
