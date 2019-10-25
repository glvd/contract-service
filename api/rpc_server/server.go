package rpc_server

import (
	"io"
	"log"
	"net"
	"net/rpc"

	"service/api"
)

type server struct {
	cfg       api.Config
	manager   *api.Manager
	rpcServer *rpc.Server
}

func (s *server) Start() {

	server := rpc.NewServer()
	server.Register("")

	ln, err := net.Listen("tcp", ":"+"port")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Print("rpc.Serve: accept:", err.Error())
			return
		}
		go serveConn(server, conn)
	}
}

func serveConn(server *rpc.Server, conn io.ReadWriteCloser) {
}
