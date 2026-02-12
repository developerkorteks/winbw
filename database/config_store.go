package database

// DBConfigStore implements the ConfigStore interface
type DBConfigStore struct {
	db interface {
		GetConfig(key string) (string, error)
		GetAllConfigs() (map[string]string, error)
		SetConfig(key, value, updatedBy string) error
	}
}

// NewConfigStore creates a new config store
func NewConfigStore() *DBConfigStore {
	return &DBConfigStore{
		db: &dbWrapper{},
	}
}

// dbWrapper wraps the package-level DB functions
type dbWrapper struct{}

func (w *dbWrapper) GetConfig(key string) (string, error) {
	return GetConfig(key)
}

func (w *dbWrapper) GetAllConfigs() (map[string]string, error) {
	return GetAllConfigs()
}

func (w *dbWrapper) SetConfig(key, value, updatedBy string) error {
	return SetConfig(key, value, updatedBy)
}

// Implement ConfigStore interface
func (s *DBConfigStore) GetConfig(key string) (string, error) {
	return s.db.GetConfig(key)
}

func (s *DBConfigStore) GetAllConfigs() (map[string]string, error) {
	return s.db.GetAllConfigs()
}

func (s *DBConfigStore) SetConfig(key, value, updatedBy string) error {
	return s.db.SetConfig(key, value, updatedBy)
}
