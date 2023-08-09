package main

import (
	"log.framework/internal/domain"
	"log.framework/internal/domain/exporter"
	"log.framework/internal/domain/layout"
)

func main() {
	rootLogger := domain.NewLogger(
		domain.RootLoggerName,
		domain.WithLoggerLevel(domain.LoggerLevelDebug),
		domain.WithLayout(layout.NewStandardLayout()),
		domain.WithExporter(
			exporter.NewCompositeExporter(
				exporter.NewFileExporter("root.log"),
				exporter.NewFileExporter("root.log.backup"),
				exporter.NewConsoleExporter(),
			),
		),
	)

	appNameLogger := domain.NewLogger(
		"app.game",
		domain.WithParentLogger(rootLogger),
		domain.WithLoggerLevel(domain.LoggerLevelInfo),
		domain.WithExporter(
			exporter.NewCompositeExporter(
				exporter.NewConsoleExporter(),
				exporter.NewCompositeExporter(
					exporter.NewFileExporter("app.game.log"),
					exporter.NewFileExporter("app.game.log.backup"),
				),
			),
		),
	)

	domain.RegisterLogger(rootLogger, appNameLogger)

	appL, _ := domain.GetLogger("app.game")
	r, _ := domain.GetLogger(domain.RootLoggerName)

	appL.Info("Hi, I'm Jack")
	r.Debug("Hi, I'm Jack")
}
