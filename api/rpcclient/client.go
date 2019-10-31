package rpcclient

import (
	"errors"
	"strings"

	"service/api"
	"service/api/pb"

	"github.com/goextension/log"
	"google.golang.org/grpc"
)

// RPCClient ...
type RPCClient interface {
	api.Runnable
	api.Client
}

type rpcclient struct {
	cfg     api.Config
	conn    *grpc.ClientConn
	manager *api.Manager
}

// Options ...
type Options func(cli *rpcclient)

// Manager ...
func Manager(manager *api.Manager) Options {
	return func(c *rpcclient) {
		c.manager = manager
	}
}

// NewClient ...
func NewClient(cfg api.Config) RPCClient {
	return &rpcclient{
		cfg: cfg,
	}
}

// WrapError ...
func WrapError(reply *pb.WorkReply, err error) error {
	if err != nil {
		return err
	}
	if reply.Error != "" {
		return errors.New(reply.Error)
	}
	return nil
}

func rpcAddr(addr, port string) string {
	if addr == "" {
		addr = "127.0.0.1"
	}
	return strings.Join([]string{addr, port}, ":")
}
func client(conn *grpc.ClientConn) pb.ServiceClient {
	return pb.NewServiceClient(conn)
}

// Start ...
func (r *rpcclient) Start() (err error) {
	log.Info("rpc client was dail on:", r.cfg.RPCPort)
	r.conn, err = grpc.Dial(rpcAddr(r.cfg.RPCAddr, r.cfg.RPCPort), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Panicf("did not connect: %v", err)
	}

	// Contact the server and print out its response.
	//name := defaultName
	//if len(os.Args) > 1 {
	//	name = os.Args[1]
	//}
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//defer cancel()
	//r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	//if err != nil {
	//	log.Fatalf("could not greet: %v", err)
	//}
	//log.Printf("Greeting: %s", r.GetMessage())
	return nil
}

// Stop ...
func (r *rpcclient) Stop() {
	e := r.conn.Close()
	if e != nil {
		log.Warnw("closed conn")
	}
}

// DeleteWork ...
func (r *rpcclient) DeleteWork(manager *api.Manager, id string) error {
	reply, e := client(r.conn).Work(manager.Context(), &pb.WorkRequest{
		Msg:      pb.MessageType_Delete,
		WorkMode: pb.WorkMode_LocalMode,
		Work: &pb.Work{
			ID: id,
		},
	})

	return WrapError(reply, e)
}

// GetWork ...
func (r *rpcclient) GetWork(manager *api.Manager, id string) (*api.Work, error) {
	reply, err := client(r.conn).Work(manager.Context(), &pb.WorkRequest{
		Msg:      pb.MessageType_Status,
		WorkMode: pb.WorkMode_LocalMode,
		Work: &pb.Work{
			ID: id,
		},
	})
	if err != nil {
		return nil, err
	}
	work := api.RPCWorkToWork(reply.Works[0])
	return work, nil
}

// GetWorks ...
func (r *rpcclient) GetWorks(manager *api.Manager) ([]*api.Work, error) {
	log.Info("rpc client", "connect", "GetWorks")
	reply, err := client(r.conn).Work(manager.Context(), &pb.WorkRequest{
		Msg:      pb.MessageType_List,
		WorkMode: pb.WorkMode_LocalMode,
	})
	if err != nil {
		return nil, err
	}
	var works []*api.Work
	for _, w := range reply.Works {
		works = append(works, api.RPCWorkToWork(w))
	}
	return works, nil
}

// AddWork ...
func (r *rpcclient) AddWork(manager *api.Manager, work api.Work) error {
	log.Infow("rpc client", "func", "AddWork")
	reply, e := client(r.conn).Work(manager.Context(), &pb.WorkRequest{
		Msg:      pb.MessageType_Add,
		WorkMode: pb.WorkMode_LocalMode,
		Work:     api.WorkToRPCWork(&work),
	})
	return WrapError(reply, e)
}

// GetNode ...
func (r *rpcclient) GetNode(manager *api.Manager, id string) {
	panic("implement me")
}

// AddNode ...
func (r *rpcclient) AddNode(manager *api.Manager, node api.Node) {
	panic("implement me")
}

// PostNode ...
func (r rpcclient) PostNode(manager *api.Manager) {
	panic("implement me")
}

// DeleteNode ...
func (r rpcclient) DeleteNode(manager *api.Manager) {
	panic("implement me")
}

// GetVideos ...
func (r rpcclient) GetVideos(manager *api.Manager) {
	panic("implement me")
}
