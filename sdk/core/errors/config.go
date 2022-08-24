package errors

import "errors"

const (
	NoConfig             = "no config, please supply a config file or environment variables"
	NoMongoDbConfig      = "no MongoDb config, please supply a config file or environment variables"
	NoCockroachDbConfig  = "no CockroachDb config, please supply a config file or environment variables"
	NoProbesConfig       = "no Probes config, please supply a config file or environment variables"
	NoRedisConfig        = "no Redis config, please supply a config file or environment variables"
	NoNATSConfig         = "no NATS config, please supply a config file or environment variables"
	NoJaegerConfig       = "no Jaeger config, please supply a config file or environment variables"
	NoEventStoreDbConfig = "no EventStoreDb config, please supply a config file or environment variables"
	NoLoggerConfig       = "no Logger config, please supply a config file or environment variables"
	NoServiceConfig      = "no Service config, please supply a config file or environment variables"
	NoHttpConfig         = "no Http config, please supply a config file or environment variables"
	NoGRPCConfig         = "no GRPC config, please supply a config file or environment variables"
	NoKafkaConfig        = "no Kafka config, please supply a config file or environment variables"
	NoProjectionConfig   = "no Projection config, please supply a config file or environment variables"
)

var (
	ErrNoConfig             = errors.New(NoConfig)
	ErrNoMongoDbConfig      = errors.New(NoMongoDbConfig)
	ErrNoCockroachDbConfig  = errors.New(NoCockroachDbConfig)
	ErrNoProbesConfig       = errors.New(NoProbesConfig)
	ErrNoRedisConfig        = errors.New(NoRedisConfig)
	ErrNoNATSConfig         = errors.New(NoNATSConfig)
	ErrNoJaegerConfig       = errors.New(NoJaegerConfig)
	ErrNoEventStoreDbConfig = errors.New(NoEventStoreDbConfig)
	ErrNoLoggerConfig       = errors.New(NoLoggerConfig)
	ErrNoServiceConfig      = errors.New(NoServiceConfig)
	ErrNoHttpConfig         = errors.New(NoHttpConfig)
	ErrNoGRPCConfig         = errors.New(NoGRPCConfig)
	ErrNoKafkaConfig        = errors.New(NoKafkaConfig)
	ErrNoProjectionConfig   = errors.New(NoProjectionConfig)
)
