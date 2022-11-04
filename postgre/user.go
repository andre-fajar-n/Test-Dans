package postgre

import (
	"context"
	"dans/entity"

	"gorm.io/gorm"
)

type User struct {
	db *gorm.DB
}

func NewUserPostgre(db *gorm.DB) entity.UserPostgre {
	return &User{
		db,
	}
}

func (r *User) GetByUsername(ctx context.Context, username string) (*entity.User, error) {
	output := &entity.User{}

	if err := r.db.Model(&output).Where("username = ?", username).First(&output).Error; err != nil {
		return nil, err
	}

	return output, nil
}

func (r *User) Create(ctx context.Context, data entity.User) error {
	if err := r.db.Model(data).Create(&data).Error; err != nil {
		return err
	}

	return nil
}
