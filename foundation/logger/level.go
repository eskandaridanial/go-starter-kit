package logger

// type 'Level' represents a logging level
type Level int

// constants 'Debug', 'Info', 'Warn', 'Error' are logging levels
const (
	Debug Level = iota
	Info
	Warn
	Error
)

// function 'String' returns the string representation of the logging level
func (l Level) String() string {
	switch l {
	case Debug:
		return "debug"
	case Info:
		return "info"
	case Warn:
		return "warn"
	case Error:
		return "error"
	default:
		return "unknown"
	}
}
