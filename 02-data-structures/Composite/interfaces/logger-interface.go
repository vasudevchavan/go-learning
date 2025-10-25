package main

import (
	"fmt"
)

// Logger is the interface that defines the logging behavior.
type Logger interface {
	Log(message string)
}

// ConsoleLogger writes logs to the console.
type ConsoleLogger struct{}

func (c ConsoleLogger) Log(message string) {
	fmt.Printf("[ConsoleLogger] %s\n", message)
}

// FileLogger simulates writing logs to a file.
type FileLogger struct {
	FileName string
}

func (f FileLogger) Log(message string) {
	// In a real application, you'd open the file and write to it.
	// Here we simulate it with console output.
	fmt.Printf("[FileLogger] Writing '%s' to file %s\n", message, f.FileName)
}

// ProcessTask simulates an application process that logs events.
func ProcessTask(task string, logger Logger) {
	logger.Log("Starting task: " + task)

	// Simulate doing some work
	fmt.Println("...doing work...")

	logger.Log("Completed task: " + task)
}

func main() {
	// Use ConsoleLogger
	consoleLogger := ConsoleLogger{}
	ProcessTask("Send Email", consoleLogger)

	fmt.Println()

	// Use FileLogger
	fileLogger := FileLogger{FileName: "app.log"}
	ProcessTask("Generate Report", fileLogger)
}
