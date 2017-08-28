package client

import "fmt"

type Logger interface {
	Debugf(format string, v ...interface{})
	Tracef(format string, v ...interface{})
	Infof(format string, v ...interface{})
	Warningf(format string, v ...interface{})
	Errorf(format string, v ...interface{})
	Fatalf(format string, v ...interface{})
}

type logger struct{}

var _ Logger = &logger{}

func (l logger) printf(level string, format string, v ...interface{}) {
	log := fmt.Sprintf(format, v...)
	fmt.Printf("[%s] %s\n", level, log)
}

func (l logger) Debugf(format string, v ...interface{}) {
	l.printf("DEBUG", format, v...)
}

func (l logger) Tracef(format string, v ...interface{}) {
	l.printf("TRACE", format, v...)
}

func (l logger) Infof(format string, v ...interface{}) {
	l.printf("INFO", format, v...)
}

func (l logger) Warningf(format string, v ...interface{}) {
	l.printf("WARN", format, v...)
}

func (l logger) Errorf(format string, v ...interface{}) {
	l.printf("ERR", format, v...)
}

func (l logger) Fatalf(format string, v ...interface{}) {
	panic(fmt.Sprintf(format, v...))
}
