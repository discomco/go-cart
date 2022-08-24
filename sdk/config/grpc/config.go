package grpc

type GrpcConfig struct {
	Port        string `mapstructure:"port"`
	Development bool   `mapstructure:"development"`
}

func (G *GrpcConfig) GetPort() string {
	return G.Port
}

func (G *GrpcConfig) IsDevelopment() bool {
	return G.Development
}
