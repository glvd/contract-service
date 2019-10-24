package service

import "service/api"

type ConversionConfig struct {
	Limit    int    `json:"limit"`
	Addr     string `json:"addr"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// Config ...
type Config struct {
	API        api.Config       `json:"api"`
	Conversion ConversionConfig `json:"conversion"`
}

func DefaultConfig() *Config {
	return &Config{
		API: api.Config{
			RestPort: "8084",
			RPCPort:  "8095",
		},
		Conversion: ConversionConfig{
			Limit:    1,
			Addr:     "localhost:3306",
			Username: "root",
			Password: "111111",
		},
	}
}
