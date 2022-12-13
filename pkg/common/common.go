package common

import (
	"context"

	env "github.com/Netflix/go-env"
	"github.com/rs/zerolog/log"
)

const (
	DEBUG = iota
	TEST
	RELEASE
)

var ReleaseModes = map[int]string{
	0: "debug",
	1: "test",
	2: "release",
}

func ReadConfigFromEnv(ctx context.Context, cfg any) {
	_, err := env.UnmarshalFromEnviron(cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("error while reading config from env")
	}
}
