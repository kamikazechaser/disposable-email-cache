package main

import (
	"github.com/kamikazechaser/disposable-email-cache/internal/server"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	log.Info().Msg("main: starting server on port 5000")

	if err := server.Start(); err != nil {
		log.Fatal().Err(err).Msg("main: could not start server")
	}
}
