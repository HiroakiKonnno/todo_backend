package handler

import (
	"net/http"

	"todo_backend/internal/model"
	"todo_backend/internal/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)
func RegisterTaskRoutes(r *gin.Engine, db *gorm.DB) {
	taskRepo := repository.NewTaskRepository(db)
	r.GET("/api/tasks", GetAllTasks(taskRepo))
	r.POST("/api/tasks", CreateTask(taskRepo))
}

// タスクの一覧を取得する
func GetAllTasks(repo *repository.TaskRepositoryImpl) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tasks, err := repo.GetAllTasks()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, tasks)
	}
}

// タスクを作成する
func CreateTask(repo *repository.TaskRepositoryImpl) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var task model.Task
		if err := ctx.ShouldBindJSON(&task); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := repo.CreateTask(&task)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusCreated, task)
	}
}
