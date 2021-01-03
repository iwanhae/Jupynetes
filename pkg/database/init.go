package database

import (
	"context"

	_ "github.com/go-sql-driver/mysql" //mysql plugin
	"github.com/rs/zerolog/log"

	"github.com/iwanhae/Jupynetes/ent"
	"github.com/iwanhae/Jupynetes/pkg/config"
)

//DBType ent support mysql, sqlite and PostgreSQL but this project will use MySQL
const DBType = "mysql"

var client = &ent.Client{}

//Init initialize database
func Init(ctx context.Context, c *config.Configs) {

	var err error
	log.Ctx(ctx).Info().Msg("connecting to db.....")
	client, err = ent.Open(DBType, c.Database.URI)
	if err != nil {
		log.Fatal().Msgf("failed opening connection to mysql: %v", err)
	}

	log.Ctx(ctx).Info().Msg("migrating db.....")
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatal().Msgf("failed creating schema resources: %v", err)
	}
}

//Close close all connections
func Close(ctx context.Context) {
	if client != nil {
		client.Close()
	}
}
