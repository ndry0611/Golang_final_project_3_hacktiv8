package service

import (
	"final_project_3/dto"
	"final_project_3/entity"
	"final_project_3/pkg/errs"
	"final_project_3/pkg/helpers"
	user_repository "final_project_3/repository/user_repository"
)

type userService struct {
	UserRepo user_repository.Repository
}

type UserService interface {
	CreateUser(userPayload *dto.NewUserRequest) (*dto.NewUserResponse, errs.Error)
}

func NewUserService(userRepo user_repository.Repository) UserService {
	return &userService{UserRepo: userRepo}
}

func (us *userService) CreateUser(userPayload *dto.NewUserRequest) (*dto.NewUserResponse, errs.Error) {
	validateErr := helpers.ValidateStruct(userPayload)
	if validateErr != nil {
		return nil, validateErr
	}
	user := entity.User{
		FullName: userPayload.FullName,
		Email:    userPayload.Email,
		Password: userPayload.Password,
	}
	createdUser, err := us.UserRepo.CreateUser(&user)
	if err != nil {
		return nil, err
	}
	response := dto.NewUserResponse{
		ID:        createdUser.ID,
		FullName:  createdUser.FullName,
		Email:     createdUser.Email,
		CreatedAt: createdUser.CreatedAt,
	}
	return &response, nil
}
