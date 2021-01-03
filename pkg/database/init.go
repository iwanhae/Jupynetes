package database

import (
	"context"

	_ "github.com/go-sql-driver/mysql" //mysql plugin
	"github.com/rs/zerolog/log"

	"github.com/iwanhae/Jupynetes/ent"
	"github.com/iwanhae/Jupynetes/ent/user"
	"github.com/iwanhae/Jupynetes/pkg/config"
)

//DBType ent support mysql, sqlite and PostgreSQL but this project will use MySQL
const DBType = "mysql"

var client = &ent.Client{}
var salt = ""

//Init initialize database
func Init(ctx context.Context, c *config.Configs) {
	salt = c.Database.Salt

	var err error
	log.Ctx(ctx).Info().Msg("connecting to db.....")
	client, err = ent.Open(DBType, c.Database.URI+"?parseTime=true")
	if err != nil {
		log.Fatal().Msgf("failed opening connection to mysql: %v", err)
	}

	log.Ctx(ctx).Info().Msg("migrating db.....")
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatal().Msgf("failed creating schema resources: %v", err)
	}

	createAdminUser(ctx)
}

func createAdminUser(ctx context.Context) {
	num, err := client.User.Query().
		Where(user.UserIDEQ("admin")).
		Count(ctx)
	if err != nil {
		log.Fatal().Msgf("fail to query admin user to db")
	}
	if num != 0 {
		log.Ctx(ctx).Info().Msg("admin user already exsists")
		return
	}
	log.Ctx(ctx).Info().Msg("admin user not detected creating one.....")
	user := client.User.Create().
		SetUserID("admin").
		SetUserPw(EncryptPassword("admin")).
		SetQuotaInstance(-1).
		SetQuotaCPU(-1).
		SetQuotaMemory(-1).
		SetQuotaNvidiaGpu(-1).
		SetQuotaStorage(-1).SaveX(ctx)
	log.Ctx(ctx).Info().Interface("admin", user).Msg("admin user created")
}

//Close close all connections
func Close(ctx context.Context) {
	if client != nil {
		client.Close()
	}
}
