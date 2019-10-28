package service

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"service/api"
)

// ConfigName ...
const ConfigName = "service.json"

// ConversionConfig ...
type ConversionConfig struct {
	Cache    string `json:"cache"`
	Limit    int    `json:"limit"`
	Addr     string `json:"addr"`
	Username string `json:"username"`
	Password string `json:"password"`
	NodeAddr string `json:"node_addr"`
}

// Config ...
type Config struct {
	API        api.Config       `json:"api"`
	Conversion ConversionConfig `json:"conversion"`
}

func jsonPath() string {
	return filepath.Join(DefaultPath, ConfigName)
}

// LoadJSON ...
func (config *Config) LoadJSON() error {
	_, e := os.Stat(jsonPath())
	//skip when config not exist
	if e != nil && os.IsNotExist(e) {
		return nil
	} else if e != nil {
		return e
	}
	file, e := os.Open(jsonPath())
	if e != nil {
		return e
	}
	dec := json.NewDecoder(file)
	e = dec.Decode(config)
	if e != nil {
		return e
	}
	return nil
}

// SaveJSON ...
func (config *Config) SaveJSON() error {
	bytes, e := json.MarshalIndent(config, "", "  ")
	if e != nil {
		return e
	}
	e = ioutil.WriteFile(jsonPath(), bytes, 0600)
	if e != nil {
		return e
	}
	return nil
}

// DefaultConfig ...
func DefaultConfig() *Config {
	return &Config{
		API: api.Config{
			RestPort: "8084",
			RPCPort:  "8085",
		},
		Conversion: ConversionConfig{
			Cache:    "cache",
			Limit:    1,
			Addr:     "localhost:3306",
			Username: "root",
			Password: "111111",
			NodeAddr: "/ip4/127.0.0.1/tcp/5001",
		},
	}
}
