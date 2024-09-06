package controllers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jair921/challenge-leal/internal/domain/entities"
	_ "github.com/jair921/challenge-leal/internal/domain/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockBranchService es una implementación mock del servicio de sucursales.
type MockBranchService struct {
	mock.Mock
}

func (m *MockBranchService) CreateBranch(branch *entities.Branch) error {
	args := m.Called(branch)
	return args.Error(0)
}

func (m *MockBranchService) GetBranchByID(id string) (*entities.Branch, error) {
	args := m.Called(id)
	return args.Get(0).(*entities.Branch), args.Error(1)
}

func (m *MockBranchService) ListBranches() ([]*entities.Branch, error) {
	args := m.Called()
	return args.Get(0).([]*entities.Branch), args.Error(1)
}

func TestCreateBranch(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	mockService := new(MockBranchService)
	controller := NewBranchController(mockService)

	r.POST("/branches", controller.CreateBranch)

	// Test case: successful branch creation
	mockService.On("CreateBranch", &entities.Branch{Name: "Test Branch", CommerceID: "1"}).Return(nil)

	payload := `{"name": "Test Branch", "commerce_id": "1"}`
	req, _ := http.NewRequest("POST", "/branches", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusCreated, rec.Code)
	mockService.AssertExpectations(t)
}

func TestGetBranchByID(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	mockService := new(MockBranchService)
	controller := NewBranchController(mockService)

	r.GET("/branches/:id", controller.GetBranchByID)

	// Test case: branch found
	mockService.On("GetBranchByID", "1").Return(&entities.Branch{ID: "1", Name: "Test Branch", CommerceID: "1"}, nil)

	req, _ := http.NewRequest("GET", "/branches/1", nil)
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockService.AssertExpectations(t)
}

func TestListBranches(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	mockService := new(MockBranchService)
	controller := NewBranchController(mockService)

	r.GET("/branches", controller.ListBranches)

	// Test case: list branches
	mockService.On("ListBranches").Return([]*entities.Branch{
		{ID: "1", Name: "Branch 1", CommerceID: "1"},
		{ID: "2", Name: "Branch 2", CommerceID: "2"},
	}, nil)

	req, _ := http.NewRequest("GET", "/branches", nil)
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockService.AssertExpectations(t)
}
