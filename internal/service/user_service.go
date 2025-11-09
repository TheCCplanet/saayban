package service

import (
	"errors"
	"fmt"
	"sayban/internal/config"
	"sayban/internal/dbmanager"

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

func (s *UserService) Register(cfg *config.Config, name string, password string) error {

	_, exists := s.users[name]
	if exists {
		return fmt.Errorf("user is already registered")
	}
	m := dbmanager.NewManger(cfg)
	s.users[name] = m
	err := m.Create(cfg, name, password)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) Lock(name string) error {
	m, exists := s.users[name]
	if !exists {
		return errors.New("nothing found")
	}

	return m.Lock()
}
func (s *UserService) Unlock(cfg *config.Config, name string, passwrod string) error {
	m, exists := s.users[name]
	if !exists {
		return fmt.Errorf("their's no db named %s", name)
	}
	err := m.Unlock(cfg, name, passwrod)
	if err != nil {
		return err
	}

	return nil
}

// func (s *UserService) GetDB(name string) (*dbmanager.Manager, error) {
// 	m, exists := s.users[name]
// 	if !exists {
// 		return nil, errors.New("nothing found")
// 	}
// 	return m, nil
// }
