package manage

import (
	"sync"

	"github.com/gofiber/contrib/websocket"
)

type Manager struct {
	Connections map[string]*websocket.Conn
	mutex       *sync.Mutex
}

func NewManager() *Manager {
	return &Manager{
		Connections: map[string]*websocket.Conn{},
		mutex:       &sync.Mutex{},
	}
}

// Adds websocket connection to connection manager
func (m *Manager) Add(username string, c *websocket.Conn) {
	m.mutex.Lock()
	m.Connections[username] = c
	m.mutex.Unlock()
}

// Removes websocket connection from the connection manager
func (m *Manager) Remove(username string) {
	m.mutex.Lock()
	delete(m.Connections, username)
	m.mutex.Unlock()
}

// Retrieve user websocket connection
func (m *Manager) GetConnection(username string) *websocket.Conn {
	return m.Connections[username]
}

// Get list of connected user's usernames
func (m *Manager) GetConnectedUsers() []string {
	var users []string
	for k := range m.Connections {
		users = append(users, k)
	}

	return users
}
