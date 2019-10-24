package rpc_server

import "net/rpc"

type server struct {
	rpcServ *rpc.Server
}

func NewServer() *server {
	rpcServ := rpc.NewServer()
	return &server{rpcServ: rpcServ}
}
