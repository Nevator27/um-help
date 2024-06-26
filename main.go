package main

import (
	"os"
	"time"

	"github.com/Nevator27/um-help/config"
	"github.com/Nevator27/um-help/logger"
	"github.com/Nevator27/um-help/repo"
	"github.com/Nevator27/um-help/server"
	"github.com/Nevator27/um-help/server/controller"
	"github.com/Nevator27/um-help/service"
	"github.com/rs/zerolog"
)

func main() {
	cfg := config.Get()
	logger := logger.New(cfg)

	repo, err := repo.New(cfg)
	if err != nil {
		end(logger, err, "failed to initialize repo")
	}

	svc, err := service.New(cfg, logger, repo)
	if err != nil {
		end(logger, err, "failed to initialize svc")
	}

	ctrl := controller.New(svc, logger)

	svr := server.New(cfg, logger, ctrl)

	if err := svr.Start(); err != nil {
		end(logger, err, "failed to start server")
	}
}

func end(logger *zerolog.Logger, err error, message string) {
	logger.Fatal().Err(err).Msgf("%+v: %+v", message, err)
	time.Sleep(time.Millisecond * 50)

	os.Exit(1)
}
