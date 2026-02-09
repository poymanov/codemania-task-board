package env

import (
	"fmt"
	"net"
	"time"

	"github.com/caarlos0/env/v11"
)

type httpEnvConfig struct {
	Host        string `env:"GATEWAY_HTTP_HOST,required"`
	Port        string `env:"GATEWAY_HTTP_PORT,required"`
	ReadTimeout string `env:"GATEWAY_HTTP_READ_TIMEOUT,required"`
}

type HttpConfig struct {
	raw httpEnvConfig
}

func NewHttpConfig() (*HttpConfig, error) {
	var raw httpEnvConfig

	if err := env.Parse(&raw); err != nil {
		return nil, err
	}

	return &HttpConfig{raw: raw}, nil
}

func (cfg *HttpConfig) Address() string {
	return net.JoinHostPort(cfg.raw.Host, cfg.raw.Port)
}

func (cfg *HttpConfig) ReadTimeout() time.Duration {
	timeDuration, err := time.ParseDuration(cfg.raw.ReadTimeout)
	if err != nil {
		panic(fmt.Errorf("failed to to parse http timeout config: %w", err))
	}

	return timeDuration * time.Second
}
