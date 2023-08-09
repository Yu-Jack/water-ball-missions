package domain

type LoggerOption func(l *logger)

func WithLoggerLevel(level LoggerLevel) LoggerOption {
	return func(l *logger) {
		l.level = level
	}
}

func WithLayout(layout Layout) LoggerOption {
	return func(l *logger) {
		l.layout = layout
	}
}

func WithExporter(exporter Exporter) LoggerOption {
	return func(l *logger) {
		l.exporter = exporter
	}
}

func WithParentLogger(parentLogger Logger) LoggerOption {
	return func(l *logger) {
		l.parentLogger = parentLogger
	}
}
