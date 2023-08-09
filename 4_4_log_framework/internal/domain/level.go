//go:generate go-enum -f=$GOFILE --nocase

package domain

// LoggerLevel
/*
ENUM(
Trace = 1
Info
Debug
Warn
Error
)
*/
type LoggerLevel int
