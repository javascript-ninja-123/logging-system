package logger2

import "sync"

type Logger struct {
	Level     Level
	Formatter Formatter
	mu        sync.Mutex
	Hook      Hook
}

func NewLogger() *Logger {
	return &Logger{
		Level:     InfoLevel,
		Formatter: &JSON_Formatter{},
	}
}

func (l *Logger) WithField(key string, value interface{}) LoggerInner {
	return NewEntity(l).WithField(key, value)
}

func (l *Logger) WithFields(fields Fields) LoggerInner {
	return NewEntity(l).WithFields(fields)
}
