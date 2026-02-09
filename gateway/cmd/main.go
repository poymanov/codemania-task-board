package main

import (
	"github.com/poymanov/codemania-task-board/gateway/app"
	"github.com/rs/zerolog/log"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatal().Err(err).Msg("failed to run app")
	}
}
