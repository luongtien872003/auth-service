package repository

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/luongtien872003/login/internal/model"
)

// PostgresUserRepository implements UserRepository interface
type PostgresUserRepository struct {
	db *sql.DB
}

// 1. NewUserRepository creates a new PostgresUserRepository
func NewUserRepository(db *sql.DB) UserRepository {
	return &PostgresUserRepository{db: db}
}

// 2. Create user
func (r *PostgresUserRepository) Create(user *model.User) error {
	query := `INSERT INTO users (id, email, password_hash, status, created_at, updated_at) 
			  VALUES ($1, $2, $3, $4, NOW(), NOW())`
	_, err := r.db.Exec(query, uuid.New(), user.Email, user.PasswordHash, user.Status)
	return err
}

// 3. Get user by email
func (r *PostgresUserRepository) GetByEmail(email string) (*model.User, error) {
	query := `SELECT id, email, password_hash, status, created_at, updated_at FROM users WHERE email = $1`
	user := &model.User{}
	err := r.db.QueryRow(query, email).Scan(
		&user.ID, &user.Email, &user.PasswordHash, &user.Status, &user.CreatedAt, &user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// 4. Get user by ID
func (r *PostgresUserRepository) GetByID(id uuid.UUID) (*model.User, error) {
	query := `SELECT id, email, password_hash, status, created_at, updated_at FROM users WHERE id = $1`
	user := &model.User{}
	err := r.db.QueryRow(query, id).Scan(
		&user.ID, &user.Email, &user.PasswordHash, &user.Status, &user.CreatedAt, &user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// 5. Update user status
func (r *PostgresUserRepository) UpdateStatus(id uuid.UUID, status string) error {
	query := `UPDATE users SET status = $1, updated_at = NOW() WHERE id = $2`
	_, err := r.db.Exec(query, status, id)
	return err
}
