package repository

import (
	"github.com/kalyanbasso/taskhub/internal/model"
	"gorm.io/gorm"
)

const (
	orderByCreatedAtDesc = "created_at desc"
)

type TaskRepository interface {
	Create(task *model.Task) error
	GetByID(id uint) (*model.Task, error)
	GetAll() ([]*model.Task, error)
	Update(task *model.Task) error
	Delete(id uint) error
	GetByCompleted(completed bool) ([]*model.Task, error)
	GetByPriority(priority model.Priority) ([]*model.Task, error)
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{
		db: db,
	}
}

func (r *taskRepository) Create(task *model.Task) error {
	return r.db.Create(task).Error
}

func (r *taskRepository) GetByID(id uint) (*model.Task, error) {
	var task model.Task
	err := r.db.First(&task, id).Error
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *taskRepository) GetAll() ([]*model.Task, error) {
	var tasks []*model.Task
	err := r.db.Order(orderByCreatedAtDesc).Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) Update(task *model.Task) error {
	return r.db.Save(task).Error
}

func (r *taskRepository) Delete(id uint) error {
	return r.db.Delete(&model.Task{}, id).Error
}

func (r *taskRepository) GetByCompleted(completed bool) ([]*model.Task, error) {
	var tasks []*model.Task
	err := r.db.Where("completed = ?", completed).Order(orderByCreatedAtDesc).Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) GetByPriority(priority model.Priority) ([]*model.Task, error) {
	var tasks []*model.Task
	err := r.db.Where("priority = ?", priority).Order(orderByCreatedAtDesc).Find(&tasks).Error
	return tasks, err
}
