package api

import (
	"context"
	"sync"
)

// Runnable ...
type Runnable interface {
	Start() error
	Stop()
}

// Manager ...
type Manager struct {
	server  Server
	clients *sync.Map
	ctx     context.Context
	cancel  context.CancelFunc
}

// NewManager ...
func NewManager(ctx context.Context) *Manager {
	m := &Manager{
		clients: &sync.Map{},
	}
	m.ctx, m.cancel = context.WithCancel(ctx)
	return m
}

// RegisterClient ...
func (m *Manager) RegisterClient(name string, client Client) {
	m.clients.Store(name, client)
}

// SetServer ...
func (m *Manager) SetServer(server Server) {
	m.server = server
}

// Client ...
func (m *Manager) Client(name string) (client Client) {
	value, ok := m.clients.Load(name)
	if !ok {
		return dummyClient{}
	}
	return value.(Client)
}

// Context ...
func (m *Manager) Context() context.Context {
	return m.ctx
}
