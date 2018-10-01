package log

import (
	"fmt"
	"os"
)

type Logger struct{}

const (
	LevelInfo  = "INFO"
	LevelDebug = "DEBUG"
	LevelError = "ERROR"
	LevelFatal = "FATAL"
)

func New() *Logger {
	return &Logger{}
}

func (l *Logger) Info(f string, ff ...interface{}) {
	l.out(LevelInfo, fmt.Sprintf(f+"\n", ff...))
}

func (l *Logger) Debug(f string, ff ...interface{}) {
	l.out(LevelDebug, fmt.Sprintf(f+"\n", ff...))
}

func (l *Logger) Error(f string, ff ...interface{}) {
	l.out(LevelError, fmt.Sprintf(f+"\n", ff...))
}

func (l *Logger) Fatal(f string, ff ...interface{}) {
	l.out(LevelFatal, fmt.Sprintf(f+"\n", ff...))
	os.Exit(1)
}

func (l *Logger) out(level, s string) {
	fmt.Printf("[%s] %s", level, s)
}
