package category_repo

import (
	"final_project_3/entity"
	"final_project_3/pkg/errs"
	"final_project_3/repository/category_repository"

	"gorm.io/gorm"
)

type categoryRepo struct {
	db *gorm.DB
}

func NewCategoryRepo(db *gorm.DB) category_repository.Repository {
	return &categoryRepo{db: db}
}

func (cr *categoryRepo) CreateCategory(categoryPayload *entity.Category) (*entity.Category, errs.Error) {
	var Category = *categoryPayload
	err := cr.db.Create(&Category).Error
	if err != nil {
		return nil, errs.NewInternalServerError(err.Error())
	}
	return &Category, nil
}