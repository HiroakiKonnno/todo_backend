package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"todo_backend/internal/handler"
	"todo_backend/internal/model"
	"todo_backend/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type MockTaskRepository struct{}

func (m *MockTaskRepository) CreateTask(task *model.Task) error {
	return nil
}

func setupTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect test db")
	}
	db.AutoMigrate(&model.Task{})
	return db
}

func TestCreateTask_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	db := setupTestDB()
	repo := &repository.TaskRepositoryImpl{DB: db}

	input := model.Task{
		Title:     "test task",
		Content:   "test description",
		UserId:    1,
	}
	body, _ := json.Marshal(input)

	req, err := http.NewRequest(http.MethodPost, "/api/tasks", bytes.NewBuffer(body))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	r := gin.Default()

	r.POST("/api/tasks", handler.CreateTask(repo))

	r.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusCreated, rec.Code)
	var response model.Task
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, input.Title, response.Title)
}