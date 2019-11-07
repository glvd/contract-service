package service

import (
	"encoding/json"
	"strings"

	"service/api"

	"github.com/glvd/conversion"
	"github.com/goextension/log"
)

type serviceHandle struct {
	cfg     Config
	task    *conversion.Task
	cluster interface{}
}

// AddNodes ...
func (s *serviceHandle) AddNode(manager *api.Manager, node api.Node) {
	panic("implement me")
}

// AddWork ...
func (s *serviceHandle) AddWork(manager *api.Manager, work api.Work) error {
	reader := strings.NewReader(work.VideoInfo)
	dec := json.NewDecoder(reader)
	info := conversion.VideoPornInfo{}
	if err := dec.Decode(&info); err != nil {
		log.Errorw("decode info error", "source", work.VideoInfo, "error", err)
		return err
	}

	iwork, e := conversion.NewInfoWork(&info,
		conversion.VideoPathOption(work.VideoPath),
		conversion.PosterPathOption(work.PosterPath),
		conversion.ThumbPathOption(work.ThumbPath),
		conversion.SamplePathOption(work.SamplePath),
		conversion.SkipOption("source"),
	)
	if e != nil {
		return e
	}
	e = s.task.AddWorker(iwork, true)
	if e != nil {
		return e
	}
	return nil
}

// DeleteNode ...
func (s *serviceHandle) DeleteNode(manager *api.Manager) {
	panic("implement me")
}

// DeleteWork ...
func (s *serviceHandle) DeleteWork(manager *api.Manager, id string) error {
	s.task.StopWork(id)
	return nil
}

// GetNode ...
func (s *serviceHandle) GetNode(manager *api.Manager, id string) {
	panic("implement me")
}

// GetVideos ...
func (s serviceHandle) GetVideos(manager *api.Manager) {
	panic("implement me")
}

// GetWork ...
func (s *serviceHandle) GetWork(manager *api.Manager, id string) (*api.Work, error) {
	iwork, err := s.task.GetWork(id)
	if err != nil {
		return nil, err
	}
	return work(iwork), nil
}

// GetWorks ...
func (s *serviceHandle) GetWorks(manager *api.Manager) (works []*api.Work, e error) {
	iworks, err := s.task.AllRun()
	if err != nil {
		return nil, err
	}
	log.Infow("get all works", "works", iworks)
	for _, w := range iworks {
		works = append(works, work(w))
	}
	return
}

// Start ...
func (s *serviceHandle) Start() error {
	if err := s.task.Start(); err != nil {
		log.Panicw(err.Error(), "func", "service|start")
	}
	return nil
}

// Stop ...
func (s *serviceHandle) Stop() {
	panic("implement me")
}

func work(iwork conversion.IWork) *api.Work {
	w := iwork.Work()
	return &api.Work{
		ID:         iwork.ID(),
		WorkStatus: int32(iwork.Status()),
		VideoPath:  w.VideoPaths,
		PosterPath: w.PosterPath,
		ThumbPath:  w.ThumbPath,
		SamplePath: w.SamplePath,
		VideoInfo:  iwork.Info(),
	}
}
