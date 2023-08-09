// Code generated by go-enum DO NOT EDIT.
// Version:
// Revision:
// Build Date:
// Built By:

package domain

import (
	"errors"
	"fmt"
	"strings"
)

const (
	// LoggerLevelTrace is a LoggerLevel of type Trace.
	LoggerLevelTrace LoggerLevel = iota + 1
	// LoggerLevelInfo is a LoggerLevel of type Info.
	LoggerLevelInfo
	// LoggerLevelDebug is a LoggerLevel of type Debug.
	LoggerLevelDebug
	// LoggerLevelWarn is a LoggerLevel of type Warn.
	LoggerLevelWarn
	// LoggerLevelError is a LoggerLevel of type Error.
	LoggerLevelError
)

var ErrInvalidLoggerLevel = errors.New("not a valid LoggerLevel")

const _LoggerLevelName = "TraceInfoDebugWarnError"

var _LoggerLevelMap = map[LoggerLevel]string{
	LoggerLevelTrace: _LoggerLevelName[0:5],
	LoggerLevelInfo:  _LoggerLevelName[5:9],
	LoggerLevelDebug: _LoggerLevelName[9:14],
	LoggerLevelWarn:  _LoggerLevelName[14:18],
	LoggerLevelError: _LoggerLevelName[18:23],
}

// String implements the Stringer interface.
func (x LoggerLevel) String() string {
	if str, ok := _LoggerLevelMap[x]; ok {
		return str
	}
	return fmt.Sprintf("LoggerLevel(%d)", x)
}

var _LoggerLevelValue = map[string]LoggerLevel{
	_LoggerLevelName[0:5]:                    LoggerLevelTrace,
	strings.ToLower(_LoggerLevelName[0:5]):   LoggerLevelTrace,
	_LoggerLevelName[5:9]:                    LoggerLevelInfo,
	strings.ToLower(_LoggerLevelName[5:9]):   LoggerLevelInfo,
	_LoggerLevelName[9:14]:                   LoggerLevelDebug,
	strings.ToLower(_LoggerLevelName[9:14]):  LoggerLevelDebug,
	_LoggerLevelName[14:18]:                  LoggerLevelWarn,
	strings.ToLower(_LoggerLevelName[14:18]): LoggerLevelWarn,
	_LoggerLevelName[18:23]:                  LoggerLevelError,
	strings.ToLower(_LoggerLevelName[18:23]): LoggerLevelError,
}

// ParseLoggerLevel attempts to convert a string to a LoggerLevel.
func ParseLoggerLevel(name string) (LoggerLevel, error) {
	if x, ok := _LoggerLevelValue[name]; ok {
		return x, nil
	}
	// Case insensitive parse, do a separate lookup to prevent unnecessary cost of lowercasing a string if we don't need to.
	if x, ok := _LoggerLevelValue[strings.ToLower(name)]; ok {
		return x, nil
	}
	return LoggerLevel(0), fmt.Errorf("%s is %w", name, ErrInvalidLoggerLevel)
}
