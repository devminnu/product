package postgres

import (
	"context"
	"fmt"

	"github.com/devminnu/learn-rest/product/pkg/common"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

type DatabaseConfig struct {
	Host     string `env:"POSTGRES_HOST,default=localhost"`
	Port     string `env:"POSTGRES_PORT,default=5432"`
	UserName string `env:"POSTGRES_USER,required=true"`
	Password string `env:"POSTGRES_PASSWORD,required=true"`
	DBName   string `env:"POSTGRES_DB,required=true"`
	SSLMode  string `env:"POSTGRES_SSL_MODE,default=disable"`
}

var (
	driver                   = "postgres"
	dbConnectionStringFormat = "host=%s port=%s user=%s password=%s dbname=%s sslmode=%s"
)

func Connect(ctx context.Context) *sqlx.DB {
	v, err := sqlx.Open(driver, dbConnectionString(ctx))
	if err != nil {
		log.Fatal().Err(err).Msgf("error connecting %v database", driver)
	}
	err = v.Ping()
	if err != nil {
		log.Fatal().Err(err).Msgf("error pinging %v database", driver)
	}
	log.Info().Msgf("%v db connection successfull", driver)

	return v
}

func dbConnectionString(ctx context.Context) (connectionString string) {
	dbConfig := new(DatabaseConfig)
	common.ReadConfigFromEnv(ctx, dbConfig)
	log.Info().Interface("dbConfig", dbConfig).Msg("postres db config")

	connectionString = fmt.Sprintf(dbConnectionStringFormat,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.UserName,
		dbConfig.Password,
		dbConfig.DBName,
		dbConfig.SSLMode)

	log.Info().Str("connectionString", connectionString).Msg("postgres db connection string")

	return
}
