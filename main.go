package main

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/savi2w/nano-go/config"
	"github.com/savi2w/nano-go/logger"
	"github.com/savi2w/nano-go/repo"
	"github.com/savi2w/nano-go/server"
	"github.com/savi2w/nano-go/server/controller"
	"github.com/savi2w/nano-go/service"
)

func main() {
	cfg := config.Get()
	logger := logger.New(cfg)

	repo, err := repo.New(cfg)
	if err != nil {
		end(logger, err, "failed to initialize repo manager")
	}

	svc := service.New(cfg, logger, repo)
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
