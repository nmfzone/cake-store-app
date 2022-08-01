package logger

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"sync"
	"time"
)

var singleton sync.Once

func InitLogger() {
	singleton.Do(func() {
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
		output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}

		log.Logger = log.Output(output)
	})
}
