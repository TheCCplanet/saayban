package service

import (
	"sayban/internal/config"
)

type UserServiceInterFace interface {
	Register(*config.Config, string, string) error
	Unlock(*config.Config, string, string) error
	Lock(string) error

	// GetDB(string, string) (*dbmanager.Manager, error)
}
