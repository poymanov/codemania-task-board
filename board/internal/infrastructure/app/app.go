package app

import (
	"fmt"
	"log"
	"net"

	"github.com/poymanov/codemania-task-board/board/internal/config"
	transportBoardV1 "github.com/poymanov/codemania-task-board/board/internal/transport/grpc/board/v1"
	boardV1 "github.com/poymanov/codemania-task-board/shared/pkg/proto/board/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type App struct {
	closer []func() error
}

func NewApp() *App {
	return &App{
		closer: make([]func() error, 2),
	}
}

func (a *App) Run() error {
	return a.runGrpcServer()
}

func (a *App) runGrpcServer() error {
	lis, err := net.Listen("tcp", config.AppConfig().Grpc.Address())
	if err != nil {
		return err
	}

	boardService := transportBoardV1.NewBoardService()

	s := grpc.NewServer()

	boardV1.RegisterBoardServiceServer(s, boardService)

	reflection.Register(s)

	go func() {
		log.Printf("🚀 gRPC server listening on %s\n", config.AppConfig().Grpc.Address())
		err = s.Serve(lis)
		if err != nil {
			log.Printf("failed to serve: %v\n", err)
			return
		}
	}()

	a.closer = append(a.closer, func() error {
		if cerr := lis.Close(); cerr != nil {
			return fmt.Errorf("failed to close listener: %w", cerr)
		}

		return nil
	})

	a.closer = append(a.closer, func() error {
		s.GracefulStop()

		return nil
	})

	return nil
}

func (a *App) Close() error {
	for _, closer := range a.closer {
		if err := closer(); err != nil {
			return err
		}
	}

	return nil
}
