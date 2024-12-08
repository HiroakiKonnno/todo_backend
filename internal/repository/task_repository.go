package repository

import (
	"todo_backend/internal/model"

	"gorm.io/gorm"
)

type TaskRepository interface {
	GetAllTasks() ([]model.Task, error)
	CreateTask(task *model.Task) error
}

type TaskRepositoryImpl struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepositoryImpl {
	return &TaskRepositoryImpl{db: db}
}

func (r *TaskRepositoryImpl) GetAllTasks() ([]model.Task, error) {
	var tasks []model.Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *TaskRepositoryImpl) CreateTask(task *model.Task) error {
	err := r.db.Create(task).Error
	return err
}

