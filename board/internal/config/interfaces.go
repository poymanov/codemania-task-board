package config

type GrpcConfig interface {
	Address() string
}

type LoggerConfig interface {
	Level() string
}
