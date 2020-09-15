package logger2

import (
	"fmt"
	"os"
	"time"
)

type LoggerInner interface {
	WithField(key string, value interface{}) LoggerInner
	WithFields(fields Fields) LoggerInner
	Debug(args ...interface{})
	Info(args ...interface{})
	Error(args ...interface{})
	Warning(args ...interface{})
	Panic(args ...interface{})
}

type Entity struct {
	Logger *Logger
	Data   Fields
}

func NewEntity(logger *Logger) *Entity {
	return &Entity{Logger: logger,
		Data: make(map[string]interface{})}
}

func (e *Entity) print() {
	bt, err := e.Logger.Formatter.Format(e)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	os.Stderr.WriteString(string(bt))

}

func (e *Entity) log(levelInt64 Level, msg string) {
	e.Data["time"] = time.Now().String()
	e.Data["level"] = levelInt64.String()
	e.Data["msg"] = msg
	e.print()
	e.Logger.Hook.Send(levelInt64, msg)
}

func (e *Entity) WithField(key string, value interface{}) LoggerInner {
	e.Data[key] = value
	return e
}

func (e *Entity) WithFields(fields Fields) LoggerInner {
	for key, value := range fields {
		e.WithField(key, value)
	}
	return e
}

func (e *Entity) Debug(args ...interface{}) {
	if e.Logger.Level >= DebugLevel {
		e.log(DebugLevel, fmt.Sprint(args...))
	}
}

func (e *Entity) Info(args ...interface{}) {
	if e.Logger.Level >= InfoLevel {
		e.log(InfoLevel, fmt.Sprint(args...))
	}
}

func (e *Entity) Error(args ...interface{}) {
	if e.Logger.Level >= ErrorLevel {
		e.log(ErrorLevel, fmt.Sprint(args...))
	}
}

func (e *Entity) Warning(args ...interface{}) {
	if e.Logger.Level >= WarnLevel {
		e.log(WarnLevel, fmt.Sprint(args...))
	}
}

func (e *Entity) Panic(args ...interface{}) {
	if e.Logger.Level >= PanicLevel {
		e.log(PanicLevel, fmt.Sprint(args...))
	}
}
