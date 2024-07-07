package manager

type Manager struct {
	storage *Storage
}

func NewManager(filename string) *Manager {
	return &Manager{
		storage: NewStorage(filename),
	}
}

func (m *Manager) ListNames() []string {
	return m.storage.ListNames()
}

func (m *Manager) SavePassword(name, password string) {
	m.storage.SavePassword(name, password)
}

func (m *Manager) GetPassword(name string) (string, bool) {
	return m.storage.GetPassword(name)
}
