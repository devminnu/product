package mongo

import (
	"context"
	"fmt"

	"github.com/devminnu/learn-rest/product/pkg/common"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type DatabaseConfig struct {
	Host     string `env:"MONGO_HOST,default=localhost"`
	Port     string `env:"MONGO_PORT,default=27017"`
	UserName string `env:"MONGO_USER,required=true"`
	Password string `env:"MONGO_PASSWORD,required=true"`
	DBName   string `env:"MONGO_DB,required=true"`
}

const dbConnectionStringFormat = "mongodb://%s:%s@%s:%s/?maxPoolSize=20&w=majority"

func Connect(ctx context.Context) *mongo.Database {
	dbConfig := new(DatabaseConfig)

	// Create a new client and connect to the server
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbConnectionString(ctx, dbConfig)))
	if err != nil {
		log.Error().Err(err).Msg("error connecting mongo db")
		panic(err)
	}
	// Ping the primary
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Error().Err(err).Msg("error pinging mongo db")
		panic(err)
	}

	return client.Database(dbConfig.DBName)
}

func dbConnectionString(ctx context.Context, dbConfig *DatabaseConfig) (connectionString string) {
	common.ReadConfigFromEnv(ctx, dbConfig)
	log.Info().Interface("dbConfig", dbConfig).Msg("mongo db config")

	connectionString = fmt.Sprintf(dbConnectionStringFormat,
		dbConfig.UserName,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
	)

	log.Info().Str("connectionString", connectionString).Msg("mongo db connection string")

	return
}
