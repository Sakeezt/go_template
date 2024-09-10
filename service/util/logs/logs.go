package logs

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"go-template/config"

	"github.com/gookit/color"
	"github.com/uniplaces/carbon"
)

type Logs struct {
	config *config.Config
}

func New(appConfig *config.Config) Log {
	return &Logs{appConfig}
}

func (l *Logs) Println(v ...interface{}) {
	l.printLog(PRINT, v)
}

func (l *Logs) Printf(format string, v ...interface{}) {
	lev := PRINT
	t := l.stampTime()
	msg := fmt.Sprintf("[%s] ", l.consoleColor(lev).Render(lev.shortString()))
	msg += fmt.Sprintf(format, v...)
	fmt.Printf("%s %s\n", t, msg)
}

func (l *Logs) Debug(v ...interface{}) {
	l.printLog(DEBUG, v)
}

func (l *Logs) Info(v ...interface{}) {
	l.printLog(INFO, v)
}

func (l *Logs) Warning(v ...interface{}) {
	l.printLog(WARN, v)
}

func (l *Logs) Error(v ...interface{}) {
	l.printLog(ERROR, v)
}

func (l *Logs) Fatal(v ...interface{}) {
	l.printLog(FATAL, v)
	os.Exit(1)
}

func (l *Logs) Panic(v ...interface{}) {
	l.printLog(PANIC, v)
	panic(v)
}

func (l *Logs) System(v ...interface{}) {
	lev := SYSTEM
	t := l.stampTime()
	msg := fmt.Sprintf("[%s] ", l.consoleColor(SYSTEM).Render(lev.shortString()))
	for i, val := range v {
		msg += fmt.Sprint(l.consoleColor(lev).Render(val))
		if i < len(v)-1 {
			msg += " "
		}
	}
	fmt.Printf("%s %s\n", t, msg)
}

func (l *Logs) Struct(data interface{}) {
	l.prettyStruct(data)
}

func (l *Logs) Line() {
	color.FgDarkGray.Println("-----------------------------------------------------")
}

// private methods

func (l *Logs) checkLevel(lev Level) bool {
	return Level(l.config.AppLogLevel) <= lev
}

func (l *Logs) stampTime() string {
	var t string
	f := "2006-01-02 15:04:05"
	now, err := carbon.NowInLocation(l.config.AppTimeZone)
	if err != nil || now == nil {
		t = time.Now().Format(f)
	} else {
		t = now.Format(f)
	}
	return fmt.Sprintf("%s", color.Normal.Render(t))
}

func (l *Logs) printLog(lev Level, v []interface{}) {
	if len(v) == 0 {
		return
	}
	if l.checkLevel(lev) {
		t := l.stampTime()
		msg := fmt.Sprintf("[%s] ", l.consoleColor(lev).Render(lev.shortString()))
		for i, val := range v {
			msg += fmt.Sprint(l.consoleColor(lev).Render(val))
			if i < len(v)-1 {
				msg += " "
			}
		}
		fmt.Printf("%s %s\n", t, msg)
	}
}

func (l *Logs) prettyStruct(v interface{}) {
	lev := PRINT
	if l.checkLevel(lev) {
		t := l.stampTime()
		msg := fmt.Sprintf("[%s] Type : %T", l.consoleColor(lev).Render(lev.shortString()), v)
		ms, _ := json.MarshalIndent(v, "", "   ")
		fmt.Printf("%s %s\n%s\n", t, msg, l.consoleColor(lev).Render(string(ms)))
	}
}

func (l *Logs) consoleColor(lev Level) color.Color {
	var clr color.Color
	if l.config.AppLogEnableColor {
		switch lev {
		case PRINT:
			clr = color.Normal
		case DEBUG:
			clr = color.FgLightGreen
		case INFO:
			clr = color.FgLightBlue
		case WARN:
			clr = color.FgLightYellow
		case ERROR:
			clr = color.FgRed
		case FATAL:
			clr = color.FgMagenta
		case SYSTEM:
			clr = color.FgLightCyan
		case PANIC:
			clr = color.FgMagenta
		default:
			clr = color.Normal
		}
	} else {
		clr = color.Normal
	}
	return clr
}
