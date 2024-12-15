package repository

import (
	"todo_backend/internal/model"

	"gorm.io/gorm"
)

type TaskRepository interface {
	GetAllTasks() ([]model.Task, error)
	GetTask() ([]model.Task, error)
	CreateTask(task *model.Task) error
	UpdateTask(task *model.Task) error
	Delete(task *model.Task) error
}

type TaskRepositoryImpl struct {
	DB *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepositoryImpl {
	return &TaskRepositoryImpl{DB: db}
}

func (r *TaskRepositoryImpl) GetAllTasks() ([]model.Task, error) {
	var tasks []model.Task
	err := r.DB.Find(&tasks).Error
	return tasks, err
}

func (r *TaskRepositoryImpl) GetTask(id int) (*model.Task, error) {
	var task model.Task
	err := r.DB.Model(&model.Task{}).Where("id = ?", id).First(&task).Error
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *TaskRepositoryImpl) CreateTask(task *model.Task) error {
	err := r.DB.Create(task).Error
	return err
}

func (r *TaskRepositoryImpl) UpdateTaskFields(id int, fields map[string]interface{}) error {
	err := r.DB.Model(&model.Task{}).Where("id = ?", id).Updates(fields).Error
	return err
}
func (r *TaskRepositoryImpl) DeleteTask(id int) error {
	err := r.DB.Model(&model.Task{}).Where("id = ?", id).Delete(&model.Task{}).Error
	return err
}
