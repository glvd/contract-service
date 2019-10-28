package api

import (
	"errors"

	"github.com/goextension/log"
)

// RPCClient ...
const RPCClient = "rpcclient"

// RPCServer ...
const RPCServer = "rpcserver"

// LocalClient ...
const LocalClient = "localclient"

// RestAPI ...
const RestAPI = "restapi"

// Client ...
type Client interface {
	Runnable
	GetWorks(manager *Manager) ([]*Work, error)
	AddWork(manager *Manager, work Work) error
	GetWork(manager *Manager, id string) error
	DeleteWork(manager *Manager, id string) error
	GetNode(manager *Manager, id string)
	AddNode(manager *Manager, node Node)
	DeleteNode(manager *Manager)
	GetVideos(manager *Manager)
}

// Work ...
type Work struct {
	VideoPath  []string
	PosterPath string
	ThumbPath  string
	SamplePath []string
	VideoInfo  string
}

// Node ...
type Node struct {
}

type dummyClient struct {
}

// AddWork ...
func (d dummyClient) AddWork(manager *Manager, work Work) error {
	panic("implement me")
}

// GetWork ...
func (d dummyClient) GetWork(manager *Manager, id string) error {
	panic("implement me")
}

// DeleteWork ...
func (d dummyClient) DeleteWork(manager *Manager, id string) error {
	panic("implement me")
}

// GetNode ...
func (d dummyClient) GetNode(manager *Manager, id string) {
	panic("implement me")
}

// AddNode ...
func (d dummyClient) AddNode(manager *Manager, node Node) {
	panic("implement me")
}

// DeleteNode ...
func (d dummyClient) DeleteNode(manager *Manager) {
	panic("implement me")
}

// GetVideos ...
func (d dummyClient) GetVideos(manager *Manager) {
	panic("implement me")
}

// Start ...
func (d dummyClient) Start() error {
	log.Error("dummy start")
	return nil
}

// Stop ...
func (d dummyClient) Stop() {
	log.Error("dummy stop")
}

// GetWorks ...
func (d dummyClient) GetWorks(manager *Manager) ([]*Work, error) {
	log.Error("dummy", "func", "GetWorks")
	return nil, errors.New("wrong client to call")
}
