package exporter

import (
	"os"

	"log.framework/internal/domain"
)

type fileExporter struct {
	file *os.File
}

func NewFileExporter(name string) domain.Exporter {
	dir, _ := os.Getwd()
	dir = dir + "/logs/"

	file, err := os.OpenFile(dir+name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	return &fileExporter{file: file}
}

func (c fileExporter) Print(log string) {
	_, _ = c.file.WriteString(log + "\n")
}
