package handler

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/luongtien872003/login/internal/model"
	"github.com/luongtien872003/login/internal/pkg/password"
	"github.com/luongtien872003/login/internal/repository"
)

// 1. Init AuthHandler
type AuthHandler struct {
	UserRepo repository.UserRepository
}

// 2. NewAuthHandler creates a new AuthHandler
func NewAuthHandler(userRepo repository.UserRepository) *AuthHandler {
	return &AuthHandler{UserRepo: userRepo}
}

// 3. RegisterRequest struct
type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

// 4. Register function
func (h *AuthHandler) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	// Check email đã tồn tại
	_, err := h.UserRepo.GetByEmail(req.Email)
	if err == nil {
		c.JSON(409, gin.H{"error": "Email already exists"})
		return
	}
	// Hash password
	log.Printf("password raw = '%s', len=%d\n", req.Password, len(req.Password))

	hashedPassword, err := password.Hash(req.Password)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to hash password"})
		return
	}
	// Tạo user với status = pending
	user := &model.User{
		Email:        req.Email,
		PasswordHash: hashedPassword,
		Status:       model.StatusPending,
	}

	if err := h.UserRepo.Create(user); err != nil {
		log.Printf("Failed to create user: %v", err)
		c.JSON(500, gin.H{"error": "Failed to create user"})
		return
	}

	// Return success
	c.JSON(201, gin.H{"message": "User created successfully"})
}
