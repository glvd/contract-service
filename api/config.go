package api

// Config ...
type Config struct {
	Timeout  int    `json:"timeout"`
	Remote   string `json:"remote"`
	RestPort string `json:"rest_port"`
	RPCAddr  string `json:"rpc_addr"`
	RPCPort  string `json:"rpc_port"`
}
