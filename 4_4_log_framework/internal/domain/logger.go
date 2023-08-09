package domain

import "fmt"

type Layout interface {
	Format(loggerName, content string, level LoggerLevel) string
}

type Exporter interface {
	Print(log string)
}

type Logger interface {
	Trace(log string)
	Info(log string)
	Debug(log string)
	Warn(log string)
	Error(log string)

	getExporter() Exporter
	getLoggerLevel() LoggerLevel
	getLayout() Layout
	getName() string
	getParentLogger() Logger
}

var RootLoggerName = "Root"

type logger struct {
	name         string
	level        LoggerLevel
	layout       Layout
	exporter     Exporter
	parentLogger Logger
}

var loggers = make(map[string]Logger)

func NewLogger(name string, options ...LoggerOption) Logger {
	l := &logger{
		name: name,
	}

	for _, option := range options {
		option(l)
	}

	if l.parentLogger != nil {
		if l.level == 0 {
			l.level = l.parentLogger.getLoggerLevel()
		}
		if l.layout == nil {
			l.layout = l.parentLogger.getLayout()
		}
		if l.exporter == nil {
			l.exporter = l.parentLogger.getExporter()
		}
	}

	return l
}

func RegisterLogger(logger ...Logger) {
	for _, l := range logger {
		if _, ok := loggers[l.getName()]; ok {
			panic(fmt.Sprintf("logger %s already registered", l.getName()))
		}

		if l.getName() != RootLoggerName && l.getParentLogger() == nil {
			panic(fmt.Sprintf("logger %s should have parent", l.getName()))

		}

		loggers[l.getName()] = l
	}
}

func GetLogger(name string) (Logger, error) {
	if l, ok := loggers[name]; !ok {
		return nil, fmt.Errorf("logger %s not found", name)
	} else {

		return l, nil
	}
}

func (l *logger) print(content string) {
	l.exporter.Print(l.layout.Format(l.name, content, l.level))
}

func (l *logger) Trace(log string) {
	if !l.isValidLevel(l.level, LoggerLevelTrace) {
		return
	}

	l.print(log)
}

func (l *logger) Info(log string) {
	if !l.isValidLevel(l.level, LoggerLevelInfo) {
		return
	}

	l.print(log)
}

func (l *logger) Debug(log string) {
	if !l.isValidLevel(l.level, LoggerLevelDebug) {
		return
	}

	l.print(log)
}

func (l *logger) Warn(log string) {
	if !l.isValidLevel(l.level, LoggerLevelWarn) {
		return
	}

	l.print(log)
}

func (l *logger) Error(log string) {
	if !l.isValidLevel(l.level, LoggerLevelError) {
		return
	}

	l.print(log)
}

func (l *logger) getExporter() Exporter {
	return l.exporter
}

func (l *logger) getLoggerLevel() LoggerLevel {
	return l.level
}

func (l *logger) getLayout() Layout {
	return l.layout
}

func (l *logger) getName() string {
	return l.name
}

func (l *logger) isValidLevel(current, target LoggerLevel) bool {
	return current >= target
}

func (l *logger) getParentLogger() Logger {
	return l.parentLogger
}
