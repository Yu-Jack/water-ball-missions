//go:generate go-enum -f=$GOFILE --nocase

package domain

// LoggerLevel
/*
ENUM(
TRACE = 1
INFO
DEBUG
WARN
ERROR
)
*/
type LoggerLevel int
