package api

import (
	"context"
	"sync"

	"github.com/goextension/log"
)

// Runnable ...
type Runnable interface {
	Start() error
	Stop()
}

// Manager ...
type Manager struct {
	mutex   *sync.Mutex
	server  Server
	clients map[string]Client
	ctx     context.Context
	cancel  context.CancelFunc
}

// NewManager ...
func NewManager(ctx context.Context) *Manager {
	m := &Manager{
		mutex:   &sync.Mutex{},
		clients: make(map[string]Client),
	}
	m.ctx, m.cancel = context.WithCancel(ctx)
	return m
}

// RegisterClient ...
func (m *Manager) RegisterClient(name string, client Client) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.clients[name] = client
}

// SetServer ...
func (m *Manager) SetServer(server Server) {
	m.server = server
}

// Client ...
func (m *Manager) Client(name string) (client Client) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	var b bool
	client, b = m.clients[name]
	if !b {
		return dummyClient{}
	}
	return
}

// Context ...
func (m *Manager) Context() context.Context {
	return m.ctx
}

// StartAll ...
func (m *Manager) StartAll() {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	go func() {
		e := m.server.Start()
		if e != nil {
			log.Panicw("can't start rpc server", "name", "rpc server", "error", e)
		}
	}()

	for name, cli := range m.clients {
		go func(name string, client Client) {
			e := cli.Start()
			if e != nil {
				log.Panicw("can't start client", "name", name, "error", e)
			}
		}(name, cli)
	}
}
