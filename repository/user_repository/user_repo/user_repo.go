package user_repo

import (
	"final_project_3/entity"
	"final_project_3/pkg/errs"
	"final_project_3/repository/user_repository"

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
