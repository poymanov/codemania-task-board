package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/poymanov/codemania-task-board/board/internal/config"
	"github.com/poymanov/codemania-task-board/board/internal/infrastructure/app"
	"github.com/rs/zerolog/log"
)

const (
	configPath = ".env"
)

func main() {
	err := config.Load(configPath)
	if err != nil {
		panic(fmt.Errorf("failed to load config: %w", err))
	}

	a, err := app.NewApp(context.Background())
	if err != nil {
		log.Fatal().Err(err).Msg("failed to initialize app")
	}

	if err := a.Run(); err != nil {
		log.Fatal().Err(err).Msg("failed to run app")
		return
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	defer func() {
		err := a.Close()
		if err != nil {
			log.Fatal().Err(err).Msg("failed to close app")
		}
	}()
}
