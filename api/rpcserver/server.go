package rpcserver

import (
	"context"
	"fmt"
	"net"

	"service/api"
	"service/api/pb"

	"github.com/goextension/log"
	"google.golang.org/grpc"
)

type server struct {
	cfg       api.Config
	manager   *api.Manager
	rpcServer *grpc.Server
}

// Options ...
type Options func(sever *server)

func init() {

}

var workHandles map[string]interface{}

// Work ...
func (s *server) Work(ctx context.Context, req *pb.WorkRequest) (*pb.WorkReply, error) {
	cli := s.manager.Client(api.LocalClient)
	work := api.RPCWorkToWork(req.Work, 0)
	switch req.Msg {
	case pb.MessageType_Add:
		if err := cli.AddWork(s.manager, *work); err != nil {
			return nil, err
		}
	case pb.MessageType_Status:
		getWork, err := cli.GetWork(s.manager, req.ID)
		if err != nil {
			return nil, err
		}
		work := api.WorkToRPCWork(getWork)
		return &pb.WorkReply{
			Msg:    req.Msg,
			Status: pb.WorkStatus(getWork.WorkStatus),
			Work:   work,
		}, nil
	}

	return &pb.WorkReply{}, nil
}

// Node ...
func (s *server) Node(context.Context, *pb.NodeRequest) (*pb.NodeReply, error) {
	return &pb.NodeReply{}, nil
}

// NewRPCServer ...
func NewRPCServer(cfg api.Config, options ...Options) api.Server {
	s := &server{
		cfg: cfg,
	}
	for _, option := range options {
		option(s)
	}
	return s
}

// Start ...
func (s *server) Start() error {
	lis, err := net.Listen("tcp", ":"+s.cfg.RPCPort)
	if err != nil {
		return fmt.Errorf("failed to listen:%w", err)
	}
	s.rpcServer = grpc.NewServer()
	pb.RegisterServiceServer(s.rpcServer, &server{})
	if err := s.rpcServer.Serve(lis); err != nil {
		log.Panicw("failed to server", "error", err)
	}
	return nil
}

// Stop ...
func (s *server) Stop() {
	if s.rpcServer != nil {
		s.rpcServer.Stop()
	}
}
