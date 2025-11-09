package repository

// import (
// 	"database/sql"
// 	"sayban/internal/models"
// 	"strings"
// 	"time"
// )

// type userRepositoryImpl struct {
// 	db *sql.DB
// }

// func NewUserRepository(db *sql.DB) UserRepository {
// 	return &userRepositoryImpl{db: db}
// }

// func (r *userRepositoryImpl) CreateUser(name, passwordHash string) (*models.User, error) {
// 	query := `INSERT INTO users (name, password_hash, created_at, updated_at) VALUES (?, ?, ?, ?)`

// 	now := time.Now()
// 	result, err := r.db.Exec(query, name, passwordHash, now, now)
// 	if err != nil {
// 		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
// 			return nil, models.ErrUserExists
// 		}
// 		return nil, err
// 	}

// 	id, err := result.LastInsertId()
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &models.User{
// 		ID:           int(id),
// 		Name:         name,
// 		PasswordHash: passwordHash,
// 		CreatedAt:    now,
// 		UpdatedAt:    now,
// 	}, nil
// }

// func (r *userRepositoryImpl) GetUserByName(name string) (*models.User, error) {
// 	var user models.User
// 	query := `SELECT id, name, password_hash, created_at, updated_at FROM users WHERE name = ?`

// 	err := r.db.QueryRow(query, name).Scan(
// 		&user.ID, &user.Name, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt,
// 	)

// 	if err == sql.ErrNoRows {
// 		return nil, models.ErrUserNotFound
// 	}
// 	return &user, err
// }

// func (r *userRepositoryImpl) GetUserByID(id int) (*models.User, error) {
// 	var user models.User
// 	query := `SELECT id, name, password_hash, created_at, updated_at FROM users WHERE id = ?`

// 	err := r.db.QueryRow(query, id).Scan(
// 		&user.ID, &user.Name, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt,
// 	)

// 	if err == sql.ErrNoRows {
// 		return nil, models.ErrUserNotFound
// 	}
// 	return &user, err
// }
