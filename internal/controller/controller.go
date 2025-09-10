package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kalyanbasso/taskhub/internal/model"
	"github.com/kalyanbasso/taskhub/internal/usecase"
)

type TaskController struct {
	taskUseCase usecase.TaskUseCase
}

func NewTaskController(taskUseCase usecase.TaskUseCase) *TaskController {
	return &TaskController{
		taskUseCase: taskUseCase,
	}
}

func (tc *TaskController) RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api/v1")
	{
		tasks := api.Group("/tasks")
		{
			tasks.POST("", tc.CreateTask)
			tasks.GET("", tc.GetAllTasks)
			tasks.GET("/:id", tc.GetTaskByID)
			tasks.PUT("/:id", tc.UpdateTask)
			tasks.DELETE("/:id", tc.DeleteTask)
			tasks.PATCH("/:id/complete", tc.CompleteTask)
			tasks.GET("/status/:completed", tc.GetTasksByStatus)
			tasks.GET("/priority/:priority", tc.GetTasksByPriority)
			tasks.GET("/overdue", tc.GetOverdueTasks)
		}
	}
}

func (tc *TaskController) CreateTask(c *gin.Context) {
	var req model.CreateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := tc.taskUseCase.CreateTask(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": task})
}

func (tc *TaskController) GetAllTasks(c *gin.Context) {
	tasks, err := tc.taskUseCase.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": tasks})
}

func (tc *TaskController) GetTaskByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrInvalidTaskID.Error()})
		return
	}

	task, err := tc.taskUseCase.GetTaskByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": task})
}

func (tc *TaskController) UpdateTask(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrInvalidTaskID.Error()})
		return
	}

	var req model.UpdateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := tc.taskUseCase.UpdateTask(uint(id), &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": task})
}

func (tc *TaskController) DeleteTask(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrInvalidTaskID.Error()})
		return
	}

	if err := tc.taskUseCase.DeleteTask(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "task deleted successfully"})
}

func (tc *TaskController) CompleteTask(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrInvalidTaskID.Error()})
		return
	}

	task, err := tc.taskUseCase.CompleteTask(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": task})
}

func (tc *TaskController) GetTasksByStatus(c *gin.Context) {
	completedStr := c.Param("completed")
	completed, err := strconv.ParseBool(completedStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid completed status"})
		return
	}

	tasks, err := tc.taskUseCase.GetTasksByStatus(completed)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": tasks})
}

func (tc *TaskController) GetTasksByPriority(c *gin.Context) {
	priority := model.Priority(c.Param("priority"))

	tasks, err := tc.taskUseCase.GetTasksByPriority(priority)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": tasks})
}

func (tc *TaskController) GetOverdueTasks(c *gin.Context) {
	tasks, err := tc.taskUseCase.GetOverdueTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": tasks})
}
