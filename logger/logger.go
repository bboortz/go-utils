package logger

import (
	"fmt"
	"github.com/bboortz/go-utils/stack"
	"log"
	"os"
	"time"
)

/*
 * Log levels
 */
type Level int

const (
	CRITICAL Level = iota
	ERROR
	WARNING
	NOTICE
	INFO
	DEBUG
)

var levelNames = []string{
	"CRITICAL",
	"ERROR",
	"WARNING",
	"NOTICE",
	"INFO",
	"DEBUG",
}

/*
 * Interface Definition
 */
type Logger interface {
	Log(level Level, caller string, args ...interface{})
	Fatal(args ...interface{})
	Critical(args ...interface{})
	Error(args ...interface{})
	Warning(args ...interface{})
	Notice(args ...interface{})
	Info(args ...interface{})
	Debug(args ...interface{})
}

type LoggerBuilder interface {
	SetLevel(Level) LoggerBuilder
	Build() Logger
}

type loggerBuilder struct {
	level  Level
	format Format
}

/*
 * Builder Methods
 */
func NewLogger() LoggerBuilder {
	return &loggerBuilder{format: defaultFormat}
}

func (b *loggerBuilder) SetLevel(level Level) LoggerBuilder {
	b.level = level
	return b
}

func (b *loggerBuilder) Build() Logger {
	return &logger{
		logStream: log.New(os.Stdout, "", 0),
		level:     b.level,
		format:    b.format,
	}
}

/*
 * OOP Methods
 */

type logger struct {
	logStream *log.Logger
	level     Level
	format    Format
}

func (l *logger) Log(level Level, caller string, args ...interface{}) {
	currentStr := time.Now().Format(time.RFC3339)
	levelStr := levelNames[l.level]
	msg := fmt.Sprintln(args...)
	logLine := fmt.Sprintf("%s %s > %-4s %s", currentStr, caller, levelStr, msg)
	l.logStream.Printf(logLine)
}

func (l *logger) Fatal(args ...interface{}) {
	callerStr := stack.GetCallingMethodName()
	l.Log(CRITICAL, callerStr, args)
	os.Exit(1)
}

func (l *logger) Critical(args ...interface{}) {
	callerStr := stack.GetCallingMethodName()
	l.Log(CRITICAL, callerStr, args)
}

func (l *logger) Error(args ...interface{}) {
	callerStr := stack.GetCallingMethodName()
	l.Log(ERROR, callerStr, args)
}

func (l *logger) Warning(args ...interface{}) {
	callerStr := stack.GetCallingMethodName()
	l.Log(WARNING, callerStr, args)
}

func (l *logger) Notice(args ...interface{}) {
	callerStr := stack.GetCallingMethodName()
	l.Log(NOTICE, callerStr, args)
}

func (l *logger) Info(args ...interface{}) {
	callerStr := stack.GetCallingMethodName()
	l.Log(INFO, callerStr, args)
}

func (l *logger) Debug(args ...interface{}) {
	callerStr := stack.GetCallingMethodName()
	l.Log(DEBUG, callerStr, args)
}
