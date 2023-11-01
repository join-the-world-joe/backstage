package logger

type Level int8

const (
	Debug Level = iota
	Info
	Warn
	Error
)

const (
	DebugLevel = "Debug"
	InfoLevel  = "Info"
	WarnLevel  = "Warn"
	ErrorLevel = "Error"
)

type Logger interface {
	Name() string
	Log(level Level, v ...interface{})
	Logf(level Level, format string, v ...interface{})
	Destroy()
}
