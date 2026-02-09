package app

import (
	"context"
	"errors"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/poymanov/codemania-task-board/board/internal/config"
	boardRepository "github.com/poymanov/codemania-task-board/board/internal/infrastructure/persistance/repository/board"
	transportBoardV1 "github.com/poymanov/codemania-task-board/board/internal/transport/grpc/board/v1"
	boardUseCase "github.com/poymanov/codemania-task-board/board/internal/usecase/board"
	"github.com/poymanov/codemania-task-board/platform/pkg/grpc/health"
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
	config           *config.Config
}

const (
	configPath = ".env"
)

func newApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.InitDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func Run() error {
	ctx := context.Background()

	a, err := newApp(ctx)

	defer func() {
		ec := a.Close()
		if ec != nil {
			log.Error().Err(ec).Msg("failed to close app")
			return
		}
	}()

	if err != nil {
		return err
	}

	err = a.runMigrator()
	if err != nil {
		return err
	}

	a.runGrpcServer()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	return nil
}

func (a *App) InitConfig(_ context.Context) error {
	cfg, err := config.Load(configPath)
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	a.config = cfg

	return nil
}

func (a *App) InitDeps(ctx context.Context) error {
	inits := []func(ctx context.Context) error{
		a.InitConfig,
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
	pool, err := pgxpool.New(ctx, a.config.Db.Uri())
	if err != nil {
		log.Fatal().Err(err).Msg("db failed connect")
		return err
	}

	err = pool.Ping(ctx)
	if err != nil {
		log.Fatal().Err(err).Msg("db not available")
		return err

	}

	a.dbConnectionPool = pool

	a.closer = append(a.closer, func() error {
		pool.Close()

		return nil
	})

	return nil
}

func (a *App) initListener(_ context.Context) error {
	list, err := net.Listen("tcp", a.config.Grpc.Address())
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
	logger.InitLogger(a.config.Logger.Level())

	return nil
}

func (a *App) runMigrator() error {
	migration := migrator.NewMigrator(a.dbConnectionPool, a.config.Db.MigrationDirectory())

	if err := migration.Up(); err != nil {
		return err
	}

	return nil
}

func (a *App) runGrpcServer() {
	br := boardRepository.NewRepository(a.dbConnectionPool)
	bus := boardUseCase.NewUseCase(br)
	boardService := transportBoardV1.NewBoardService(bus)

	s := grpc.NewServer()

	boardV1.RegisterBoardServiceServer(s, boardService)
	health.RegisterService(s)

	reflection.Register(s)

	go func() {
		log.Info().Msg(fmt.Sprintf("🚀 gRPC server listening on %s\n", a.config.Grpc.Address()))
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
}

func (a *App) Close() error {
	for _, closer := range a.closer {
		if err := closer(); err != nil {
			return err
		}
	}

	return nil
}
