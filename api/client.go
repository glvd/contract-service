package api

import (
	"errors"

	"github.com/goextension/log"
)

const RPCCLient = "rpcclient"
const RPCServer = "rpcserver"
const RestAPI = "restapi"

type Client interface {
	Runnable
	GetWorks(manager *Manager) ([]*Work, error)
	AddWork(manager *Manager) error
	GetWork(manager *Manager, id string) error
	DeleteWork(manager *Manager, id string) error
	GetNode(manager *Manager)
	PostNode(manager *Manager)
	DeleteNode(manager *Manager)
	GetVideos(manager *Manager)
}
type Work struct {
	VideoPath  []string
	PosterPath string
	ThumbPath  string
	SamplePath string
	VideoInfo  string
}

type dummyClient struct {
}

func (d dummyClient) AddWork(manager *Manager) error {
	panic("implement me")
}

func (d dummyClient) GetWork(manager *Manager, id string) error {
	panic("implement me")
}

func (d dummyClient) DeleteWork(manager *Manager, id string) error {
	panic("implement me")
}

func (d dummyClient) GetNode(manager *Manager) {
	panic("implement me")
}

func (d dummyClient) PostNode(manager *Manager) {
	panic("implement me")
}

func (d dummyClient) DeleteNode(manager *Manager) {
	panic("implement me")
}

func (d dummyClient) GetVideos(manager *Manager) {
	panic("implement me")
}

func (d dummyClient) Start() error {
	log.Error("dummy start")
	return nil
}

func (d dummyClient) Stop() {
	log.Error("dummy stop")
}

func (d dummyClient) GetWorks(manager *Manager) ([]*Work, error) {
	log.Error("dummy", "func", "GetWorks")
	return nil, errors.New("wrong client to call")
}
