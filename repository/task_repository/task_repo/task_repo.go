package task_repo

import (
	"errors"
	"final_project_3/dto"
	"final_project_3/entity"
	"final_project_3/pkg/errs"
	"final_project_3/repository/task_repository"
	"strconv"

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

func (tr *taskRepo) GetTasks() (*[]dto.GetTasksResponse, errs.Error) {
	var tasks []entity.Task
	err := tr.db.Preload("User").Find(&tasks).Error
	if err != nil {
		return nil, errs.NewInternalServerError(err.Error())
	}
	var res []dto.GetTasksResponse
	for _, task := range tasks {
		res = append(res, dto.GetTasksResponse{
			ID:          task.ID,
			Title:       task.Title,
			Status:      task.Status,
			Description: task.Description,
			UserID:      task.UserID,
			CategoryID:  task.CategoryID,
			User: dto.GetTaskUser{
				ID:       task.User.ID,
				Email:    task.User.Email,
				FullName: task.User.FullName,
			},
		})
	}
	return &res, nil
}

func (tr *taskRepo) FindTaskById(id int) (*entity.Task, errs.Error) {
	var Task entity.Task
	err := tr.db.Where("id = ?", id).Preload("User").First(&Task).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			msg := "Task with id: " + strconv.Itoa(id) + " not found"
			return nil, errs.NewNotFoundError(msg)
		}
	}
	return &Task, nil
}

func (tr *taskRepo) UpdateTask(taskPayload *entity.Task) (*entity.Task, errs.Error) {
	var Task = *taskPayload
	err := tr.db.Model(&Task).Updates(map[string]interface{}{"title": Task.Title, "description": Task.Description}).First(&Task).Error
	if err != nil {
		return nil, errs.NewInternalServerError(err.Error())
	}
	return &Task, nil
}

func (tr *taskRepo) UpdateTaskStatus(taskPayload *entity.Task) (*entity.Task, errs.Error) {
	var Task = *taskPayload
	err := tr.db.Model(&Task).Update("status", Task.Status).First(&Task).Error
	if err != nil {
		return nil, errs.NewInternalServerError(err.Error())
	}
	return &Task, nil
}

func (tr *taskRepo) UpdateTaskCategory(taskPayload *entity.Task) (*entity.Task, errs.Error) {
	var Task = *taskPayload
	err := tr.db.Model(&Task).Update("category_id", Task.CategoryID).First(&Task).Error
	if err != nil {
		return nil, errs.NewInternalServerError(err.Error())
	}
	return &Task, nil
}

func (tr *taskRepo) DeleteTask(id int) errs.Error {
	err := tr.db.Where("id = ?", id).Delete(entity.Task{}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			msg := "task with id: " + strconv.Itoa(id) + " not found"
			return errs.NewNotFoundError(msg)
		}
		return errs.NewInternalServerError("something went wrong")
	}
	return nil
}