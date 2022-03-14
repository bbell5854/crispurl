package logger

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func Setup(level string) {
	lvl := converStringLevelToInt(level)

	log.SetLevel(lvl)
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
}

func converStringLevelToInt(level string) log.Level {
	switch level {
	case "trace":
		return log.TraceLevel
	case "debug":
		return log.DebugLevel
	case "info":
		return log.InfoLevel
	case "warn":
		return log.WarnLevel
	case "error":
		return log.ErrorLevel
	case "fatal":
		return log.FatalLevel
	case "panic":
		return log.PanicLevel
	}

	// Default log level
	return log.WarnLevel
}
