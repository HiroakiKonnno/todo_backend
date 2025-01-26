package handler

import (
	"net/http"
	"strconv"

	"todo_backend/internal/model"
	"todo_backend/internal/repository"

	"os"

	"encoding/csv"

	"time"

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
	r.POST("/api/tasks/export", exportCSV(taskRepo))
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

func exportCSV(repo *repository.TaskRepositoryImpl) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		go func() {
			file, err := os.Create("task.csv")
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			tasks, err := repo.GetAllTasks()
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			data := [][]string{
				{"ID", "Title", "Content", "StartDate", "EndDate"},
			}

			for _, task := range tasks {
				data = append(data, []string{
					strconv.Itoa(task.ID),
					task.Title,
					task.Content,
					task.StartDate.Format(time.RFC3339),
					task.EndDate.Format(time.RFC3339),
				})
			}
			
			defer file.Close()
			w := csv.NewWriter(file)

			defer w.Flush()
			for _, record := range data {
				if err := w.Write(record); err != nil {
					ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
					return
				}
			}
		}()
		ctx.File("task.csv")
	}
}

func HelloWorld() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	}
}
