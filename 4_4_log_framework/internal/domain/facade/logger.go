package facade

import (
	"encoding/json"
	"io"
	"os"

	"log.framework/internal/domain"
	domainExport "log.framework/internal/domain/exporter"
	"log.framework/internal/domain/layout"
)

type root struct {
	Loggers logger `json:"loggers"`
}

type logger struct {
	LevelThreshold string   `json:"levelThreshold"`
	Exporter       exporter `json:"exporter"`
	Layout         string   `json:"layout"`
}

type exporter struct {
	Type     string     `json:"type"`
	FileName string     `json:"fileName"`
	Children []exporter `json:"children"`
}

func ParseConfig(filename string) {
	file, _ := os.Open("config.json")
	bytes, _ := io.ReadAll(file)
	var m map[string]interface{}
	_ = json.Unmarshal(bytes, &m)

	bytes, _ = json.Marshal(m["loggers"])
	rootLogger := parseLogger(bytes)

	rootL := domain.NewLogger(
		domain.RootLoggerName,
		domain.WithLoggerLevel(getLevel(rootLogger.LevelThreshold)),
		domain.WithLayout(getLayout(rootLogger.Layout)),
		domain.WithExporter(getExporter(rootLogger.Exporter)),
	)

	var leftParents map[string]interface{}
	var allLoggers []domain.Logger
	_ = json.Unmarshal(bytes, &leftParents)

	parseChildren(leftParents, rootL, &allLoggers)

	domain.RegisterLogger(allLoggers...)
}

func parseChildren(leftPart map[string]interface{}, parentLogger domain.Logger, allLoggers *[]domain.Logger) {
	*allLoggers = append(*allLoggers, parentLogger)

	for k, v := range leftPart {
		if isWhileListKey(k) {
			continue
		}

		bytes, _ := json.Marshal(v)
		l := parseLogger(bytes)

		newLogger := domain.NewLogger(
			k,
			domain.WithParentLogger(parentLogger),
			domain.WithLoggerLevel(getLevel(l.LevelThreshold)),
			domain.WithLayout(getLayout(l.Layout)),
			domain.WithExporter(getExporter(l.Exporter)),
		)

		var newLeftPart map[string]interface{}
		_ = json.Unmarshal(bytes, &newLeftPart)

		parseChildren(newLeftPart, newLogger, allLoggers)
	}
}

func parseLogger(b []byte) logger {
	var l logger
	_ = json.Unmarshal(b, &l)
	return l
}

func isWhileListKey(key string) bool {
	return key == "levelThreshold" || key == "layout" || key == "exporter"
}

func getLevel(threshold string) domain.LoggerLevel {
	level, _ := domain.ParseLoggerLevel(threshold)
	return level
}

func getExporter(e exporter) domain.Exporter {
	var output domain.Exporter

	if e.Type == "console" {
		output = domainExport.NewConsoleExporter()
	}

	if e.Type == "file" {
		output = domainExport.NewFileExporter(e.FileName)
	}

	if e.Type == "composite" {
		allExporters := make([]domain.Exporter, 0)
		for _, child := range e.Children {
			allExporters = append(allExporters, getExporter(child))
		}
		output = domainExport.NewCompositeExporter(allExporters...)
	}

	return output
}

func getLayout(l string) domain.Layout {
	if l == "standard" {
		return layout.NewStandardLayout()
	}
	return nil
}
