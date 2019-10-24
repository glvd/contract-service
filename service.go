package service

import (
	"github.com/glvd/conversion"
	"github.com/goextension/log"
	"github.com/goextension/tool"
)

// Service ...
type Service struct {
	config Config
	secret string
	task   *conversion.Task
}

// NewService ...
func NewService() *Service {

	secret := tool.GenerateRandomString(32)
	task := initConversion()
	return &Service{
		secret: secret,
		task:   task,
	}
}

// Start ...
func (s *Service) Start() error {
	//start conversion task
	go func() {
		if err := s.task.Start(); err != nil {
			log.Panicw(err.Error(), "func", "service|start")
			return
		}
	}()
	return nil
}

// Stop ...
func (s *Service) Stop() {
	s.task.Stop()
}
