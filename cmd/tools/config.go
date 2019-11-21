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
	Gateway  string
	KeyPass  string
	KeyPath  string
	DNode    string
	DTag     string
	DMessage string
}

// DefaultConfigName ...
var DefaultConfigName = "deploy.json"

// DefaultConfigPath ...
var DefaultConfigPath = "config"

// ConfigPath ...
var ConfigPath = ""

var _config Config

// DefaultDBConfig ...
func DefaultDBConfig() DBConfig {
	return DBConfig{
		DBType:   "sqlite3",
		Schema:   "deploy",
		Username: "",
		Password: "",
		Address:  ".",
	}

}

// LoadConfig ...
func LoadConfig(path string) (e error) {
	defer func() {
		if e != nil {
			_config.DBConfig = DefaultDBConfig()
		}
	}()
	bys, e := ioutil.ReadFile(path)
	if e != nil {
		return e
	}

	e = json.Unmarshal(bys, &_config.DBConfig)
	if e != nil {
		return e
	}
	return nil
}
