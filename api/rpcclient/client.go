package rpcclient

import (
	"errors"
	"strings"

	"service/api"
	"service/api/pb"

	"github.com/goextension/log"
	"google.golang.org/grpc"
)

type rpcclient struct {
	cfg  api.Config
	conn *grpc.ClientConn
}

// NewClient ...
func NewClient(cfg api.Config) api.Client {
	return &rpcclient{
		cfg: cfg,
	}
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
	// Set up a connection to the server.
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
		log.Warnw("close conn")
	}
}

// DeleteWork ...
func (r *rpcclient) DeleteWork(manager *api.Manager, id string) error {
	panic("implement me")
}

// GetWork ...
func (r *rpcclient) GetWork(manager *api.Manager, id string) (*api.Work, error) {
	reply, err := client(r.conn).Work(manager.Context(), &pb.WorkRequest{
		Msg:      pb.MessageType_Status,
		WorkMode: pb.WorkMode_LocalMode,
		ID:       id,
	})
	if err != nil {
		return nil, err
	}
	work := api.RPCWorkToWork(reply.Work, reply.Status)
	return work, nil
}

// GetWorks ...
func (r *rpcclient) GetWorks(manager *api.Manager) ([]*api.Work, error) {
	return nil, nil
}

// AddWork ...
func (r *rpcclient) AddWork(manager *api.Manager, work api.Work) error {
	reply, e := client(r.conn).Work(manager.Context(), &pb.WorkRequest{
		Msg:      pb.MessageType_Add,
		WorkMode: pb.WorkMode_LocalMode,
		Work:     api.WorkToRPCWork(&work),
	})
	if e != nil {
		return e
	}
	if reply.Error != "" {
		return errors.New(reply.Error)
	}
	return nil
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
