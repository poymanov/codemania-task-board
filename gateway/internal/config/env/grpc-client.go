package env

import (
	"net"

	"github.com/caarlos0/env/v11"
)

type grpcClientBoardEnvConfig struct {
	Host string `env:"BOARD_GRPC_HOST,required"`
	Port string `env:"BOARD_GRPC_PORT,required"`
}

type GrpcClient struct {
	rawBoard grpcClientBoardEnvConfig
}

func NewGrpcClient() (*GrpcClient, error) {
	var rawBoard grpcClientBoardEnvConfig

	if err := env.Parse(&rawBoard); err != nil {
		return nil, err
	}

	return &GrpcClient{
		rawBoard: rawBoard,
	}, nil
}

func (cfg *GrpcClient) BoardAddress() string {
	return net.JoinHostPort(cfg.rawBoard.Host, cfg.rawBoard.Port)
}
