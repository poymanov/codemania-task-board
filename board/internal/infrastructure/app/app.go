package app

import (
	"context"
	"errors"
	"fmt"
	"net"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/poymanov/codemania-task-board/board/internal/config"
	boardRepository "github.com/poymanov/codemania-task-board/board/internal/infrastructure/persistance/repository/board"
	transportBoardV1 "github.com/poymanov/codemania-task-board/board/internal/transport/grpc/board/v1"
	boardUseCase "github.com/poymanov/codemania-task-board/board/internal/usecase/board"
	"github.com/poymanov/codemania-task-board/platform/pkg/logger"
	"github.com/poymanov/codemania-task-board/platform/pkg/migrator"
	boardV1 "github.com/poymanov/codemania-task-board/shared/pkg/proto/board/v1"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type App struct {
	closer           []func() error
	listener         net.Listener
	dbConnectionPool *pgxpool.Pool
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{
		closer: make([]func() error, 0),
	}

	err := a.InitDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Run() error {
	migration := migrator.NewMigrator(a.dbConnectionPool, config.AppConfig().Db.MigrationDirectory())

	if err := migration.Up(); err != nil {
		return err
	}

	return a.runGrpcServer()
}

func (a *App) InitDeps(ctx context.Context) error {
	inits := []func(ctx context.Context) error{
		a.InitLogger,
		a.InitDB,
		a.initListener,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) InitDB(ctx context.Context) error {
	pool, err := pgxpool.New(ctx, config.AppConfig().Db.Uri())
	if err != nil {
		panic(fmt.Errorf("db failed connect: %w", err))
	}

	err = pool.Ping(ctx)
	if err != nil {
		panic(fmt.Errorf("db not available: %w", err))
	}

	a.dbConnectionPool = pool

	a.closer = append(a.closer, func() error {
		pool.Close()

		return nil
	})

	return nil
}

func (a *App) initListener(_ context.Context) error {
	list, err := net.Listen("tcp", config.AppConfig().Grpc.Address())
	if err != nil {
		log.Fatal().Err(err).Msg("failed to start listener")
		return err
	}

	a.listener = list

	a.closer = append(a.closer, func() error {
		lerr := list.Close()

		if lerr != nil && !errors.Is(lerr, net.ErrClosed) {
			log.Fatal().Err(err).Msg("failed to close listener")

			return lerr
		}

		return nil
	})

	return nil
}

func (a *App) InitLogger(_ context.Context) error {
	logger.InitLogger(config.AppConfig().Logger.Level())

	return nil
}

func (a *App) runGrpcServer() error {
	br := boardRepository.NewRepository(a.dbConnectionPool)
	bus := boardUseCase.NewUseCase(br)
	boardService := transportBoardV1.NewBoardService(bus)

	s := grpc.NewServer()

	boardV1.RegisterBoardServiceServer(s, boardService)

	reflection.Register(s)

	go func() {
		log.Printf("🚀 gRPC server listening on %s\n", config.AppConfig().Grpc.Address())
		err := s.Serve(a.listener)
		if err != nil {
			log.Fatal().Err(err).Msg("failed to serve grpc server")
			return
		}
	}()

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
