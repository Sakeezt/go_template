package logs

import "fmt"

type Level uint8

const (
	PRINT Level = iota
	DEBUG
	INFO
	WARN
	ERROR
	FATAL
	SYSTEM
	PANIC
	OFF = 99
)

func (lev Level) string() string {
	switch lev {
	case PRINT:
		return "PRINT"
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARN:
		return "WARN"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	case SYSTEM:
		return "SYSTEM"
	case PANIC:
		return "PANIC"
	case OFF:
		return "OFF"
	default:
		return fmt.Sprintf("LEVEL(%d)", lev)
	}
}

func (lev Level) shortString() string {
	switch lev {
	case PRINT:
		return "PNT"
	case DEBUG:
		return "DEB"
	case INFO:
		return "INF"
	case WARN:
		return "WAR"
	case ERROR:
		return "ERR"
	case FATAL:
		return "FAT"
	case SYSTEM:
		return "SYS"
	case PANIC:
		return "PAN"
	default:
		return "---"
	}
}
