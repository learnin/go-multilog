# go-multilog

A simple wrapper for multiple logger instances.

## Example

example using https://github.com/sirupsen/logrus
```go
stdOutLogger := logrus.New()
stdOutLogger.Formatter = &logrus.TextFormatter{DisableColors: false}
stdOutLogger.Level = logrus.WarnLevel

logf, _ := os.OpenFile("example.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
fileLogger := logrus.New()
fileLogger.Out = logf
fileLogger.Formatter = &logrus.TextFormatter{DisableColors: true}
fileLogger.Level = logrus.DebugLevel

log = multilog.New(stdOutLogger, fileLogger)

log.Debug("debug message")
log.Warn("warn message")
```

## License

MIT
