package controllers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jair921/challenge-leal/internal/domain/entities"
	_ "github.com/jair921/challenge-leal/internal/domain/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockCampaignService es una implementación mock del servicio de campañas.
type MockCampaignService struct {
	mock.Mock
}

func (m *MockCampaignService) CreateCampaign(campaign *entities.Campaign) error {
	args := m.Called(campaign)
	return args.Error(0)
}

func (m *MockCampaignService) GetByID(id string) (*entities.Campaign, error) {
	args := m.Called(id)
	return args.Get(0).(*entities.Campaign), args.Error(1)
}

func (m *MockCampaignService) GetCampaignsByCommerceAndBranch(commerceID, branchID string) ([]*entities.Campaign, error) {
	args := m.Called(commerceID, branchID)
	return args.Get(0).([]*entities.Campaign), args.Error(1)
}

func (m *MockCampaignService) GetActiveCampaigns(commerceID, branchID string) ([]*entities.Campaign, error) {
	args := m.Called(commerceID, branchID)
	return args.Get(0).([]*entities.Campaign), args.Error(1)
}

func (m *MockCampaignService) UpdateCampaign(campaign *entities.Campaign) error {
	args := m.Called(campaign)
	return args.Error(0)
}

func (m *MockCampaignService) DeleteCampaign(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestCreateCampaign(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	mockService := new(MockCampaignService)
	controller := NewCampaignController(mockService)

	r.POST("/campaigns", controller.CreateCampaign)

	// Test case: successful campaign creation
	campaign := &entities.Campaign{
		CommerceID: 1,
		BranchID:   1,
		StartDate:  time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		EndDate:    time.Date(2024, 12, 31, 23, 59, 59, 0, time.UTC),
		Multiplier: 1.5,
	}

	mockService.On("CreateCampaign", campaign).Return(nil)

	payload := `{"commerce_id": 1, "branch_id": 1, "start_date": "2024-01-01T00:00:00Z", "end_date": "2024-12-31T23:59:59Z", "multiplier": 1.5}`
	req, _ := http.NewRequest("POST", "/campaigns", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusCreated, rec.Code)
	mockService.AssertExpectations(t)
}

func TestGetByID(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	mockService := new(MockCampaignService)
	controller := NewCampaignController(mockService)

	r.GET("/campaigns/:id", controller.GetByID)

	// Test case: campaign found
	mockService.On("GetByID", "1").Return(&entities.Campaign{ID: "1", CommerceID: 1, BranchID: 1}, nil)

	req, _ := http.NewRequest("GET", "/campaigns/1", nil)
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockService.AssertExpectations(t)
}

func TestGetCampaignsByCommerceAndBranch(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	mockService := new(MockCampaignService)
	controller := NewCampaignController(mockService)

	r.GET("/campaigns/:commerceID/:branchID", controller.GetCampaignsByCommerceAndBranch)

	// Test case: campaigns found
	mockService.On("GetCampaignsByCommerceAndBranch", "1", "1").Return([]*entities.Campaign{
		{ID: "1", CommerceID: 1, BranchID: 1},
		{ID: "2", CommerceID: 1, BranchID: 1},
	}, nil)

	req, _ := http.NewRequest("GET", "/campaigns/1/1", nil)
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockService.AssertExpectations(t)
}

func TestGetActiveCampaigns(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	mockService := new(MockCampaignService)
	controller := NewCampaignController(mockService)

	r.GET("/campaigns/active/:commerceID/:branchID", controller.GetActiveCampaigns)

	// Test case: active campaigns found
	mockService.On("GetActiveCampaigns", "1", "1").Return([]*entities.Campaign{
		{ID: "1", CommerceID: 1, BranchID: 1},
	}, nil)

	req, _ := http.NewRequest("GET", "/campaigns/active/1/1", nil)
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockService.AssertExpectations(t)
}

func TestUpdateCampaign(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	mockService := new(MockCampaignService)
	controller := NewCampaignController(mockService)

	r.PUT("/campaigns", controller.UpdateCampaign)

	// Test case: successful campaign update
	campaign := &entities.Campaign{
		ID:         "1",
		CommerceID: 1,
		BranchID:   1,
		StartDate:  time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		EndDate:    time.Date(2024, 12, 31, 23, 59, 59, 0, time.UTC),
		Multiplier: 1.5,
	}

	mockService.On("UpdateCampaign", campaign).Return(nil)

	payload := `{"id": "1", "commerce_id": 1, "branch_id": 1, "start_date": "2024-01-01T00:00:00Z", "end_date": "2024-12-31T23:59:59Z", "multiplier": 1.5}`
	req, _ := http.NewRequest("PUT", "/campaigns", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockService.AssertExpectations(t)
}

func TestDeleteCampaign(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	mockService := new(MockCampaignService)
	controller := NewCampaignController(mockService)

	r.DELETE("/campaigns/:id", controller.DeleteCampaign)

	// Test case: successful campaign deletion
	mockService.On("DeleteCampaign", "1").Return(nil)

	req, _ := http.NewRequest("DELETE", "/campaigns/1", nil)
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusNoContent, rec.Code)
	mockService.AssertExpectations(t)
}
