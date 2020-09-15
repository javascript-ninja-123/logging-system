package logger2

import "fmt"

// Level will be exported
type Level int64

// DebugLevel will be exported
const (
	DebugLevel Level = iota
	WarnLevel
	InfoLevel
	PanicLevel
	ErrorLevel
)

var cannotStrinifyErrorLevel = "canonot stringify"

func (l Level) String() string {
	switch l {
	case DebugLevel:
		return "debug"
	case WarnLevel:
		return "warn"
	case PanicLevel:
		return "panic"
	case ErrorLevel:
		return "error"
	case InfoLevel:
		return "info"
	default:
		return fmt.Sprintf("level %v", cannotStrinifyErrorLevel)
	}
}
