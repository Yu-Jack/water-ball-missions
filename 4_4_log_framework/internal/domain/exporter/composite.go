package exporter

import (
	"log.framework/internal/domain"
)

type compositeExporter struct {
	exporters []domain.Exporter
}

func NewCompositeExporter(exporters ...domain.Exporter) domain.Exporter {
	return &compositeExporter{exporters: exporters}
}

func (c *compositeExporter) Print(log string) {
	for _, exporter := range c.exporters {
		exporter.Print(log)
	}
}
