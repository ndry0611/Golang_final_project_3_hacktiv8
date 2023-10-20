package user_repository

import (
	"final_project_3/entity"
	"final_project_3/pkg/errs"
)

type Repository interface {
	CreateUser(userPayload *entity.User) (*entity.User, errs.Error)
}
