package api

import (
	"sync"

	"github.com/goextension/log"
)

type Manager struct {
	sync.Mutex
	clients map[string]Client
}

func NewManager() *Manager {
	return &Manager{
		clients: make(map[string]Client),
	}
}

func (m *Manager) Register(name string, client Client) {
	m.Lock()
	defer m.Unlock()
	m.clients[name] = client
}

func (m *Manager) Client(name string) (client Client) {
	m.Lock()
	defer m.Unlock()
	var b bool
	client, b = m.clients[name]
	if !b {
		return dummyClient{}
	}
	return
}

func (m *Manager) StartAll() {
	m.Lock()
	defer m.Unlock()
	for name, cli := range m.clients {
		go func(name string, client Client) {
			e := cli.Start()
			if e != nil {
				log.Panicw("can't start client", "name", name, "error", e)
			}
		}(name, cli)
	}
}
