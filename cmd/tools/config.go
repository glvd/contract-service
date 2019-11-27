package main

import (
	"encoding/json"
	"io/ioutil"

	"service/contract"
)

// DBConfig ...
type DBConfig struct {
	DBType   string
	Schema   string
	Username string
	Password string
	Address  string
}

// ContractConfig ...
type ContractConfig struct {
	DNode    string
	DTag     string
	DMessage string
}

// Config ...
type Config struct {
	DBConfig
	ContractConfig
	Gateway  string
	KeyPass  string
	KeyPath  string
	DNode    string
	DTag     string
	DMessage string
}

// DefaultConfigName ...
var DefaultConfigName = "config.json"

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

func initConfig() {
	if _config.DMessage != contract.DefaultMessageAddress {
		contract.DefaultMessageAddress = _config.DMessage
	}
	if _config.DTag != contract.DefaultTagAddress {
		contract.DefaultTagAddress = _config.DTag
	}
	if _config.DNode != contract.DefaultNodeAddress {
		contract.DefaultNodeAddress = _config.DNode
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
	var cfg Config
	e = json.Unmarshal(bys, &cfg)
	if e != nil {
		return e
	}
	_config = cfg
	return nil
}

// SaveConfig ...
func SaveConfig(path string) (e error) {
	bys, e := json.MarshalIndent(_config, "", " ")
	if e != nil {
		return e
	}
	return ioutil.WriteFile(path, bys, 0755)
}
