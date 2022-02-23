package logger

import "log"

type Logger struct {
	logger   *log.Logger
	logLevel uint8
}

func (l *Logger) Debug(v ...interface{}) {
	if l.logLevel >= LevelDebug {
		l.logger.Println(v...)
	}
}

func (l *Logger) Info(v ...interface{}) {
	if l.logLevel >= LevelInfo {
		l.logger.Println(v...)
	}
}

func (l *Logger) Warn(v ...interface{}) {
	if l.logLevel >= LevelWarn {
		l.logger.Println(v...)
	}
}

func (l *Logger) Error(v ...interface{}) {
	if l.logLevel >= LevelError {
		l.logger.Println(v...)
	}
}

func (l *Logger) Fatal(v ...interface{}) {
	l.logger.Fatal(v...)
}
