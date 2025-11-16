package dbmanager

import (
	"fmt"
	"sayban/internal/models"
)

func (m *Manager) AddAccount(acc models.Account) error {
	m.Mu.Lock()
	defer m.Mu.Unlock()

	_, err := m.DB.Exec(`
	INSERT INTO accounts (name, onion_address, identity_key) VALUES($1,$2,$3)
	`, acc.Name, acc.OnionAddress, acc.IdentityKey)
	if err != nil {
		return fmt.Errorf("faild to add account : %w", err)
	}

	return nil
}

func (m *Manager) DeleteAccountByAddress(accountAddress string) error {
	m.Mu.Lock()
	defer m.Mu.Unlock()

	_, err := m.DB.Exec(`
	DELETE FROM accounts WHERE address=$1
	`, accountAddress)
	if err != nil {
		return fmt.Errorf("faild to delete account by id: %w", err)
	}
	return nil
}

func (m *Manager) GetAccountList() ([]*models.Account, error) {
	m.Mu.Lock()
	defer m.Mu.Unlock()

	rows, err := m.DB.Query(`
	SELECT id, name, onion_address, identity_key FROM accounts
	`)
	if err != nil {
		return nil, fmt.Errorf("faild to get accounts list : %w", err)
	}
	defer rows.Close()

	var accounts []*models.Account
	for rows.Next() {
		var acc models.Account

		rows.Scan(&acc.ID, &acc.Name, &acc.OnionAddress, &acc.IdentityKey)
		accounts = append(accounts, &acc)
	}

	return accounts, nil
}
