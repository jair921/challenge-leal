package controllers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jair921/challenge-leal/internal/domain/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockCommerceService es una implementación mock del servicio de comercios.
type MockCommerceService struct {
	mock.Mock
}

func (m *MockCommerceService) CreateCommerce(commerce *entities.Commerce) error {
	args := m.Called(commerce)
	return args.Error(0)
}

func (m *MockCommerceService) GetCommerceByID(id string) (*entities.Commerce, error) {
	args := m.Called(id)
	return args.Get(0).(*entities.Commerce), args.Error(1)
}

func (m *MockCommerceService) ListCommerces() ([]*entities.Commerce, error) {
	args := m.Called()
	return args.Get(0).([]*entities.Commerce), args.Error(1)
}

func TestCreateCommerce(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	mockService := new(MockCommerceService)
	controller := NewCommerceController(mockService)

	r.POST("/commerces", controller.CreateCommerce)

	// Test case: successful commerce creation
	mockService.On("CreateCommerce", &entities.Commerce{Name: "Test Commerce"}).Return(nil)

	payload := `{"name": "Test Commerce"}`
	req, _ := http.NewRequest("POST", "/commerces", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusCreated, rec.Code)
	mockService.AssertExpectations(t)
}

func TestGetCommerceByID(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	mockService := new(MockCommerceService)
	controller := NewCommerceController(mockService)

	r.GET("/commerces/:id", controller.GetCommerceByID)

	// Test case: commerce found
	mockService.On("GetCommerceByID", "1").Return(&entities.Commerce{ID: "1", Name: "Test Commerce"}, nil)

	req, _ := http.NewRequest("GET", "/commerces/1", nil)
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockService.AssertExpectations(t)
}

func TestListCommerces(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	mockService := new(MockCommerceService)
	controller := NewCommerceController(mockService)

	r.GET("/commerces", controller.ListCommerces)

	// Test case: commerces found
	mockService.On("ListCommerces").Return([]*entities.Commerce{
		{ID: "1", Name: "Commerce 1"},
		{ID: "2", Name: "Commerce 2"},
	}, nil)

	req, _ := http.NewRequest("GET", "/commerces", nil)
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockService.AssertExpectations(t)
}
