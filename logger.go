package utils

import (
	gologging "github.com/op/go-logging"
	"os"
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
	// setup logging backend
	logFormat := gologging.MustStringFormatter(string(b.format))
	var logBackend = gologging.NewLogBackend(os.Stdout, "", 0)
	var logBackendFormatter = gologging.NewBackendFormatter(logBackend, logFormat)
	var logBackendLeveled = gologging.AddModuleLevel(logBackend)
	logBackendLeveled.SetLevel(gologging.ERROR, "")

	// Set the backends to be used.
	gologging.SetBackend(logBackendLeveled, logBackendFormatter)

	log := gologging.MustGetLogger("go-utils-logger")
	return &logger{
		log:    log,
		level:  b.level,
		format: b.format,
	}
}

/*
 * OOP Methods
 */

type logger struct {
	log    *gologging.Logger
	level  Level
	format Format
}

func (l *logger) Fatal(args ...interface{}) {
	l.log.Fatal(args...)
}

func (l *logger) Debug(args ...interface{}) {
	l.log.Debug(args...)
}
