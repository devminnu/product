package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/devminnu/learn-rest/product/pkg/common"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type ServerConfig struct {
	ReleaseMode  int32         `env:"REST_RELEASE_MODE"`
	ServiceName  string        `env:"REST_SERVICE_NAME"`
	Port         string        `env:"REST_PORT,default=8080"`
	ReadTimeout  time.Duration `env:"REST_READ_TIMEOUT,default=10s"`
	WriteTimeout time.Duration `env:"REST_WRITE_TIMEOUT,default=10s"`
}

func Run(ctx context.Context, registerRestHandlers func(ctx context.Context, r *gin.Engine)) {
	// read server config
	serverConfig := new(ServerConfig)
	common.ReadConfigFromEnv(ctx, serverConfig)
	log.Info().Interface("config", serverConfig).Msg("server config")

	router := gin.New()

	// setup server
	s := &http.Server{
		Addr:           fmt.Sprintf(":%s", serverConfig.Port),
		Handler:        router,
		ReadTimeout:    serverConfig.ReadTimeout,
		WriteTimeout:   serverConfig.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	// register middlewares
	// router.Use(middleware.Logger())

	//register ping health status route
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, map[string]string{
			"service": serverConfig.ServiceName,
			"status":  "running",
		})
	})

	// register other handlers
	registerRestHandlers(ctx, router)

	// start server
	s.ListenAndServe()
}
