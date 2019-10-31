package main

import (
	"github.com/rs/zerolog"
	"github.com/hookover/logging"
)

func init() {
	conf := &logging.Conf{DefaultLogFile: "./logs/app.log"}

	conf.Channels = append(conf.Channels,
		&logging.Channel{
			Name:    "sql",
			LogFile: "./logs/sql.log",
		},
		&logging.Channel{
			Name:    "console",
			Days:    3,
			Level:   zerolog.WarnLevel,
			LogFile: "./logs/console.log",
		},
	)

	if _, err := logging.Initialization(conf); err != nil {
		panic(err)
	}
}

func main() {
	logging.Info().Msg("test")
	logging.Chan("sql").Warn().Msg("sql warning")
	logging.Chan("console").Error().Msg("console error")
	logging.Chan("console").Debug().Msg("console debug")
}
