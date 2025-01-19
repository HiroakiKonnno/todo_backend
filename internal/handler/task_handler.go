package handler

import (
	"net/http"
	"strconv"

	"todo_backend/internal/model"
	"todo_backend/internal/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)
func RegisterTaskRoutes(r *gin.RouterGroup, db *gorm.DB) {
	taskRepo := repository.NewTaskRepository(db)
	r.GET("/api/helloworld", HelloWorld())
	r.GET("/api/tasks", GetAllTasks(taskRepo))
	r.GET("/api/tasks/:id", GetTask(taskRepo))
	r.POST("/api/tasks", CreateTask(taskRepo))
	r.PATCH("/api/tasks/:id", UpdateTask(taskRepo))
	r.DELETE("/api/tasks/:id", DeleteTask(taskRepo))
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

// タスクを取得する
func GetTask(repo *repository.TaskRepositoryImpl) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err !=nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "IDが無効です"})
			return
		}
		task, err := repo.GetTask(id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, task)
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
		if err := task.Validate(); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Validation failed", "details": err.Error()})
		}
		err := repo.CreateTask(&task)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusCreated, task)
	}
}

// タスクを更新する
func UpdateTask(repo *repository.TaskRepositoryImpl) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err !=nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "IDが無効です"})
			return
		}

		var fields map[string]interface{}
		if err := ctx.ShouldBindJSON(&fields); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "無効なリクエストボディ"})
			return
		}

		repo.UpdateTaskFields(id, fields)

		updatedTask, err := repo.GetTask(id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, updatedTask)
	}
}

// タスクを更新する
func DeleteTask(repo *repository.TaskRepositoryImpl) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err !=nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "IDが無効です"})
			return
		}

		if err := repo.DeleteTask(id); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, err)
	}
}
func HelloWorld() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	}
}
