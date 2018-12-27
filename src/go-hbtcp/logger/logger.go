package logger

import (
	"fmt"
	"log"
)

// Level represent the type of the logging level
type Level int

// List of the logging level
const (
	LevelError Level = iota
	LevelWarn
	LevelInfo
	LevelDebug
)

// Logger represent an active logging object which writes lines to the native logger
type Logger struct {
	level Level
	err   *log.Logger
	warn  *log.Logger
	info  *log.Logger
	debug *log.Logger
	depth int
}

// SetLevel change the current logging level to the given level
func (l *Logger) SetLevel(level Level) {
	l.level = level
}

// Error write the line to the error output channel
func (l *Logger) Error(format string, v ...interface{}) {
	if LevelError <= l.level {
		l.err.Output(l.depth, fmt.Sprintf(format, v...))
	}
}

// Warn write the line to the warning output channel
func (l *Logger) Warn(format string, v ...interface{}) {
	if LevelWarn <= l.level {
		l.warn.Output(l.depth, fmt.Sprintf(format, v...))
	}
}

// Info write the line to the information output channel
func (l *Logger) Info(format string, v ...interface{}) {
	if LevelInfo <= l.level {
		l.info.Output(l.depth, fmt.Sprintf(format, v...))
	}
}

// Debug write the line to the debug output channel
func (l *Logger) Debug(format string, v ...interface{}) {
	if LevelDebug <= l.level {
		l.debug.Output(l.depth, fmt.Sprintf(format, v...))
	}
}
