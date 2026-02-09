package repository

import (
	"github.com/google/uuid"
	"github.com/luongtien872003/login/internal/model"
)

// 1. Init UserRepository interface
type UserRepository interface {
	Create(user *model.User) error
	GetByEmail(email string) (*model.User, error)
	GetByID(id uuid.UUID) (*model.User, error)
	UpdateStatus(id uuid.UUID, status string) error
}
