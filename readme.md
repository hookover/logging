# logging
zerolog + file-rotatelogs 简单封装 


```go

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


```

日志文件

```go
-rw-r--r--  1 hookover  staff    69B Oct 31 17:15 app-2019-10-31.log
-rw-r--r--  1 hookover  staff    79B Oct 31 17:15 console-2019-10-31.log
-rw-r--r--  1 hookover  staff    76B Oct 31 17:15 sql-2019-10-31.log

```