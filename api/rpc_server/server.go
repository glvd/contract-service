package rpc_server

import "net/rpc"

type server struct {
	rpcServer *rpc.Server
}

func NewServer() *server {
	rpcServer := rpc.NewServer()
	return &server{rpcServer: rpcServer}
}
