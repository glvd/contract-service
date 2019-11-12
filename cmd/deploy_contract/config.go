package main

import (
	"encoding/json"
	"io/ioutil"
)

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

var _config Config

// DefaultConfig ...
func DefaultConfig() Config {
	return Config{
		DBConfig: DBConfig{
			DBType:   "sqlite3",
			Schema:   "deploy",
			Username: "",
			Password: "",
			Address:  ".",
		},
	}

}

// LoadConfig ...
func LoadConfig(path string) error {
	bys, e := ioutil.ReadFile(path)
	if e != nil {
		_config = DefaultConfig()
		return e
	}
	var cfg Config

	e = json.Unmarshal(bys, &cfg)
	if e != nil {
		_config = DefaultConfig()
		return e
	}
	_config = cfg
	return nil
}
