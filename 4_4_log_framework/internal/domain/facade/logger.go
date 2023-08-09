package facade

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"log.framework/internal/domain"
	domainExport "log.framework/internal/domain/exporter"
	"log.framework/internal/domain/layout"
)

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

// ParseConfig parses the config file, but I use panic here for convenience.
// In generally, it's not a good practice to use panic here.
func ParseConfig(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(fmt.Errorf("failed to open config file: %w", err))
	}

	bytes, err := io.ReadAll(file)
	if err != nil {
		panic(fmt.Errorf("failed to read config file: %w", err))
	}

	var m map[string]interface{}
	err = json.Unmarshal(bytes, &m)
	if err != nil {
		panic(fmt.Errorf("failed to parse config file: %w", err))
	}

	var allLoggers []domain.Logger

	rootLogger, leftPart := parseRoot(bytes, m, err)
	parseChildren(leftPart, rootLogger, &allLoggers)

	domain.RegisterLogger(allLoggers...)
}

func parseRoot(bytes []byte, m map[string]interface{}, err error) (domain.Logger, map[string]interface{}) {
	bytes, _ = json.Marshal(m["loggers"])
	rootLogger := parseLogger(bytes)

	rootL := domain.NewLogger(
		domain.RootLoggerName,
		domain.WithLoggerLevel(getLevel(rootLogger.LevelThreshold)),
		domain.WithLayout(getLayout(rootLogger.Layout)),
		domain.WithExporter(getExporter(rootLogger.Exporter)),
	)

	var leftPart map[string]interface{}
	err = json.Unmarshal(bytes, &leftPart)
	if err != nil {
		panic(fmt.Errorf("failed to parse loggers key of remaining part: %w", err))
	}

	return rootL, leftPart
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
		err := json.Unmarshal(bytes, &newLeftPart)
		if err != nil {
			panic(fmt.Errorf("failed to parse loggers key of remaining child part: %w", err))
		}

		parseChildren(newLeftPart, newLogger, allLoggers)
	}
}

func parseLogger(b []byte) logger {
	var l logger
	err := json.Unmarshal(b, &l)
	if err != nil {
		panic(fmt.Errorf("failed to parse logger: %w", err))
	}
	return l
}

func isWhileListKey(key string) bool {
	return key == "levelThreshold" || key == "layout" || key == "exporter"
}

func getLevel(threshold string) domain.LoggerLevel {
	level, err := domain.ParseLoggerLevel(threshold)
	if err != nil {
		panic(fmt.Errorf("failed to parse logger level: %w", err))
	}
	return level
}

func getExporter(e exporter) domain.Exporter {
	var output domain.Exporter

	switch e.Type {
	case "console":
		output = domainExport.NewConsoleExporter()
	case "file":
		output = domainExport.NewFileExporter(e.FileName)
	case "composite":
		allExporters := make([]domain.Exporter, 0)
		for _, child := range e.Children {
			allExporters = append(allExporters, getExporter(child))
		}
		output = domainExport.NewCompositeExporter(allExporters...)
	case "":
		// not set, inherit from parent
	default:
		panic(fmt.Errorf("unsupported exporter type: %s", e.Type))
	}

	return output
}

func getLayout(l string) domain.Layout {
	if l == "standard" {
		return layout.NewStandardLayout()
	}
	return nil
}
