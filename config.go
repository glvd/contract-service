package service

type DatabaseConfig struct {
	Addr     string `json:"addr"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// Config ...
type Config struct {
	RestPort        string         `json:"rest_port"`
	RPCPort         string         `json:"rpc_port"`
	ConversionLimit int            `json:"conversion_limit"`
	Database        DatabaseConfig `json:"database"`
}

func DefaultConfig() *Config {
	return &Config{
		RestPort:        "8084",
		RPCPort:         "8095",
		ConversionLimit: 1,
		Database: DatabaseConfig{
			Addr:     "localhost:3306",
			Username: "root",
			Password: "111111",
		},
	}
}
