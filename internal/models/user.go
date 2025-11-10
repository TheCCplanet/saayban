package models

import (
	"errors"
	"strings"
)

var (
	ErrInvalidName        = errors.New("invalid name")
	ErrInvalidPassword    = errors.New("password must be at least 6 characters")
	ErrUserExists         = errors.New("user already exists")
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidCredentials = errors.New("invalid credentials")
)

// type User struct {
// 	ID           int       `json:"id"`
// 	Name         string    `json:"name"`
// 	PasswordHash string    `json:"-"`
// 	CreatedAt    time.Time `json:"created_at"`
// 	UpdatedAt    time.Time `json:"updated_at"`
// }

type ErrorResponse struct {
	Error string `json:"error"`
}

type RegisterRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (r *RegisterRequest) Validate() error {
	r.Name = strings.TrimSpace(r.Name)
	if r.Name == "" {
		return ErrInvalidName
	}
	if len(r.Password) < 6 {
		return ErrInvalidPassword
	}
	return nil
}

type UnlockRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (u *UnlockRequest) Validate() error {
	u.Name = strings.TrimSpace(u.Name)
	if u.Name == "" {
		return ErrInvalidName
	}
	if len(u.Password) < 6 {
		return ErrInvalidPassword
	}
	return nil
}

type LockRequest struct {
	Name string `json:"name"`
}

func (u *LockRequest) Validate() error {
	u.Name = strings.TrimSpace(u.Name)
	if u.Name == "" {
		return ErrInvalidName
	}

	return nil
}

// type Account struct {
// 	ID          string
// 	Onion       string
// 	IdentityKey string
// 	CreatedAt   time.Time
// 	UpdateAt    time.Time
// }
