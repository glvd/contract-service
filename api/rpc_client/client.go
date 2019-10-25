package rpc_client

import (
	"net/rpc"

	"service/api"
)

type rpcclient struct {
	rpcClient *rpc.Client
}

func NewClient() api.Client {
	rpcClient := rpc.NewClient(nil)
	return &rpcclient{
		rpcClient: rpcClient,
	}
}

func (r rpcclient) Start() error {
	panic("start")
}

func (r rpcclient) Stop() {
	panic("implement me")
}

func (r rpcclient) GetTasks(manager *api.Manager) ([]*api.Task, error) {
	panic("implement me")
}

func (r rpcclient) PostTask(manager *api.Manager) error {
	panic("implement me")
}

func (r rpcclient) GetTask(manager *api.Manager, id string) error {
	panic("implement me")
}

func (r rpcclient) DeleteTask(manager *api.Manager, id string) error {
	panic("implement me")
}

func (r rpcclient) GetNode(manager *api.Manager) {
	panic("implement me")
}

func (r rpcclient) PostNode(manager *api.Manager) {
	panic("implement me")
}

func (r rpcclient) DeleteNode(manager *api.Manager) {
	panic("implement me")
}

func (r rpcclient) GetVideos(manager *api.Manager) {
	panic("implement me")
}
