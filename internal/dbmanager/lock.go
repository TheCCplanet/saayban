package dbmanager

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sayban/internal/config"
	"time"
)

func (m *Manager) Create(cfg *config.Config, name string, password string) error {
	err := os.MkdirAll(cfg.DBUrl, 0700)
	if err != nil {
		return fmt.Errorf("faild to init database: %e", err)
	}

	dbPath := filepath.Join(cfg.DBUrl, name+".db")

	_, err = os.Stat(dbPath)
	if err == nil {
		return fmt.Errorf("database is already exists for user %s", name)
	}
	db, err := sql.Open("sqlite3", fmt.Sprintf("%s?_key=%s", dbPath, password))
	if err != nil {
		return fmt.Errorf("database error: %w", err)
	}
	defer db.Close()

	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS messages (id INTEGER PRIMARY KEY, content TEXT);"); err != nil {
		return fmt.Errorf("failed to init schema: %w", err)
	}

	return nil
}

func (m *Manager) StartAutoLockTimer() {
	if m.Timer != nil {
		m.Timer.Stop()
	}
	m.Timer = time.AfterFunc(m.AutoLockTime, func() {
		_ = m.Lock() // Ignore error for auto-lock timer
	})
}

func (m *Manager) Unlock(cfg *config.Config, name string, password string) error {
	if !m.IsLocked {
		return fmt.Errorf("database is alredy unlocked")
	}

	m.Mu.Lock()
	defer m.Mu.Unlock()

	dbPath := filepath.Join(cfg.DBUrl, name+".db")

	_, err := os.Stat(dbPath)
	if err != nil {
		return fmt.Errorf("database is not exists for user %s", name)
	}

	db, err := sql.Open("sqlite3", fmt.Sprintf("%s?_key=%s", dbPath, password))
	if err != err {
		return err
	}
	err = db.Ping()
	if err != nil {
		return fmt.Errorf("unable to open db for %s", name)
	}

	m.DB = db
	m.IsLocked = false

	return nil
}

func (m *Manager) Lock() error {
	m.Mu.Lock()
	defer m.Mu.Unlock()

	if m.IsLocked {
		return fmt.Errorf("database is already locked")
	}

	if m.Timer != nil {
		m.Timer.Stop()
		m.Timer = nil
	}

	if m.DB != nil {
		if err := m.DB.Close(); err != nil {
			log.Printf("Error closing database: %v", err)
			// Continue anyway - database is effectively locked
		}
	}
	m.DB = nil
	m.IsLocked = true
	log.Println("Database is locked")
	return nil
}
