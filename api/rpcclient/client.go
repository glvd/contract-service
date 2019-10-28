package rpcclient

import (
	"errors"
	"log"
	"strings"

	"service/api"
	"service/api/pb"

	"google.golang.org/grpc"
)

type rpcclient struct {
	cfg       api.Config
	rpcClient pb.ServiceClient
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

// Start ...
func (r *rpcclient) Start() error {
	// Set up a connection to the server.
	conn, err := grpc.Dial(rpcAddr(r.cfg.RPCAddr, r.cfg.RPCPort), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	r.rpcClient = pb.NewServiceClient(conn)

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
}

// DeleteWork ...
func (r *rpcclient) DeleteWork(manager *api.Manager, id string) error {
	panic("implement me")
}

// GetWork ...
func (r *rpcclient) GetWork(manager *api.Manager, id string) error {
	panic("implement me")
}

// GetWorks ...
func (r *rpcclient) GetWorks(manager *api.Manager) ([]*api.Work, error) {
	panic("implement me")
}

// AddWork ...
func (r *rpcclient) AddWork(manager *api.Manager, work api.Work) error {
	reply, e := r.rpcClient.Work(manager.Context(), &pb.WorkRequest{
		Msg:        pb.MessageType_Add,
		WorkMode:   pb.WorkMode_LocalMode,
		VideoPath:  work.VideoPath,
		PosterPath: work.PosterPath,
		ThumbPath:  work.ThumbPath,
		SamplePath: work.SamplePath,
		VideoInfo:  work.VideoInfo,
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
