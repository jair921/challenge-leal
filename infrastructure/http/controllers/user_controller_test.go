package controllers_test

import (
	"github.com/gin-gonic/gin"
	"github.com/jair921/challenge-leal/infrastructure/http/controllers"
	"github.com/jair921/challenge-leal/internal/domain/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// MockUserService es un mock del servicio de usuarios
type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) CreateUser(user *entities.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserService) GetUserByID(id string) (*entities.User, error) {
	args := m.Called(id)
	return args.Get(0).(*entities.User), args.Error(1)
}

func (m *MockUserService) ListUsers() ([]*entities.User, error) {
	args := m.Called()
	return args.Get(0).([]*entities.User), args.Error(1)
}

func TestCreateUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	mockService := new(MockUserService)
	controller := controllers.NewUserController(mockService)

	r.POST("/users", controller.CreateUser)

	// Test case: successful user creation
	mockService.On("CreateUser", mock.MatchedBy(func(u *entities.User) bool {
		return u.Name == "John Doe" && u.Email == "john.doe@example.com"
	})).Return(nil)

	payload := `{"name": "John Doe", "email": "john.doe@example.com"}`
	req, _ := http.NewRequest("POST", "/users", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusCreated, rec.Code)
	mockService.AssertExpectations(t)
}

func TestGetUserByID(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	mockService := new(MockUserService)
	controller := controllers.NewUserController(mockService)

	r.GET("/users/:id", controller.GetUserByID)

	// Test case: successful user retrieval
	mockUser := &entities.User{ID: "1", Name: "John Doe", Email: "john.doe@example.com"}
	mockService.On("GetUserByID", "1").Return(mockUser, nil)

	req, _ := http.NewRequest("GET", "/users/1", nil)
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockService.AssertExpectations(t)
}

func TestListUsers(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	mockService := new(MockUserService)
	controller := controllers.NewUserController(mockService)

	r.GET("/users", controller.ListUsers)

	// Test case: successful users list retrieval
	mockUsers := []*entities.User{
		{ID: "1", Name: "John Doe", Email: "john.doe@example.com"},
		{ID: "2", Name: "Jane Doe", Email: "jane.doe@example.com"},
	}
	mockService.On("ListUsers").Return(mockUsers, nil)

	req, _ := http.NewRequest("GET", "/users", nil)
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockService.AssertExpectations(t)
}
