package logger

import (
	"redisrenamer/pkg/model"

	"github.com/mborders/logmatic"
)

var log *logmatic.Logger

// Init initialises the Log object and return it
func Init(argConfig model.ArgConfig) *logmatic.Logger {
	if log == nil {
		var level logmatic.LogLevel = logmatic.INFO

		if argConfig.Verbose {
			level = logmatic.TRACE
		}

		log = logmatic.NewLogger()
		log.SetLevel(level)
		log.ExitOnFatal = true
	}

	return log
}

// Trace logs a trace statement
func Trace(format string, a ...interface{}) {
	if log != nil {
		log.Trace(format, a...)
	}
}

// Debug logs a debug statement
func Debug(format string, a ...interface{}) {
	if log != nil {
		log.Debug(format, a...)
	}
}

// Info logs a info statement
func Info(format string, a ...interface{}) {
	if log != nil {
		log.Info(format, a...)
	}
}

// Warn logs a warn statement
func Warn(format string, a ...interface{}) {
	if log != nil {
		log.Warn(format, a...)
	}
}

// Error logs a error statement
func Error(format string, a ...interface{}) {
	if log != nil {
		log.Error(format, a...)
	}
}

// Fatal logs a fatal statement
func Fatal(format string, a ...interface{}) {
	if log != nil {
		log.Fatal(format, a...)
	}
}
