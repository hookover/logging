package logging

import (
	"errors"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/rs/zerolog"
	"strings"
	"time"
)

var (
	defaultLogFile string        = "./logs/app.log"
	defaultMaxDay  time.Duration = 7
	defaultFormat  string        = "%Y-%m-%d"
	defaultLevel   zerolog.Level = zerolog.DebugLevel

	logger *Logging
)

type Logging struct {
	defaultChannel string
	conf           *Conf
	loggers        map[string]*zerolog.Logger
}

type Channel struct {
	Name    string
	LogFile string
	Format  string
	Level   zerolog.Level
	Days    time.Duration
}

type Conf struct {
	Channels       []*Channel
	DefaultLogFile string
}

func Initialization(config *Conf) (*Logging, error) {
	if config.DefaultLogFile != "" {
		defaultLogFile = config.DefaultLogFile
	}

	if logger == nil {
		logger = &Logging{
			defaultChannel: "default",
			loggers:        map[string]*zerolog.Logger{},
		}
	}

	for _, channel := range config.Channels {
		if channel.Name == "" {
			return nil, errors.New("channel name params can not be null")
		}

		if channel.LogFile == "" {
			channel.LogFile = defaultLogFile
		}

		if channel.Format == "" {
			channel.Format = defaultFormat
		}

		if channel.Days == 0 {
			channel.Days = defaultMaxDay
		}

		writer, err := rotate(channel.Days*time.Hour*24, channel.LogFile, channel.Format)
		if err != nil {
			return nil, err
		}

		l := zerolog.New(writer).Level(channel.Level).With().Timestamp().Logger()

		logger.loggers[channel.Name] = &l
	}

	//set default logger
	if _, ok := logger.loggers[logger.defaultChannel]; !ok {
		writer, err := rotate(7*time.Hour*24, defaultLogFile, defaultFormat)
		if err != nil {
			return nil, err
		}
		l := zerolog.New(writer).Level(zerolog.Level(defaultLevel)).With().Timestamp().Logger()

		logger.loggers[logger.defaultChannel] = &l
	}

	return logger, nil
}

func rotate(maxAge time.Duration, logPath string, format string) (r *rotatelogs.RotateLogs, err error) {
	tmp := strings.Split(logPath, ".")
	if tmp[len(tmp)-1] == "log" {
		tmp = tmp[:len(tmp)-1]
	}

	logPath = strings.Join(tmp, ".") + "-" + format + ".log"
	r, err = rotatelogs.New(
		logPath,
		rotatelogs.WithMaxAge(maxAge),
		rotatelogs.WithRotationTime(time.Hour*24),
	)
	return
}
