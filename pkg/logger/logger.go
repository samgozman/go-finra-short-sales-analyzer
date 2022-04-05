package logger

import (
	"fmt"
	"time"
)

// Log info message
func Info(serviceName, message string) {
	printMsg(serviceName, message, "INFO")
}

// Log error message
func Error(serviceName, message string) {
	printMsg(serviceName, message, "ERROR")
}

func printMsg(serviceName, message, level string) {
	dt := time.Now()
	fmt.Printf("[%s] %s - %s: %s\n", serviceName, level, dt.Format(time.UnixDate), message)
}
