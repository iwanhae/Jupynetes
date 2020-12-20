package main

import (
	"os"
	"time"

	"git.iwanhae.kr/wan/jupynetes/pkg/config"
	"git.iwanhae.kr/wan/jupynetes/pkg/kubeclient"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Msg("Hello")

	c := config.GetConfigs()

	// Set Logging Format
	if c.Deploy == config.EnvDeployProd {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC1123})
	}
	log.Logger = log.Logger.With().Caller().Timestamp().Logger()

	log.Info().Interface("configs", c).Msg("Configuration Loaded")
	kubeclient.Init(c)
}
