package service

import (
	"io/ioutil"
	"path/filepath"

	"github.com/glvd/conversion"
	"github.com/goextension/log"
	"github.com/goextension/tool"
)

var Path = ".service"

// Service ...
type Service struct {
	secret string
	config *Config
	task   *conversion.Task
}

// NewService ...
func NewService() *Service {
	cfg := DefaultConfig()

	e := cfg.LoadJSON()
	if e != nil {
		log.Panicw("can't load json file", "error", e)
	}
	secret := tool.GenerateRandomString(32)
	e = ioutil.WriteFile(filepath.Join(Path, "secret"), []byte(secret), 0600)
	if e != nil {
		log.Panicw("can't save secret file", "error", e)
	}
	task, err := initConversion(cfg.Conversion)
	if err != nil {
		log.Error(err)
		panic(err)
	}

	cfg.SaveJSON()

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
