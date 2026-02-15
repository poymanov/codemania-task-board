package app

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/adaptor"
	loggerMiddleware "github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/poymanov/codemania-task-board/gateway/internal/config"
	boardGrpcClientV1 "github.com/poymanov/codemania-task-board/gateway/internal/transport/grpc/client/board/v1/board"
	columnGrpcClientV1 "github.com/poymanov/codemania-task-board/gateway/internal/transport/grpc/client/board/v1/column"
	apiV1 "github.com/poymanov/codemania-task-board/gateway/internal/transport/http/gateway/v1"
	createBoardUseCase "github.com/poymanov/codemania-task-board/gateway/internal/usecase/board/create"
	getAllBoardUseCase "github.com/poymanov/codemania-task-board/gateway/internal/usecase/board/get_all"
	createColumnUseCase "github.com/poymanov/codemania-task-board/gateway/internal/usecase/column/create"
	"github.com/poymanov/codemania-task-board/platform/pkg/logger"
	gatewayV1 "github.com/poymanov/codemania-task-board/shared/pkg/openapi/gateway/v1"
	boardV1 "github.com/poymanov/codemania-task-board/shared/pkg/proto/board/v1"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const defaultInitializationTimeout = time.Second * 10

type App struct {
	closer       []func() error
	config       *config.Config
	boardClient  *boardGrpcClientV1.BoardClient
	columnClient *columnGrpcClientV1.Client
}

func newApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
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

func (a *App) initConfig(ctx context.Context) error {
	if _, ok := ctx.Deadline(); !ok {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, defaultInitializationTimeout)
		defer cancel()
	}

	chDone := make(chan struct{})

	var configErr error

	go func() {
		configPath := flag.String("env", ".env", "path to .env file")
		flag.Parse()

		cfg, err := config.Load(*configPath)

		if err != nil {
			configErr = fmt.Errorf("failed to load config: %w, config path: %s", err, *configPath)
		} else {
			a.config = cfg
		}

		chDone <- struct{}{}
	}()

	select {
	case <-ctx.Done():
		configErr = fmt.Errorf("config initialization timed out")
	case <-chDone:
	}

	if configErr != nil {
		return configErr
	}

	return nil
}

func (a *App) initGrpcClients(ctx context.Context) error {
	if _, ok := ctx.Deadline(); !ok {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, defaultInitializationTimeout)
		defer cancel()
	}

	chDone := make(chan struct{})

	var clientErr error

	go func() {
		conn, err := grpc.NewClient(
			a.config.GrpcClient.BoardAddress(),
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		if err != nil {
			clientErr = fmt.Errorf("failed to connect grpc: %w", err)
		}

		boardServiceClient := boardV1.NewBoardServiceClient(conn)
		columnServiceClient := boardV1.NewColumnServiceClient(conn)

		a.boardClient = boardGrpcClientV1.NewClient(boardServiceClient)
		a.columnClient = columnGrpcClientV1.NewClient(columnServiceClient)

		a.closer = append(a.closer, func() error {
			if cerr := conn.Close(); cerr != nil {
				return cerr
			}

			return nil
		})

		chDone <- struct{}{}
	}()

	select {
	case <-ctx.Done():
		clientErr = fmt.Errorf("gRPC clients initialization timed out")
	case <-chDone:
	}

	if clientErr != nil {
		return clientErr
	}

	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	if _, ok := ctx.Deadline(); !ok {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, defaultInitializationTimeout)
		defer cancel()
	}

	inits := []func(ctx context.Context) error{
		a.initConfig,
		a.initLogger,
		a.initGrpcClients,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initLogger(_ context.Context) error {
	logger.InitLogger(a.config.Logger.Level())

	return nil
}

func (a *App) runHttpServer() error {
	cbuc := createBoardUseCase.NewUseCase(a.boardClient)
	galbuc := getAllBoardUseCase.NewUseCase(a.boardClient)
	ccuc := createColumnUseCase.NewUseCase(a.columnClient)

	api := apiV1.NewApi(cbuc, galbuc, ccuc)

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
