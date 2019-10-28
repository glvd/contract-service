package rpc_server

import (
	"context"
	"fmt"
	"net"

	"service/api"
	pb "service/api/pb"

	"github.com/goextension/log"
	"google.golang.org/grpc"
)

type server struct {
	cfg       api.Config
	manager   *api.Manager
	rpcServer *grpc.Server
}

func (s *server) Work(context.Context, *pb.WorkRequest) (*pb.WorkReply, error) {
	panic("implement me")
}

func (s *server) Node(context.Context, *pb.NodeRequest) (*pb.NodeReply, error) {
	panic("implement me")
}

type Options func(sever *server)

func init() {

}
func NewRPCServer(cfg api.Config, options ...Options) api.Server {
	s := &server{
		cfg: cfg,
	}
	for _, option := range options {
		option(s)
	}
	return s
}

func (s *server) Start() error {
	lis, err := net.Listen("tcp", ":"+s.cfg.RPCPort)
	if err != nil {
		return fmt.Errorf("failed to listen:%w", err)
	}
	s.rpcServer = grpc.NewServer()
	pb.RegisterServiceServer(s.rpcServer, &server{})
	go func() {
		if err := s.rpcServer.Serve(lis); err != nil {
			log.Panicw("failed to server", "error", err)
		}
	}()
	return nil
}
func (s *server) Stop() {
	if s.rpcServer != nil {
		s.rpcServer.Stop()
	}
}
