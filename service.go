package service

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"

	"service/api"
	"service/api/restapi"
	"service/api/rpcserver"

	"github.com/glvd/conversion"
	"github.com/goextension/log"
	"github.com/goextension/tool"
)

// DefaultFolder ...
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
	ctx     context.Context
	cancel  context.CancelFunc
	done    chan bool
	secret  string
	config  *Config
	manager *api.Manager
	task    *conversion.Task
}

// NewService ...
func NewService() *Service {
	ctx, cancel := context.WithCancel(context.Background())
	cfg := DefaultConfig()

	e := cfg.LoadJSON()
	if e != nil {
		log.Panicw("can't load json file", "error", e)
	}
	secret := tool.GenerateRandomString(64)
	fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!SECRET!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
	fmt.Println(secret)
	fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!SECRET!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
	e = ioutil.WriteFile(filepath.Join(DefaultPath, "secret"), []byte(secret), 0600)
	if e != nil {
		log.Panicw("can't save secret file", "error", e)
	}
	task, err := initConversion(cfg.Conversion)
	if err != nil {
		log.Panicw("init conversion failed", "error", err)
	}

	manager := api.NewManager(ctx)

	rpcServer := rpcserver.NewRPCServer(cfg.API)
	manager.SetServer(rpcServer)

	rest := restapi.NewRestAPI(cfg.API, restapi.Manager(manager))
	manager.RegisterClient(api.RestAPI, rest)
	local := &core{}
	manager.RegisterClient(api.LocalClient, local)

	e = cfg.SaveJSON()
	if e != nil {
		log.Panicw("can't save json file", "error", e)
	}
	return &Service{
		ctx:     ctx,
		cancel:  cancel,
		done:    make(chan bool),
		secret:  secret,
		config:  cfg,
		manager: manager,
		task:    task,
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

	go func() {

	}()

	s.manager.StartAll()

	return nil
}

// Wait ...
func (s *Service) Wait() {
	<-s.done
}

// Stop ...
func (s *Service) Stop() {
	s.task.Stop()
}
