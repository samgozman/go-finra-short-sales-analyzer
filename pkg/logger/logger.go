package logger

import (
	"fmt"
	"time"
)

func Info(serviceName, message string) {
	printMsg(serviceName, message, "INFO")
}

func Error(serviceName, message string) {
	printMsg(serviceName, message, "ERROR")
}

func printMsg(serviceName, message, level string) {
	dt := time.Now()
	fmt.Printf("[%s] %s - %s: %s\n", serviceName, level, dt.Format(time.UnixDate), message)
}
