package user_repo

import (
	"final_project_3/entity"
	"final_project_3/pkg/errs"
	"final_project_3/repository/user_repository"
	"errors"

	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) user_repository.Repository {
	return &userRepo{db: db}
}

func (ur *userRepo) CreateUser(userPayload *entity.User) (*entity.User, errs.Error) {
	var User = *userPayload
	err := ur.db.Create(&User).Error
	if err != nil {
		return nil, errs.NewInternalServerError(err.Error())
	}
	return &User, nil
}

func (ur *userRepo) FindOneUserByEmail(email string) (*entity.User, errs.Error) {
	var User entity.User
	err := ur.db.Where("email = ?", email).First(&User).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			msg := "user with email: " + email + " not found"
			return nil, errs.NewNotFoundError(msg)
		}
	}
	return &User, nil
}