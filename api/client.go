package api

import (
	"errors"

	"github.com/goextension/log"
)

const RPCCLient = "rpcclient"
const RPCServer = "rpcserver"
const RestAPI = "restapi"

type Client interface {
	Start() error
	Stop()
	GetTasks(manager *Manager) ([]*Task, error)
	PostTask(manager *Manager) error
	GetTask(manager *Manager, id string) error
	DeleteTask(manager *Manager, id string) error
	GetNode(manager *Manager)
	PostNode(manager *Manager)
	DeleteNode(manager *Manager)
	GetVideos(manager *Manager)
}
type Task struct {
}

type dummyClient struct {
}

func (d dummyClient) Start() error {
	log.Error("dummy start")
	return nil
}

func (d dummyClient) Stop() {
	log.Error("dummy stop")
}

func (d dummyClient) GetTasks(manager *Manager) ([]*Task, error) {
	log.Error("dummy", "func", "GetTasks")
	return nil, errors.New("wrong client to call")
}
