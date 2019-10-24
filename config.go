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
