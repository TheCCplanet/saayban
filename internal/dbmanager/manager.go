package dbmanager

import (
	"database/sql"
	"errors"
	"sayban/internal/config"
	"sync"
	"time"
)

type Manager struct {
	DB           *sql.DB
	DBPath       string
	IsLocked     bool
	Timer        *time.Timer
	AutoLockTime time.Duration
	Mu           sync.Mutex
}
type DBPool struct {
	Mu  sync.Mutex
	DBs map[string]*Manager
}

func NewDBPool() *DBPool {
	return &DBPool{
		Mu:  sync.Mutex{},
		DBs: make(map[string]*Manager),
	}
}

func NewManger(cfg *config.Config) *Manager {
	return &Manager{
		DBPath:       cfg.DBUrl,
		AutoLockTime: cfg.AutoLockTime,
		IsLocked:     true,
	}
}

func (m *Manager) GetDB() (*sql.DB, error) {
	m.Mu.Lock()
	defer m.Mu.Unlock()

	if m.IsLocked {
		return nil, errors.New("database is locked")
	}

	return m.DB, nil
}

func (m *Manager) Close() {
	m.Mu.Lock()
	defer m.Mu.Unlock()

	if m.DB != nil {
		m.DB.Close()
		m.DB = nil
	}
	m.IsLocked = true
}
