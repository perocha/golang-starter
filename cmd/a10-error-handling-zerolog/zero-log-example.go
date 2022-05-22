package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Print("This is a test of zerolog logging package")

	log.Debug().
		Int("error code", 101).
		Msg("this is a generic message")

	log.Debug().Str("str1", "str2").Send()
}
