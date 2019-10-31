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
	Tmp         string `json:"tmp"`
	Cache       string `json:"cache"`
	Limit       int    `json:"limit"`
	Mode        string `json:"mode"`
	Addr        string `json:"addr"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	SingleAddr  string `json:"node_addr"`
	ClusterAddr string `json:"cluster_addr"`
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
			Timeout:  30,
			Remote:   "127.0.0.1",
			RestPort: "8084",
			RPCAddr:  "127.0.0.1",
			RPCPort:  "8085",
		},
		Conversion: ConversionConfig{
			Tmp:         filepath.Join(DefaultPath, "tmp"),
			Cache:       "cache",
			Limit:       1,
			Mode:        NodeModeCluster,
			Addr:        "localhost:3306",
			Username:    "root",
			Password:    "111111",
			SingleAddr:  "/ip4/127.0.0.1/tcp/5001",
			ClusterAddr: "/ip4/127.0.0.1/tcp/9094",
		},
	}
}
