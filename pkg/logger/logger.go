package logger

import (
	"context"
	"os"

	"github.com/rs/zerolog/pkgerrors"

	"github.com/devminnu/learn-rest/product/pkg/common"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type LoggerConfig struct {
	ReleaseMode int32 `env:"RELEASE_MODE"`
}

func Init(ctx context.Context) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = zerolog.New(os.Stderr).With().Timestamp().Caller().Logger().Hook(Caller{})
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	loggerConfig := new(LoggerConfig)
	common.ReadConfigFromEnv(ctx, loggerConfig)
	switch loggerConfig.ReleaseMode {
	case common.TEST:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
}

type Caller struct{}

func (c Caller) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	// e.CallerSkipFrame(3).Caller()
}
