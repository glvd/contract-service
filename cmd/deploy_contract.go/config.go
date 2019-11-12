package main

// DBConfig ...
type DBConfig struct {
	DBType   string
	Schema   string
	Username string
	Password string
	Address  string
}

// Config ...
type Config struct {
	DBConfig
}

// DefaultConfigName ...
var DefaultConfigName = "deploy.json"

// DefaultConfigPath ...
var DefaultConfigPath = "config"

// ConfigPath ...
var ConfigPath = ""

// DefaultConfig ...
func DefaultConfig() Config {
	return Config{
		DBConfig: DBConfig{
			DBType:   "sqlite3",
			Schema:   "deploy.db",
			Username: "",
			Password: "",
			Address:  ".",
		},
	}

}

// LoadConfig ...
func LoadConfig(path string) {

}
