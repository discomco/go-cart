package config

import (
	"fmt"
	"os"

	"github.com/discomco/go-cart/config/cockroach_db"
	"github.com/discomco/go-cart/config/eventstore_db"
	"github.com/discomco/go-cart/config/grpc"
	"github.com/discomco/go-cart/config/http"
	"github.com/discomco/go-cart/config/jaeger"
	"github.com/discomco/go-cart/config/kafka"
	"github.com/discomco/go-cart/config/logger"
	"github.com/discomco/go-cart/config/mongo_db"
	"github.com/discomco/go-cart/config/nats"
	"github.com/discomco/go-cart/config/probes"
	"github.com/discomco/go-cart/config/redis"
	"github.com/discomco/go-cart/config/service"
	"github.com/discomco/go-cart/core/constants"
	errors2 "github.com/discomco/go-cart/core/errors"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

// Path We give the configuration path its own type, for DI
type Path string

type Config struct {
	Service      *service.Config                 `mapstructure:"service"`
	Logger       *logger.Config                  `mapstructure:"logger"`
	Jaeger       *jaeger.Config                  `mapstructure:"jaeger"`
	GRPC         *grpc.GrpcConfig                `mapstructure:"grpc"`
	Probes       *probes.Config                  `mapstructure:"probes"`
	NATS         *nats.Config                    `mapstructure:"nats"`
	EventStoreDb *eventstore_db.Config           `mapstructure:"eventStoreDb"`
	Projection   *eventstore_db.ProjectionConfig `mapstructure:"projection"`
	Http         *http.HttpConfig                `mapstructure:"http"`
	Redis        *redis.Config                   `mapstructure:"redis"`
	CockroachDb  *cockroach_db.Config            `mapstructure:"cockroach_db"`
	Kafka        *kafka.Config                   `mapstructure:"kafka"`
	MongoDB      *mongo_db.Config                `mapstructure:"mongoDb"`
}

func (cfg *Config) GetServiceConfig() IServiceConfig {
	if cfg.Service == nil {
		panic(errors2.ErrNoServiceConfig)
	}
	return cfg.Service
}

func (cfg *Config) GetMongoDbConfig() IMongoDbConfig {
	if cfg.MongoDB == nil {
		panic(errors2.ErrNoMongoDbConfig)
	}
	return cfg.MongoDB
}

func (cfg *Config) GetCockroachDbConfig() ICockroachDbConfig {
	if cfg.CockroachDb == nil {
		panic(errors2.ErrNoCockroachDbConfig)
	}
	return cfg.CockroachDb
}

func (cfg *Config) GetProbesConfig() IProbesConfig {
	if cfg.Probes == nil {
		panic(errors2.ErrNoProbesConfig)
	}
	return cfg.Probes
}

func (cfg *Config) GetRedisConfig() IRedisConfig {
	if cfg.Redis == nil {
		panic(errors2.ErrNoRedisConfig)
	}
	return cfg.Redis
}

func (cfg *Config) GetNATSConfig() INATSConfig {
	if cfg.NATS == nil {
		panic(errors2.ErrNoNATSConfig)
	}
	return cfg.NATS
}

func (cfg *Config) GetJaegerConfig() IJaegerConfig {
	if cfg.Jaeger == nil {
		panic(errors2.ErrNoJaegerConfig)
	}
	return cfg.Jaeger
}

func (cfg *Config) GetESDBConfig() IESDBConfig {
	if cfg.EventStoreDb == nil {
		panic(errors2.ErrNoEventStoreDbConfig)
	}
	return cfg.EventStoreDb
}

func (cfg *Config) GetLoggerConfig() ILoggerConfig {
	if cfg.Logger == nil {
		panic(errors2.ErrNoLoggerConfig)
	}
	return cfg.Logger
}

func (cfg *Config) GetProjectionConfig() IProjectionConfig {
	if cfg.Projection == nil {
		panic(errors2.ErrNoProjectionConfig)
	}
	return cfg.Projection
}

func (cfg *Config) GetHttpConfig() IHttpConfig {
	if cfg.Http == nil {
		panic(errors2.ErrNoHttpConfig)
	}
	return cfg.Http
}

func (cfg *Config) GetGRPCConfig() IGRPCConfig {
	if cfg.GRPC == nil {
		panic(errors2.ErrNoGRPCConfig)
	}
	return cfg.GRPC
}

func (cfg *Config) GetKafkaConfig() IKafkaConfig {
	if cfg.Kafka == nil {
		panic(errors2.ErrNoKafkaConfig)
	}
	return cfg.Kafka
}

func AppConfig(configPath Path) (IAppConfig, error) {
	if configPath == "" {
		configPathFromEnv := Path(os.Getenv(constants.ConfigPath))
		if configPathFromEnv != "" {
			configPath = configPathFromEnv
		} else {
			getwd, err := os.Getwd()
			if err != nil {
				return nil, errors.Wrap(err, "os.Getwd")
			}
			configPath = Path(fmt.Sprintf("%s/config/config.yaml", getwd))
		}
	}
	cfg := &Config{}
	viper.SetConfigType(constants.Yaml)
	viper.SetConfigFile(string(configPath))
	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.Wrap(err, "viper.ReadInConfig")
	}
	if err := viper.Unmarshal(cfg); err != nil {
		return nil, errors.Wrap(err, "viper.Unmarshal")
	}

	if cfg.GRPC != nil {
		grpcPort := os.Getenv(constants.GrpcPort)
		if grpcPort != "" {
			if cfg.GRPC == nil {
				cfg.GRPC = &grpc.GrpcConfig{}
			}
			cfg.GRPC.Port = grpcPort
		}
	}

	if cfg.Jaeger != nil {
		jaegerAddr := os.Getenv(constants.JaegerHostPort)
		if jaegerAddr != "" {
			if cfg.Jaeger == nil {
				cfg.Jaeger = &jaeger.Config{}
			}
			cfg.Jaeger.HostPort = jaegerAddr
		}
	}

	if cfg.EventStoreDb != nil {
		eventStoreConnectionString := os.Getenv(constants.EventStoreConnectionString)
		if eventStoreConnectionString != "" {
			cfg.EventStoreDb.ConnectionString = eventStoreConnectionString
		}
	}

	if cfg.NATS != nil {
		natsUrl := os.Getenv(constants.NatsUrl)
		if natsUrl != "" {
			if cfg.NATS == nil {
				cfg.NATS = &nats.Config{}
			}
			cfg.NATS.Url = natsUrl
		}
	}

	if cfg.Redis != nil {
		redisURL := os.Getenv(constants.RedisUrl)
		if redisURL != "" {
			if cfg.Redis == nil {
				cfg.Redis = &redis.Config{}
			}
			cfg.Redis.Url = redisURL
		}
	}

	if cfg.MongoDB != nil {
		mongoUri := os.Getenv(constants.MongoDbUri)
		if mongoUri != "" {
			if cfg.MongoDB == nil {
				cfg.MongoDB = &mongo_db.Config{}
			}
			cfg.MongoDB.Uri = mongoUri

			mongoUser := os.Getenv(constants.MongoDbUser)
			if mongoUser != "" {
				cfg.MongoDB.User = mongoUser
			}
			mongoPassword := os.Getenv(constants.MongoDbPassword)
			if mongoPassword != "" {
				cfg.MongoDB.Password = mongoPassword
			}
			mongoAuthMechanism := os.Getenv(constants.MongoDbAuthMechanism)
			if mongoAuthMechanism != "" {
				cfg.MongoDB.AuthMechanism = mongoAuthMechanism
			}
			if cfg.MongoDB.AuthMechanism == "" {
				cfg.MongoDB.AuthMechanism = "SCRAM-SHA-1"
			}
		}
	}
	return cfg, nil
}
