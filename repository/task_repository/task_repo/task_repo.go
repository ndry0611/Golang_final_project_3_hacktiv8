package task_repo

import (
	"final_project_3/entity"
	"final_project_3/pkg/errs"
	"final_project_3/repository/task_repository"

	"gorm.io/gorm"
)

type taskRepo struct {
	db *gorm.DB
}

func NewTaskRepo(db *gorm.DB) task_repository.Repository {
	return &taskRepo{db: db}
}

func (tr *taskRepo) CreateTask(taskPayload *entity.Task) (*entity.Task, errs.Error) {
	var Task = *taskPayload
	err := tr.db.Create(&Task).Error
	if err != nil {
		return nil, errs.NewInternalServerError(err.Error())
	}
	return &Task, nil
}