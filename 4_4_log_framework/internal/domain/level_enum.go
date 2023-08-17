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
	// LoggerLevelTRACE is a LoggerLevel of type TRACE.
	LoggerLevelTRACE LoggerLevel = iota + 1
	// LoggerLevelINFO is a LoggerLevel of type INFO.
	LoggerLevelINFO
	// LoggerLevelDEBUG is a LoggerLevel of type DEBUG.
	LoggerLevelDEBUG
	// LoggerLevelWARN is a LoggerLevel of type WARN.
	LoggerLevelWARN
	// LoggerLevelERROR is a LoggerLevel of type ERROR.
	LoggerLevelERROR
)

var ErrInvalidLoggerLevel = errors.New("not a valid LoggerLevel")

const _LoggerLevelName = "TRACEINFODEBUGWARNERROR"

var _LoggerLevelMap = map[LoggerLevel]string{
	LoggerLevelTRACE: _LoggerLevelName[0:5],
	LoggerLevelINFO:  _LoggerLevelName[5:9],
	LoggerLevelDEBUG: _LoggerLevelName[9:14],
	LoggerLevelWARN:  _LoggerLevelName[14:18],
	LoggerLevelERROR: _LoggerLevelName[18:23],
}

// String implements the Stringer interface.
func (x LoggerLevel) String() string {
	if str, ok := _LoggerLevelMap[x]; ok {
		return str
	}
	return fmt.Sprintf("LoggerLevel(%d)", x)
}

var _LoggerLevelValue = map[string]LoggerLevel{
	_LoggerLevelName[0:5]:                    LoggerLevelTRACE,
	strings.ToLower(_LoggerLevelName[0:5]):   LoggerLevelTRACE,
	_LoggerLevelName[5:9]:                    LoggerLevelINFO,
	strings.ToLower(_LoggerLevelName[5:9]):   LoggerLevelINFO,
	_LoggerLevelName[9:14]:                   LoggerLevelDEBUG,
	strings.ToLower(_LoggerLevelName[9:14]):  LoggerLevelDEBUG,
	_LoggerLevelName[14:18]:                  LoggerLevelWARN,
	strings.ToLower(_LoggerLevelName[14:18]): LoggerLevelWARN,
	_LoggerLevelName[18:23]:                  LoggerLevelERROR,
	strings.ToLower(_LoggerLevelName[18:23]): LoggerLevelERROR,
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