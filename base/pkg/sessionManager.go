package pkg

import "sync"

type SessionManager struct {
	sessions map[string]string
	mutex    sync.Mutex
}

func NewSessionManager() *SessionManager {
	return &SessionManager{
		sessions: make(map[string]string),
	}
}

func (m *SessionManager) SetSession(userID, sessionID string) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.sessions[userID] = sessionID
}

func (m *SessionManager) GetSession(userID string) (string, bool) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	sessionID, ok := m.sessions[userID]
	return sessionID, ok
}

func (m *SessionManager) DeleteSession(userID string) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	delete(m.sessions, userID)
}
