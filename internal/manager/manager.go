package manage

import "github.com/gofiber/contrib/websocket"

type Manager struct {
	Connections map[string]*websocket.Conn
}

// Adds websocket connection to connection manager
func (m *Manager) Add(username string, c *websocket.Conn) {
	m.Connections[username] = c
}

// Removes websocket connection from the connection manager
func (m *Manager) Remove(username string) {
	delete(m.Connections, username)
}
