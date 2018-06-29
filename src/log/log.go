package log

import (
	"fmt"
	"path/filepath"
	"runtime"
	"time"
)

type Level int

const (
	LOG_DEBUG Level = iota
	LOG_INFO
	LOG_ERROR
)

var (
	LogLevel   Level  = LOG_INFO
	DateFormat string = "2006-01-02 15:04:05"

	FormatFunctions map[Level]func(string, string, int, string, string) = map[Level]func(string, string, int, string, string){
		LOG_DEBUG: logDefault,
		LOG_INFO:  logDefault,
		LOG_ERROR: logDefault,
	}

	// The current maximum length printed for caller information. This is updated each time something gets printed
	CallerColumnWidth = 0

	levelStrings map[Level]string = map[Level]string{
		LOG_DEBUG: "[DEBUG]",
		LOG_INFO:  "[INFO] ",
		LOG_ERROR: "[ERROR]",
	}
)

func Info(message string) {
	log(LOG_INFO, message)
}

func Debug(message string) {
	log(LOG_DEBUG, message)
}

func Error(message string) {
	log(LOG_ERROR, message)
}

func log(level Level, message string) {
	caller := getCallerDetails()

	updateCallerColumnWidth(caller)

	if LogLevel <= level {
		FormatFunctions[level](time.Now().Format(DateFormat), levelStrings[level], CallerColumnWidth, caller, message)
	}
}

func updateCallerColumnWidth(caller string) {
	if len(caller) > CallerColumnWidth {
		CallerColumnWidth = len(caller)
	}
}

func getCallerDetails() string {
	name := "???"
	line := -1

	// A bit hacky: We know here that the stack contains two calls from inside
	// this file. The third frame comes from the file that initially called a
	// function in this file (e.g. Info())
	pc, _, _, ok := runtime.Caller(3)

	if ok {
		details := runtime.FuncForPC(pc)

		name, line = details.FileLine(pc)
		name = filepath.Base(name)
	}

	caller := fmt.Sprintf("%s:%d", name, line)

	return caller
}

func logDefault(time, level string, maxLength int, caller, message string) {
	fmt.Printf("%s %s %-*s | %s\n", time, level, maxLength, caller, message)
}
