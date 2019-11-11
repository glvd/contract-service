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

// DefaultConfig ...
var DefaultConfig = "deploy.json"

// DefaultConfigPath ...
var DefaultConfigPath = "config"

// LoadConfig ...
func LoadConfig() {

}
