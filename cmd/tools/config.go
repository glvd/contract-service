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
	Database DBConfig
	Contract ContractConfig
	Gateway  string
	KeyPass  string
	KeyPath  string
	DNode    string `json:"-"`
	DTag     string `json:"-"`
	DMessage string `json:"-"`
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
	contract.DefaultMessageAddress = compareAssign(contract.DefaultMessageAddress, _config.DMessage, _config.Contract.DMessage)
	contract.DefaultTagAddress = compareAssign(contract.DefaultTagAddress, _config.DTag, _config.Contract.DTag)
	contract.DefaultNodeAddress = compareAssign(contract.DefaultNodeAddress, _config.DNode, _config.Contract.DNode)
}

func compareAssign(source string, v ...string) string {
	for _, vv := range v {
		if source != vv {
			return vv
		}
	}
	return source
}

// LoadConfig ...
func LoadConfig(path string) (e error) {
	defer func() {
		if e != nil {
			_config.Database = DefaultDBConfig()
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
