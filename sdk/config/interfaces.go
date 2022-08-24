package config

type IAppConfig interface {
	GetServiceConfig() IServiceConfig
	GetProjectionConfig() IProjectionConfig
	GetNATSConfig() INATSConfig
	GetHttpConfig() IHttpConfig
	GetProbesConfig() IProbesConfig
	GetGRPCConfig() IGRPCConfig
	GetLoggerConfig() ILoggerConfig
	GetRedisConfig() IRedisConfig
	GetESDBConfig() IESDBConfig
	GetJaegerConfig() IJaegerConfig
	GetCockroachDbConfig() ICockroachDbConfig
	GetKafkaConfig() IKafkaConfig
	GetMongoDbConfig() IMongoDbConfig
}

type IServiceConfig interface {
	GetNamespace() string
	GetSubSystem() string
	GetServiceName() string
}

type IKafkaConfig interface {
	GetBootstrapServers() string
	GetGroupId() string
	GetAutoOffsetReset() string
	GetRetentionMs() string
}

type IHttpConfig interface {
	GetPort() string
	IsDevelopment() bool
	GetBasePath() string
	GetAppPath() string
	WithDebugErrorsResponse() bool
	GetIgnoreLogUrls() []string
}

type IGRPCConfig interface {
	GetPort() string
	IsDevelopment() bool
}

type IProbesConfig interface {
	GetReadinessPath() string
	GetLivenessPath() string
	GetPort() string
	GetPProf() string
	GetPrometheusPath() string
	GetCheckIntervalSeconds() int
}

type IProjectionConfig interface {
	GetPoolSize() int
	GetEventPrefix() string
	GetName() string
	GetGroup() string
}

type ILoggerConfig interface {
	GetLogLevel() string
	IsDevelopment() bool
	GetEncoder() string
}

type IJaegerConfig interface {
	GetServiceName() string
	GetHostPort() string
	IsEnabled() bool
	UseLogSpans() bool
}

//IMongoDbConfig is the interface for the mongo config.
type IMongoDbConfig interface {
	GetUri() string
	GetUser() string
	GetPassword() string
	GetAuthMechanism() string
}

type ICockroachDbConfig interface {
	GetDSN() string
}

// IRedisConfig is an interface to the Redis Configuration
type IRedisConfig interface {
	GetUrl() string
}

// IESDBConfig is the configuration interface for EventStoreDB
type IESDBConfig interface {
	GetConnectionString() string
	GetUser() string
	GetPwd() string
}

// INATSConfig is and interface to nats configuration
type INATSConfig interface {
	GetUrl() string
	GetUser() string
	GetPwd() string
}
