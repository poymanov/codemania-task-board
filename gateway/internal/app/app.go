package app

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/adaptor"
	loggerMiddleware "github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/poymanov/codemania-task-board/gateway/internal/config"
	boardGrpcClientV1 "github.com/poymanov/codemania-task-board/gateway/internal/transport/grpc/client/board/v1"
	apiV1 "github.com/poymanov/codemania-task-board/gateway/internal/transport/http/gateway/v1"
	createBoardUseCase "github.com/poymanov/codemania-task-board/gateway/internal/usecase/board/create"
	"github.com/poymanov/codemania-task-board/platform/pkg/logger"
	gatewayV1 "github.com/poymanov/codemania-task-board/shared/pkg/openapi/gateway/v1"
	boardV1 "github.com/poymanov/codemania-task-board/shared/pkg/proto/board/v1"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type App struct {
	closer      []func() error
	config      *config.Config
	boardClient *boardGrpcClientV1.BoardClient
}

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
	if err != nil {
		return err
	}

	defer func() {
		ec := a.Close()
		if ec != nil {
			log.Error().Err(ec).Msg("failed to close app")
			return
		}
	}()

	err = a.runHttpServer()
	if err != nil {
		return err
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	return nil
}

func (a *App) InitConfig(_ context.Context) error {
	configPath := flag.String("env", ".env", "path to .env file")

	flag.Parse()

	cfg, err := config.Load(*configPath)
	if err != nil {
		return fmt.Errorf("failed to load config: %w, config path: %s", err, *configPath)
	}

	a.config = cfg

	return nil
}

func (a *App) InitBoardClient(_ context.Context) error {
	conn, err := grpc.NewClient(
		a.config.GrpcClient.BoardAddress(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return fmt.Errorf("failed to connect board grpc: %w", err)
	}

	boardServiceClient := boardV1.NewBoardServiceClient(conn)

	a.boardClient = boardGrpcClientV1.NewClient(boardServiceClient)

	a.closer = append(a.closer, func() error {
		if cerr := conn.Close(); cerr != nil {
			return cerr
		}

		return nil
	})

	return nil
}

func (a *App) InitDeps(ctx context.Context) error {
	inits := []func(ctx context.Context) error{
		a.InitConfig,
		a.InitLogger,
		a.InitBoardClient,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) InitLogger(_ context.Context) error {
	logger.InitLogger(a.config.Logger.Level())

	return nil
}

func (a *App) runHttpServer() error {
	cbUseCase := createBoardUseCase.NewUseCase(a.boardClient)
	api := apiV1.NewApi(cbUseCase)

	gatewayServer, err := gatewayV1.NewServer(api)
	if err != nil {
		return err
	}

	app := fiber.New(fiber.Config{
		ReadTimeout: a.config.Http.ReadTimeout(),
	})
	app.Use(loggerMiddleware.New())
	app.Use("/", adaptor.HTTPHandler(gatewayServer))

	go func() {
		if err := app.Listen(a.config.Http.Address()); err != nil {
			log.Fatal().Err(err).Msg("failed to serve http server")
		}
	}()

	a.closer = append(a.closer, func() error {
		esh := app.Shutdown()
		if esh != nil {
			return esh
		}

		return nil
	})

	return nil
}

func (a *App) Close() error {
	for _, closer := range a.closer {
		if err := closer(); err != nil {
			log.Fatal().Err(err).Msg("failed to close application component")
		}
	}

	return nil
}
