package logger

import "log"

const (
	LevelOff = iota
	LevelError
	LevelWarn
	LevelInfo
	LevelDebug
	LevelVerbose
)

func New(logLevel uint8) *Logger {
	return &Logger{
		log.New(log.Writer(), log.Prefix(), log.Flags()),
		logLevel,
	}
}
