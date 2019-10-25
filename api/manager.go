package api

import "sync"

type Manager struct {
	sync.Mutex
	clients map[string]Client
}

var manager = initManager()

func initManager() *Manager {
	return &Manager{
		clients: make(map[string]Client),
	}
}

func (m *Manager) Register(name string, client Client) {
	m.Lock()
	defer m.Unlock()
	m.clients[name] = client
}

func (m *Manager) Client(name string) (client Client, b bool) {
	m.Lock()
	defer m.Unlock()
	client, b = m.clients[name]
	return
}
