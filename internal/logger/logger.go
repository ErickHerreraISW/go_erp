package logger

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Setup(appEnv string) {
	if appEnv == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	} else {
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	}
}
