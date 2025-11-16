package service

import (
	"fmt"
	"sayban/internal/config"
	"sayban/internal/dbmanager"
	"sayban/internal/models"

	_ "github.com/mattn/go-sqlite3"
)

type UserService struct {
	users map[string]*dbmanager.Manager
}

func NewUserService() *UserService {
	return &UserService{
		users: make(map[string]*dbmanager.Manager),
	}
}

func (s *UserService) Register(cfg *config.Config, userName string, password string) error {

	_, exists := s.users[userName]
	if exists {
		return fmt.Errorf("user is already registered")
	}
	m := dbmanager.NewManger(cfg)
	s.users[userName] = m
	err := m.Create(cfg, userName, password)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) Lock(userName string) error {
	m, exists := s.users[userName]
	if !exists {
		return fmt.Errorf("there's no db userNamed %s", userName)
	}

	return m.Lock()
}
func (s *UserService) Unlock(cfg *config.Config, userName string, password string) error {
	m, exists := s.users[userName]
	if !exists {
		return fmt.Errorf("there's no db userNamed %s", userName)
	}
	err := m.Unlock(cfg, userName, password)
	if err != nil {
		return err
	}

	return nil
}

// func (s *UserService) SendKeyBundle() error {

// 	return nil
// }

func (s *UserService) DeleteAccountByID(userName string, onionAddress string) error {
	m, exists := s.users[userName]
	if !exists {
		return fmt.Errorf("their's no db userNamed %s", userName)
	}
	err := m.DeleteAccountByAddress(onionAddress)
	return err
}

func (s *UserService) RegisterAccount(userName string, account models.Account) error {
	m, exists := s.users[userName]
	if !exists {
		return fmt.Errorf("their's no db userNamed %s", userName)
	}
	err := m.AddAccount(account)

	return err
}

func (s *UserService) GetAccountList(userName string) ([]*models.Account, error) {
	m, exists := s.users[userName]
	if !exists {
		return nil, fmt.Errorf("their's no db userNamed %s", userName)
	}
	accounts, err := m.GetAccountList()

	return accounts, err
}
