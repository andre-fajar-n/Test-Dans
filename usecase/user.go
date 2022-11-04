package usecase

import (
	"context"
	"dans/entity"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	userPostgre entity.UserPostgre
}

func NewUser(userPostgre entity.UserPostgre) entity.UserUsecase {
	return &User{
		userPostgre,
	}
}

func (u *User) Register(ctx context.Context, req *entity.UserRegisterRequest) error {
	user, err := u.userPostgre.GetByUsername(ctx, req.Username)
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	if user != nil {
		return errors.New("username already registered")
	}
	fmt.Println("REGISTER")

	// hash password
	hashed, err := HashAndSalt(req.Password)
	if err != nil {
		// ser.Logger.Error().
		// 	Msgf("failed to encrypt password with error: %s", err)
		return err
	}

	if err := u.userPostgre.Create(ctx, entity.User{
		Username: req.Username,
		Password: hashed,
	}); err != nil {
		return err
	}

	return nil
}

func HashAndSalt(password string) (string, error) {
	// hash and salt the password
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
