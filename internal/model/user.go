package model

import (
	"time"

	"github.com/google/uuid"
)

// 1. Create Type User
type User struct {
	ID           uuid.UUID `json:"id" db:"id"`
	Email        string    `json:"email" db:"email"`
	PasswordHash string    `json:"-" db:"password_hash"` // Không bao giờ return
	Status       string    `json:"status" db:"status"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

// 2. Create Status
const (
	StatusPending = "pending"
	StatusActive  = "active"
	StatusBanned  = "banned"
)
