package db

import (
	"context"
	"fmt"

	"github.com/fair-n-square-co/transactions/internal/db/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User interface {
	CreateUser(ctx context.Context, user CreateUserFields) (uuid.UUID, error)
}

type CreateUserFields struct {
	Email string
	Username string
	FirstName string
	LastName string
}

type user struct {
	db *gorm.DB
}

func (u *user) CreateUser(ctx context.Context, user CreateUserFields) (uuid.UUID, error) {
	userModel := models.User {
		Email: user.Email,
		Username: user.Username,
		FirstName: user.FirstName,
		LastName: user.LastName,
	}

	result := u.db.Create(&userModel)

	if result.Error != nil {
		return uuid.Nil, fmt.Errorf("failed to create user: %v", result.Error)
	}
	return userModel.ID, nil
}

func newUser(db *gorm.DB) User {
	return &user{
		db: db,
	}
}