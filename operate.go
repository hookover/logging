package logging

import (
	"github.com/rs/zerolog"
	"io"
)

func Chan(channel string) *zerolog.Logger {
	if logger, ok := logger.loggers[channel]; ok {
		return logger
	}
	return logger.loggers[logger.defaultChannel]
}

func Log() *zerolog.Event {
	return logger.loggers[logger.defaultChannel].Log()
}

func Level(lvl zerolog.Level) zerolog.Logger {
	return logger.loggers[logger.defaultChannel].Level(lvl)
}

func With() zerolog.Context {
	return logger.loggers[logger.defaultChannel].With()
}

func WithLevel(lvl zerolog.Level) *zerolog.Event {
	return logger.loggers[logger.defaultChannel].WithLevel(lvl)
}

func Debug() *zerolog.Event {
	return logger.loggers[logger.defaultChannel].Debug()
}

func Info() *zerolog.Event {
	return logger.loggers[logger.defaultChannel].Info()
}

func Warn() *zerolog.Event {
	return logger.loggers[logger.defaultChannel].Warn()
}

func Error() *zerolog.Event {
	return logger.loggers[logger.defaultChannel].Error()
}

func Panic() *zerolog.Event {
	return logger.loggers[logger.defaultChannel].Panic()
}

func Fatal() *zerolog.Event {
	return logger.loggers[logger.defaultChannel].Fatal()
}

func Output(w io.Writer) zerolog.Logger {
	return logger.loggers[logger.defaultChannel].Output(w)
}

func Write(p []byte) (n int, err error) {
	return logger.loggers[logger.defaultChannel].Write(p)
}

func Printf(format string, v ...interface{}) {
	logger.loggers[logger.defaultChannel].Printf(format, v...)
}

func Hook(h zerolog.Hook) zerolog.Logger {
	return logger.loggers[logger.defaultChannel].Hook(h)
}

func Sample(s zerolog.Sampler) zerolog.Logger {
	return logger.loggers[logger.defaultChannel].Sample(s)
}

func UpdateContext(update func(c zerolog.Context) zerolog.Context) {
	logger.loggers[logger.defaultChannel].UpdateContext(update)
}
