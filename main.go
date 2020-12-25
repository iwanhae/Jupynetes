package main

//go:generate oapi-codegen --package=server --generate types -o pkg/server/types.gen.go api/api.yaml
//go:generate oapi-codegen --package=server --generate spec -o pkg/server/spec.gen.go api/api.yaml

import (
	"net/http"
	"os"
	"time"

	"github.com/iwanhae/Jupynetes/pkg/config"
	"github.com/iwanhae/Jupynetes/pkg/kubeclient"
	"github.com/iwanhae/Jupynetes/pkg/server"
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

	// Initializing pkgs
	kubeclient.Init(c)
	r := server.InitRouter(c)

	// Listening
	log.Info().Msg("Listening on :3000")
	http.ListenAndServe(":3000", r)
}
