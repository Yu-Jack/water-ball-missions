package exporter

import (
	"fmt"

	"log.framework/internal/domain"
)

type consoleExporter struct {
}

func NewConsoleExporter() domain.Exporter {
	return &consoleExporter{}
}

func (c *consoleExporter) Print(log string) {
	fmt.Println(log)
}
