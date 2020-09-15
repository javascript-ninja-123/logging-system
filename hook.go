package logger2

type Hook func(level Level, msg string)

func (h Hook) Send(level Level, msg string) {
	h(level, msg)
}
