package layout

import (
	"fmt"
	"time"

	"log.framework/internal/domain"
)

type standardLayout struct{}

func NewStandardLayout() domain.Layout {
	return &standardLayout{}
}

func (l *standardLayout) Format(loggerName, content string, level domain.LoggerLevel) string {
	return fmt.Sprintf("%s |-%s %s - %s", time.Now(), level.String(), loggerName, content)
}
