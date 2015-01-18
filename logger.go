package multilog

type Logger interface {
	Fatal(v ...interface{})
	Fatalf(format string, v ...interface{})
	Fatalln(v ...interface{})
	Panic(v ...interface{})
	Panicf(format string, v ...interface{})
	Panicln(v ...interface{})
	Print(v ...interface{})
	Printf(format string, v ...interface{})
	Println(v ...interface{})

	Debugf(format string, v ...interface{})
	Infof(format string, v ...interface{})
	Warnf(format string, v ...interface{})
	Warningf(format string, v ...interface{})
	Errorf(format string, v ...interface{})
	Debug(v ...interface{})
	Info(v ...interface{})
	Warn(v ...interface{})
	Warning(v ...interface{})
	Error(v ...interface{})
	Debugln(v ...interface{})
	Infoln(v ...interface{})
	Warnln(v ...interface{})
	Warningln(v ...interface{})
	Errorln(v ...interface{})
}

type MultiLogger struct {
	loggers []Logger
}

func New(loggers ...Logger) *MultiLogger {
	return &MultiLogger{loggers: loggers}
}

func (l *MultiLogger) Fatal(v ...interface{}) {
	for _, logger := range l.loggers {
		logger.Fatal(v)
	}
}

func (l *MultiLogger) Fatalf(format string, v ...interface{}) {
	for _, logger := range l.loggers {
		logger.Fatalf(format, v)
	}
}

func (l *MultiLogger) Fatalln(v ...interface{}) {
	for _, logger := range l.loggers {
		logger.Fatalln(v)
	}
}

func (l *MultiLogger) Panic(v ...interface{}) {
	for _, logger := range l.loggers {
		logger.Panic(v)
	}
}

func (l *MultiLogger) Panicf(format string, v ...interface{}) {
	for _, logger := range l.loggers {
		logger.Panicf(format, v)
	}
}

func (l *MultiLogger) Panicln(v ...interface{}) {
	for _, logger := range l.loggers {
		logger.Panicln(v)
	}
}

func (l *MultiLogger) Print(v ...interface{}) {
	for _, logger := range l.loggers {
		logger.Print(v)
	}
}

func (l *MultiLogger) Printf(format string, v ...interface{}) {
	for _, logger := range l.loggers {
		logger.Printf(format, v)
	}
}

func (l *MultiLogger) Println(v ...interface{}) {
	for _, logger := range l.loggers {
		logger.Println(v)
	}
}

func (l *MultiLogger) Debugf(format string, v ...interface{}) {
	for _, logger := range l.loggers {
		logger.Debugf(format, v)
	}
}

func (l *MultiLogger) Infof(format string, v ...interface{}) {
	for _, logger := range l.loggers {
		logger.Infof(format, v)
	}
}

func (l *MultiLogger) Warnf(format string, v ...interface{}) {
	for _, logger := range l.loggers {
		logger.Warnf(format, v)
	}
}

func (l *MultiLogger) Warningf(format string, v ...interface{}) {
	for _, logger := range l.loggers {
		logger.Warningf(format, v)
	}
}

func (l *MultiLogger) Errorf(format string, v ...interface{}) {
	for _, logger := range l.loggers {
		logger.Errorf(format, v)
	}
}

func (l *MultiLogger) Debug(v ...interface{}) {
	for _, logger := range l.loggers {
		logger.Debug(v)
	}
}

func (l *MultiLogger) Info(v ...interface{}) {
	for _, logger := range l.loggers {
		logger.Info(v)
	}
}

func (l *MultiLogger) Warn(v ...interface{}) {
	for _, logger := range l.loggers {
		logger.Warn(v)
	}
}

func (l *MultiLogger) Warning(v ...interface{}) {
	for _, logger := range l.loggers {
		logger.Warning(v)
	}
}

func (l *MultiLogger) Error(v ...interface{}) {
	for _, logger := range l.loggers {
		logger.Error(v)
	}
}

func (l *MultiLogger) Debugln(v ...interface{}) {
	for _, logger := range l.loggers {
		logger.Debugln(v)
	}
}

func (l *MultiLogger) Infoln(v ...interface{}) {
	for _, logger := range l.loggers {
		logger.Infoln(v)
	}
}

func (l *MultiLogger) Warnln(v ...interface{}) {
	for _, logger := range l.loggers {
		logger.Warnln(v)
	}
}

func (l *MultiLogger) Warningln(v ...interface{}) {
	for _, logger := range l.loggers {
		logger.Warningln(v)
	}
}

func (l *MultiLogger) Errorln(v ...interface{}) {
	for _, logger := range l.loggers {
		logger.Errorln(v)
	}
}
