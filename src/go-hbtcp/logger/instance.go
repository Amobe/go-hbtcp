package logger

import (
	"log"
	"os"
	"sync"
)

var gOnce sync.Once
var gLoggerInstance *Logger

// createLoggerInstance generate an instance of the logger if it's nil
func createLoggerInstance() {
	if gLoggerInstance == nil {
		logger := new(Logger)
		writer := os.Stdout
		flag := log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile

		logger.err = log.New(writer, "[E] ", flag)
		logger.warn = log.New(writer, "[W] ", flag)
		logger.info = log.New(writer, "[I] ", flag)
		logger.debug = log.New(writer, "[D] ", flag)
		logger.SetLevel(LevelDebug)
		logger.depth = 2

		gLoggerInstance = logger
	}
}

// GetLoggerInstance return a global instance of the logger object
func GetLoggerInstance() *Logger {
	if gLoggerInstance == nil {
		gOnce.Do(createLoggerInstance)
	}
	return gLoggerInstance
}
