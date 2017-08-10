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
 * Format
 */
type Format string

//var defaultFormat Format = `%{color}%{time:15:04:05.000} %{id:03x} %{module:15s} %{shortfunc:20s} ▶ %{level:.4s} %{color:reset} %{message}`
var defaultFormat Format = `%{color}%{time:15:04:05.000} %{id:03x} ▶ %{level:.4s} %{color:reset} %{message}`

/*
 * Interface Definition
 */
type Logger interface {
	/*
		Critical(args ...interface{})
		Error(args ...interface{})
		Warning(args ...interface{})
		Notice(args ...interface{})
		Info(args ...interface{})
	*/
	Log(level Level, args ...interface{})
	Fatal(args ...interface{})
	Debug(args ...interface{})
}

type LoggerBuilder interface {
	SetLevel(Level) LoggerBuilder
	SetFormat(Format) LoggerBuilder
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

func (b *loggerBuilder) SetFormat(format Format) LoggerBuilder {
	b.format = format
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

func (l *logger) Log(level Level, args ...interface{}) {
	currentStr := time.Now().Format(time.RFC3339)
	callerStr := stack.GetCallingMethodName()
	levelStr := levelNames[l.level]
	msg := fmt.Sprintln(args...)
	logLine := fmt.Sprintf("%s %s > %-4s %s", currentStr, callerStr, levelStr, msg)
	fmt.Print(logLine)
}

func (l *logger) Fatal(args ...interface{}) {
	l.Log(CRITICAL, args)
}

func (l *logger) Debug(args ...interface{}) {
	l.Log(DEBUG, args)
	//newArgs := fmt.Sprintf("%s%s%s", args...)
	//l.log.Debug(newArgs)
}
