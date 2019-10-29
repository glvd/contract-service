package api

import (
	"errors"

	"service/api/pb"

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
	//Runnable
	GetWorks(manager *Manager) ([]*Work, error)
	AddWork(manager *Manager, work Work) error
	GetWork(manager *Manager, id string) (*Work, error)
	DeleteWork(manager *Manager, id string) error
	GetNode(manager *Manager, id string)
	AddNode(manager *Manager, node Node)
	DeleteNode(manager *Manager)
	GetVideos(manager *Manager)
}

// Work ...
type Work struct {
	ID         string   `json:"id"`
	WorkStatus int32    `json:"work_status"`
	VideoPath  []string `json:"video_path"`
	PosterPath string   `json:"poster_path"`
	ThumbPath  string   `json:"thumb_path"`
	SamplePath []string `json:"sample_path"`
	VideoInfo  string   `json:"video_info"`
}

// Node ...
type Node struct {
}

type dummyClient struct {
}

var _ Client = &dummyClient{}

// RPCWorkToWork ...
func RPCWorkToWork(work *pb.Work) *Work {
	if work == nil {
		return &Work{}
	}
	return &Work{
		WorkStatus: (int32)(work.Status),
		ID:         work.GetID(),
		VideoPath:  work.GetVideoPath(),
		PosterPath: work.GetPosterPath(),
		ThumbPath:  work.GetThumbPath(),
		SamplePath: work.GetSamplePath(),
		VideoInfo:  work.GetVideoInfo(),
	}
}

// WorkToRPCWork ...
func WorkToRPCWork(work *Work) *pb.Work {
	if work == nil {
		return &pb.Work{}
	}
	return &pb.Work{
		Status:     pb.WorkStatus(work.WorkStatus),
		ID:         work.ID,
		VideoPath:  work.VideoPath,
		PosterPath: work.PosterPath,
		ThumbPath:  work.ThumbPath,
		SamplePath: work.SamplePath,
		VideoInfo:  work.VideoInfo,
	}
}

// AddWork ...
func (d dummyClient) AddWork(manager *Manager, work Work) error {
	panic("implement me")
}

// DeleteWork ...
func (d dummyClient) DeleteWork(manager *Manager, id string) error {
	panic("implement me")
}

// GetWork ...
func (d dummyClient) GetWork(manager *Manager, id string) (*Work, error) {
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
