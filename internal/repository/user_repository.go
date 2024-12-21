package repository

import (
	"todo_backend/internal/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *model.User) error
}

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{DB: db}
}

func (r *UserRepositoryImpl) CreateUser(user *model.User) error {
	err := r.DB.Create(user).Error
	return err
}

func (r *UserRepositoryImpl) FindByUserId(LoginId string) (model.User, error) {
	var user model.User
	err := r.DB.Where("login_id = ?", LoginId).First(&user).Error
	return user, err
}