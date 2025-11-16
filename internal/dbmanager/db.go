package dbmanager

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sayban/internal/config"
)

func (m *Manager) Create(cfg *config.Config, name string, password string) error {
	err := os.MkdirAll(cfg.DBUrl, 0700)
	if err != nil {
		return fmt.Errorf("failed to init database: %w", err)
	}

	dbPath := filepath.Join(cfg.DBUrl, name+".db")

	_, err = os.Stat(dbPath)
	if err == nil {
		return fmt.Errorf("database is already exists for user %s", name)
	}
	db, err := sql.Open("sqlite3", fmt.Sprintf("%s?_key=%s", dbPath, password))
	log.Println("db was created")
	if err != nil {
		return fmt.Errorf("database error: %w", err)
	}
	defer db.Close()

	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS messages (id INTEGER PRIMARY KEY, content TEXT);"); err != nil {
		return fmt.Errorf("failed to init schema: %w", err)
	}
	log.Println("DB table was created")

	// identity keys table
	if _, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS identity_keys (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	identity_public_key TEXT NOT NULL,
	identity_private_key TEXT NOT NULL,
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	)`); err != nil {
		return fmt.Errorf("failed to init schema: %w", err)
	}
	log.Println("DB table was created")

	// signed pre-keys table
	if _, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS signed_prekeys (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	user_id TEXT,
	key_id INTEGER,
	public_key TEXT NOT NULL,
	private_key TEXT NOT NULL,
	signature TEXT NOT NULL,
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
	expires_at DATETIME,
	is_active INTEGER DEFAULT 1
	)`); err != nil {
		return fmt.Errorf("failed to init schema: %w", err)
	}

	// One-Time-pre-keys table
	if _, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS one_time_prekeys (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	key_id INTEGER,
	public_key TEXT NOT NULL,
	private_key TEXT NOT NULL,
	used_at DATETIME,
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	)`); err != nil {
		return fmt.Errorf("failed to init schema: %w", err)
	}

	// Account-Table
	if _, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS accounts (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name VARCHAR(100) UNIQUE,
	onion_address VARCHAR(64) UNIQUE NOT NULL,
	identity_key TEXT NOT NULL,
	
	)
	`); err != nil {
		return fmt.Errorf("faild to init schema: %w", err)
	}

	return nil
}
