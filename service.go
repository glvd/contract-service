package service

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"

	"github.com/glvd/conversion"
	"github.com/goextension/log"
	"github.com/goextension/tool"
)

var (
	DefaultFolder = ".service"
	DefaultPath   string
)

func init() {

	// We try guessing user's home from the HOME variable. This
	// allows HOME hacks for things like Snapcraft builds. HOME
	// should be set in all UNIX by the OS. Alternatively, we fall back to
	// usr.HomeDir (which should work on Windows etc.).
	dir, e := os.Getwd()
	if e != nil {
		dir := os.Getenv("HOME")
		if dir == "" {
			usr, err := user.Current()
			if err != nil {
				panic(fmt.Sprintf("cannot get current user: %s", err))
			}
			dir = usr.HomeDir
		}
	}
	DefaultPath = filepath.Join(dir, DefaultFolder)
}

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
	log.Warnw("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!SECRET!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!", "secret", secret)
	e = ioutil.WriteFile(filepath.Join(DefaultPath, "secret"), []byte(secret), 0600)
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
