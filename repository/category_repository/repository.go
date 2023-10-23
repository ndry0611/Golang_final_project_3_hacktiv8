package category_repository

import (
	"final_project_3/entity"
	"final_project_3/pkg/errs"
)

type Repository interface {
	CreateCategory (categoryPayload *entity.Category) (*entity.Category, errs.Error)
}