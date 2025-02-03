package repository

import (
	"todo_backend/internal/model"

	"time"

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

func (r *TaskRepositoryImpl) GetAllTasks(start_date *time.Time, end_date *time.Time) ([]model.Task, error) {
	query := r.DB
	
	if start_date != nil {
		query = query.Where("start_date >= ?", *start_date)
	}
	if end_date != nil {
		query = query.Where("end_date <= ?", *end_date)
	}

	var tasks []model.Task
	err := query.Find(&tasks).Error
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
