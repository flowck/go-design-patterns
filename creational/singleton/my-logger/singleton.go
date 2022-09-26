package my_logger

import "fmt"

type MyLogger struct {
	level int
}

func (l *MyLogger) SetLevel(value int) {
	l.level = value
}

func (l *MyLogger) Log(value any) {
	fmt.Println(value)
}

var logger *MyLogger

func GetLoggerInstance() *MyLogger {
	if logger == nil {
		fmt.Println("Creating a logger instance")
		logger = &MyLogger{}
	}

	return logger
}
