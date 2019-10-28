package service

import (
	"encoding/json"
	"strings"

	"service/api"

	"github.com/glvd/conversion"
)

type core struct {
	cfg     Config
	task    *conversion.Task
	cluster interface{}
}

// AddNode ...
func (c core) AddNode(manager *api.Manager, node api.Node) {
	panic("implement me")
}

// AddWork ...
func (c core) AddWork(manager *api.Manager, work api.Work) error {
	reader := strings.NewReader(work.VideoInfo)
	dec := json.NewDecoder(reader)
	info := conversion.VideoPornInfo{}
	if err := dec.Decode(&info); err != nil {
		return err
	}

	iWork, e := conversion.NewInfoWork(&info,
		conversion.VideoPathOption(work.VideoPath),
		conversion.PosterPathOption(work.PosterPath),
		conversion.ThumbPathOption(work.ThumbPath),
		conversion.SamplePathOption(work.SamplePath),
	)
	if e != nil {
		return e
	}
	e = c.task.AddWorker(iWork)
	if e != nil {
		return e
	}
	return nil
}

// DeleteNode ...
func (c core) DeleteNode(manager *api.Manager) {
	panic("implement me")
}

// DeleteWork ...
func (c core) DeleteWork(manager *api.Manager, id string) error {
	panic("implement me")
}

// GetNode ...
func (c core) GetNode(manager *api.Manager, id string) {
	panic("implement me")
}

// GetVideos ...
func (c core) GetVideos(manager *api.Manager) {
	panic("implement me")
}

// GetWork ...
func (c core) GetWork(manager *api.Manager, id string) error {
	panic("implement me")
}

// GetWorks ...
func (c core) GetWorks(manager *api.Manager) ([]*api.Work, error) {
	panic("implement me")
}

// Start ...
func (c core) Start() error {
	panic("implement me")
}

// Stop ...
func (c core) Stop() {
	panic("implement me")
}
