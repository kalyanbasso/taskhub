package usecase

import (
	"errors"
	"time"

	"github.com/kalyanbasso/taskhub/internal/model"
	"github.com/kalyanbasso/taskhub/internal/repository"
)

type TaskUseCase interface {
	CreateTask(req *model.CreateTaskRequest) (*model.TaskResponse, error)
	GetTaskByID(id uint) (*model.TaskResponse, error)
	GetAllTasks() ([]*model.TaskResponse, error)
	UpdateTask(id uint, req *model.UpdateTaskRequest) (*model.TaskResponse, error)
	DeleteTask(id uint) error
	CompleteTask(id uint) (*model.TaskResponse, error)
	GetTasksByStatus(completed bool) ([]*model.TaskResponse, error)
	GetTasksByPriority(priority model.Priority) ([]*model.TaskResponse, error)
	GetOverdueTasks() ([]*model.TaskResponse, error)
}

type taskUseCase struct {
	taskRepo repository.TaskRepository
}

func NewTaskUseCase(taskRepo repository.TaskRepository) TaskUseCase {
	return &taskUseCase{
		taskRepo: taskRepo,
	}
}

func (uc *taskUseCase) CreateTask(req *model.CreateTaskRequest) (*model.TaskResponse, error) {
	if req.Title == "" {
		return nil, errors.New("title is required")
	}

	// Validar prioridade
	if req.Priority != "" {
		if req.Priority != model.PriorityLow && req.Priority != model.PriorityMedium && req.Priority != model.PriorityHigh {
			return nil, ErrInvalidPriority
		}
	} else {
		req.Priority = model.PriorityMedium // default
	}

	// Validar due_date
	if req.DueDate != nil && req.DueDate.Before(time.Now()) {
		return nil, errors.New("due date cannot be in the past")
	}

	task := &model.Task{
		Title:       req.Title,
		Description: req.Description,
		Priority:    req.Priority,
		DueDate:     req.DueDate,
		Completed:   false,
	}

	if err := uc.taskRepo.Create(task); err != nil {
		return nil, err
	}

	return task.ToResponse(), nil
}

func (uc *taskUseCase) GetTaskByID(id uint) (*model.TaskResponse, error) {
	if id == 0 {
		return nil, ErrInvalidTaskID
	}

	task, err := uc.taskRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return task.ToResponse(), nil
}

func (uc *taskUseCase) GetAllTasks() ([]*model.TaskResponse, error) {
	tasks, err := uc.taskRepo.GetAll()
	if err != nil {
		return nil, err
	}

	responses := make([]*model.TaskResponse, len(tasks))
	for i, task := range tasks {
		responses[i] = task.ToResponse()
	}

	return responses, nil
}

func (uc *taskUseCase) UpdateTask(id uint, req *model.UpdateTaskRequest) (*model.TaskResponse, error) {
	if id == 0 {
		return nil, ErrInvalidTaskID
	}

	task, err := uc.taskRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Atualizar campos se fornecidos
	if req.Title != nil {
		if *req.Title == "" {
			return nil, errors.New("title cannot be empty")
		}
		task.Title = *req.Title
	}

	if req.Description != nil {
		task.Description = *req.Description
	}

	if req.Completed != nil {
		task.Completed = *req.Completed
	}

	if req.Priority != nil {
		if *req.Priority != model.PriorityLow && *req.Priority != model.PriorityMedium && *req.Priority != model.PriorityHigh {
			return nil, ErrInvalidPriority
		}
		task.Priority = *req.Priority
	}

	if req.DueDate != nil && req.DueDate.Before(time.Now()) {
		return nil, errors.New("due date cannot be in the past")
	}

	if req.DueDate != nil {
		task.DueDate = req.DueDate
	}

	if err := uc.taskRepo.Update(task); err != nil {
		return nil, err
	}

	return task.ToResponse(), nil
}

func (uc *taskUseCase) DeleteTask(id uint) error {
	if id == 0 {
		return ErrInvalidTaskID
	}

	// Verificar se o task existe
	_, err := uc.taskRepo.GetByID(id)
	if err != nil {
		return err
	}

	return uc.taskRepo.Delete(id)
}

func (uc *taskUseCase) CompleteTask(id uint) (*model.TaskResponse, error) {
	if id == 0 {
		return nil, ErrInvalidTaskID
	}

	task, err := uc.taskRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if task.Completed {
		return nil, errors.New("task is already completed")
	}

	task.Completed = true

	if err := uc.taskRepo.Update(task); err != nil {
		return nil, err
	}

	return task.ToResponse(), nil
}

func (uc *taskUseCase) GetTasksByStatus(completed bool) ([]*model.TaskResponse, error) {
	tasks, err := uc.taskRepo.GetByCompleted(completed)
	if err != nil {
		return nil, err
	}

	responses := make([]*model.TaskResponse, len(tasks))
	for i, task := range tasks {
		responses[i] = task.ToResponse()
	}

	return responses, nil
}

func (uc *taskUseCase) GetTasksByPriority(priority model.Priority) ([]*model.TaskResponse, error) {
	if priority != model.PriorityLow && priority != model.PriorityMedium && priority != model.PriorityHigh {
		return nil, ErrInvalidPriority
	}

	tasks, err := uc.taskRepo.GetByPriority(priority)
	if err != nil {
		return nil, err
	}

	responses := make([]*model.TaskResponse, len(tasks))
	for i, task := range tasks {
		responses[i] = task.ToResponse()
	}

	return responses, nil
}

func (uc *taskUseCase) GetOverdueTasks() ([]*model.TaskResponse, error) {
	tasks, err := uc.taskRepo.GetAll()
	if err != nil {
		return nil, err
	}

	now := time.Now()
	var overdueTasks []*model.TaskResponse

	for _, task := range tasks {
		if task.DueDate != nil && task.DueDate.Before(now) && !task.Completed {
			overdueTasks = append(overdueTasks, task.ToResponse())
		}
	}

	return overdueTasks, nil
}
