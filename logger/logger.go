package logger

import (
	"fmt"
	"log"
	"time"
)

// LogLevel defines the loglevels
type LogLevel string

// const for different log levels
const (
	INFO  LogLevel = "[INFO] "
	DEBUG LogLevel = "[DEBUG] "
	ERROR LogLevel = "[ERROR] "
	FATAL LogLevel = "[FATAL] "
)

func init() {
	log.SetPrefix("# ")
}

// TimeLogger will log the time taken by the request
func TimeLogger(operation string, start time.Time) {
	Log(INFO, operation, time.Since(start), "\n")
}

// Log will do the logging
func Log(logLevel LogLevel, msg ...interface{}) {
	switch logLevel {
	case INFO:
		log.Print(INFO, fmt.Sprintln(msg...))
	case DEBUG:
		log.Print(DEBUG, fmt.Sprintln(msg...))
	case ERROR:
		log.Print(ERROR, fmt.Sprintln(msg...))
	case FATAL:
		log.Fatal(FATAL, fmt.Sprintln(msg...))
	default:
		log.Print(fmt.Sprintln(msg...))
	}
}
