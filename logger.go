package golib

import (
	"fmt"
	"runtime"
	"time"
)

const (
	formatTime = "2006-01-02 15:04:05"
	ERROR      = "ERROR"
	WARN       = "WARN"
	INFO       = "INFO"
	TRACE      = "TRACE"
	DEBUG      = "DEBUG"
)

func formatLog(message string, level string, err error, data ...interface{}) {
	pc, fn, line, _ := runtime.Caller(1)
	messagePrint := message
	for _, val := range data {
		messagePrint = fmt.Sprintf(messagePrint+" %v", val)
	}
	if err != nil {
		messagePrint = fmt.Sprintf("%s %v", messagePrint, err.Error())
	}
	timeNow := time.Now()
	fmt.Println(fmt.Sprintf("%s %s %s[%s:%d] - %s", timeNow.Format(formatTime), level, runtime.FuncForPC(pc).Name(), fn, line, messagePrint))
}

func Debug(message string, err error, data ...interface{}) {
	pc, fn, line, _ := runtime.Caller(1)
	messagePrint := message
	for _, val := range data {
		messagePrint = fmt.Sprintf(messagePrint+" %v", val)
	}
	if err != nil {
		messagePrint = fmt.Sprintf("%s %v", messagePrint, err.Error())
	}
	timeNow := time.Now()
	fmt.Println(fmt.Sprintf("%s %s %s[%s:%d] - %s", timeNow.Format(formatTime), DEBUG, runtime.FuncForPC(pc).Name(), fn, line, messagePrint))
}

func Fatal(message string, err error, data ...interface{}) {
	pc, fn, line, _ := runtime.Caller(1)
	messagePrint := message
	for _, val := range data {
		messagePrint = fmt.Sprintf(messagePrint+" %v", val)
	}
	if err != nil {
		messagePrint = fmt.Sprintf("%s %v", messagePrint, err.Error())
	}
	timeNow := time.Now()
	fmt.Println(fmt.Sprintf("%s %s %s[%s:%d] - %s", timeNow.Format(formatTime), ERROR, runtime.FuncForPC(pc).Name(), fn, line, messagePrint))
	panic(err)
}

func Error(message string, err error, data ...interface{}) {
	pc, fn, line, _ := runtime.Caller(1)
	messagePrint := message
	for _, val := range data {
		messagePrint = fmt.Sprintf(messagePrint+" %v", val)
	}
	if err != nil {
		messagePrint = fmt.Sprintf("%s %v", messagePrint, err.Error())
	}
	timeNow := time.Now()
	fmt.Println(fmt.Sprintf("%s %s %s[%s:%d] - %s", timeNow.Format(formatTime), ERROR, runtime.FuncForPC(pc).Name(), fn, line, messagePrint))
}

func Info(message string, data ...interface{}) {
	pc, fn, line, _ := runtime.Caller(1)
	messagePrint := message
	for _, val := range data {
		messagePrint = fmt.Sprintf(messagePrint+" %v", val)
	}
	timeNow := time.Now()
	fmt.Println(fmt.Sprintf("%s %s %s[%s:%d] - %s", timeNow.Format(formatTime), INFO, runtime.FuncForPC(pc).Name(), fn, line, messagePrint))
}

func Warn(message string, err error, data ...interface{}) {
	pc, fn, line, _ := runtime.Caller(1)
	messagePrint := message
	for _, val := range data {
		messagePrint = fmt.Sprintf(messagePrint+" %v", val)
	}
	timeNow := time.Now()
	fmt.Println(fmt.Sprintf("%s %s %s[%s:%d] - %s", timeNow.Format(formatTime), WARN, runtime.FuncForPC(pc).Name(), fn, line, messagePrint))
}

func Trace(message string, err error, data ...interface{}) {
	pc, fn, line, _ := runtime.Caller(1)
	messagePrint := message
	for _, val := range data {
		messagePrint = fmt.Sprintf(messagePrint+" %v", val)
	}
	timeNow := time.Now()
	fmt.Println(fmt.Sprintf("%s %s %s[%s:%d] - %s", timeNow.Format(formatTime), TRACE, runtime.FuncForPC(pc).Name(), fn, line, messagePrint))
}
