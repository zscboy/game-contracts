package client

const (
	NetworkIDMainnet  = 314
	NetworkIDCalibnet = 314159

	EndpointMainnet  = "https://api.node.glif.io/"
	EndpointCalibnet = "https://api.calibration.node.glif.io/"
)

type Config struct {
	endpoint   string
	networkID  int64
	privateKey string
}

type Option func(opts *Config)

func EndpointOption(endpoint string) Option {
	return func(opts *Config) {
		opts.endpoint = endpoint
	}
}

func NetworkIDOption(id int64) Option {
	return func(opts *Config) {
		opts.networkID = id
	}
}

func PrivateKeyOption(privateKey string) Option {
	return func(opts *Config) {
		opts.privateKey = privateKey
	}
}

func defaultConfig() Config {
	return Config{
		endpoint:  EndpointCalibnet,
		networkID: NetworkIDCalibnet,
	}
}
