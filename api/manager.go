package api

import (
	"sync"

	"github.com/goextension/log"
)

type Manager struct {
	mutex   *sync.Mutex
	clients map[string]Client
}

func NewManager() *Manager {
	return &Manager{
		mutex:   &sync.Mutex{},
		clients: make(map[string]Client),
	}
}

func (m *Manager) Register(name string, client Client) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.clients[name] = client
}

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

func (m *Manager) StartAll() {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	for name, cli := range m.clients {
		go func(name string, client Client) {
			e := cli.Start()
			if e != nil {
				log.Panicw("can't start client", "name", name, "error", e)
			}
		}(name, cli)
	}
}
