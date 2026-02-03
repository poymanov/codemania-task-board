package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/poymanov/codemania-task-board/board/internal/config"
	"github.com/poymanov/codemania-task-board/board/internal/infrastructure/app"
)

const (
	configPath = ".env"
)

func main() {
	err := config.Load(configPath)
	if err != nil {
		panic(fmt.Errorf("failed to load config: %w", err))
	}

	a := app.NewApp()

	if err := a.Run(); err != nil {
		log.Fatal(err)
		return
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	if err := a.Close(); err != nil {
		log.Fatal(err)
	}
}
