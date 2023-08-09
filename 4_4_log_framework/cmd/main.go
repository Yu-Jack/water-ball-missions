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
		domain.WithExporter(exporter.NewConsoleExporter()),
	)

	appNameLogger := domain.NewLogger(
		"app.game",
		domain.WithParentLogger(rootLogger),
		domain.WithLoggerLevel(domain.LoggerLevelInfo),
		domain.WithExporter(
			exporter.NewCompositeExporter(
				exporter.NewConsoleExporter(),
				exporter.NewCompositeExporter(
					exporter.NewFileExporter("game.log"),
					exporter.NewFileExporter("game.backup.log"),
				),
			),
		),
	)

	aiNameLogger := domain.NewLogger(
		"app.game.ai",
		domain.WithParentLogger(appNameLogger),
		domain.WithLoggerLevel(domain.LoggerLevelTrace),
		domain.WithLayout(layout.NewStandardLayout()),
	)

	domain.RegisterLogger(rootLogger, appNameLogger, aiNameLogger)

	r, _ := domain.GetLogger(domain.RootLoggerName)
	appL, _ := domain.GetLogger("app.game")
	appAi, _ := domain.GetLogger("app.game.ai")

	r.Debug("This is root")
	appL.Trace("Hi, app.game is running")
	appAi.Trace("Hi, app.game.ai is running, but traced.")
}
