package service

import (
	"final_project_3/dto"
	"final_project_3/entity"
	"final_project_3/pkg/errs"
	"final_project_3/pkg/helpers"
	"final_project_3/repository/category_repository"
)

type categoryService struct {
	CategoryRepo category_repository.Repository
}

type CategoryService interface {
	CreateCategory(categoryPayload *dto.NewCategoryRequest) (*dto.NewCategoryResponse, errs.Error)
	GetCategoriesWithTasks() (*[]dto.GetCategoriesResponse, errs.Error)
}

func NewCategoryService(categoryRepo category_repository.Repository) CategoryService {
	return &categoryService{CategoryRepo: categoryRepo}
}

func (cs *categoryService) CreateCategory(categoryPayload *dto.NewCategoryRequest) (*dto.NewCategoryResponse, errs.Error) {
	validateErr := helpers.ValidateStruct(categoryPayload)
	if validateErr != nil {
		return nil, validateErr
	}
	category := entity.Category{
		Type: categoryPayload.Type,
	}
	createdCategory, err := cs.CategoryRepo.CreateCategory(&category)
	if err != nil {
		return nil, err
	}
	response := dto.NewCategoryResponse{
		ID:        createdCategory.ID,
		Type:      createdCategory.Type,
		CreatedAt: createdCategory.CreatedAt,
	}
	return &response, nil
}

func (cs *categoryService) GetCategoriesWithTasks() (*[]dto.GetCategoriesResponse, errs.Error) {
	response, err := cs.CategoryRepo.GetCategoriesWithTasks()
	if err != nil {
		return nil, err
	}
	return response, nil
}
