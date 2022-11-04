package entity

import (
	"context"

	"gorm.io/gorm"
)

// Base Model
type (
	User struct {
		gorm.Model
		Username string `gorm:"column:username;uniqueIndex:uniq_key_username;type:varchar(20);not null:true" json:"username"`
		Password string `gorm:"column:password;type:varchar(255);not null:true" json:"password"`
	}
)

// Interface
type (
	UserPostgre interface {
		GetByUsername(ctx context.Context, username string) (*User, error)
		Create(ctx context.Context, data User) error
	}

	UserUsecase interface {
		Register(ctx context.Context, req *UserRegisterRequest) error
	}
)

// Request
type (
	UserRegisterRequest struct {
		Username string `validate:"required" json:"username"`
		Password string `validate:"required" json:"password"`
	}
)
