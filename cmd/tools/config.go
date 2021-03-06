package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/glvd/contract-service/contract"
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
}

// InitConfig ...
type InitConfig struct {
	Database DBConfig
	Contract ContractConfig
	Gateway  string
	KeyPass  string
	KeyPath  string
}

// DefaultConfigName ...
var DefaultConfigName = "config.json"

// DefaultKeyPass ...
var DefaultKeyPass = "123"

// DefaultKeyPath ...
var DefaultKeyPath = "945d35cd4a6549213e8d37feb5d708ec98906902"

// ConfigPath ...
var ConfigPath = ""

var _config Config
var _init InitConfig

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

func initDefault() {
	contract.DefaultMessageAddress = compareAssign(contract.DefaultMessageAddress, _init.Contract.DMessage, _config.Contract.DMessage)
	contract.DefaultTagAddress = compareAssign(contract.DefaultTagAddress, _init.Contract.DTag, _config.Contract.DTag)
	contract.DefaultNodeAddress = compareAssign(contract.DefaultNodeAddress, _init.Contract.DNode, _config.Contract.DNode)
	contract.DefaultGatway = compareAssign(contract.DefaultGatway, _init.Gateway, _config.Gateway)
	DefaultKeyPass = compareAssign(DefaultKeyPass, _init.KeyPass)
	DefaultKeyPath = compareAssign(DefaultKeyPath, _init.KeyPath)

	_config.Gateway = contract.DefaultGatway
	_config.Contract.DMessage = contract.DefaultMessageAddress
	_config.Contract.DTag = contract.DefaultTagAddress
	_config.Contract.DNode = contract.DefaultNodeAddress
}

func defaultAssign(source string, v ...string) string {
	for _, vv := range v {
		if source != vv {
			return vv
		}
	}
	return source
}

func compareAssign(source string, v ...string) string {
	for _, vv := range v {
		if source != vv && vv != "" {
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
