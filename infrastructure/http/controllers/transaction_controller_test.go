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

// MockTransactionService es una implementación mock del servicio de transacciones.
type MockTransactionService struct {
	mock.Mock
}

func (m *MockTransactionService) AccumulatePoints(transaction *entities.Transaction) error {
	args := m.Called(transaction)
	return args.Error(0)
}

func (m *MockTransactionService) GetTransactionsByUserID(userID string) ([]*entities.Transaction, error) {
	args := m.Called(userID)
	return args.Get(0).([]*entities.Transaction), args.Error(1)
}

func TestAccumulatePoints(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	mockService := new(MockTransactionService)
	controller := NewTransactionController(mockService)

	r.POST("/transactions", controller.AccumulatePoints)

	// Test case: successful points accumulation
	mockService.On("AccumulatePoints", mock.MatchedBy(func(t *entities.Transaction) bool {
		return t.UserID == 1 && t.CommerceID == 1 && t.BranchID == 1 && t.CampaignID == 1 && t.Amount == 100
	})).Return(nil)

	payload := `{
		"user_id": 1,
		"commerce_id": 1,
		"branch_id": 1,
		"campaign_id": 1,
		"amount": 100,
		"points_earned": 10,
		"cashback_earned": 5
	}`
	req, _ := http.NewRequest("POST", "/transactions", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusCreated, rec.Code)
	mockService.AssertExpectations(t)
}

func TestGetTransactionsByUserID(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	mockService := new(MockTransactionService)
	controller := NewTransactionController(mockService)

	r.GET("/transactions/:userID", controller.GetTransactionsByUserID)

	// Test case: transactions found
	mockService.On("GetTransactionsByUserID", "1").Return([]*entities.Transaction{
		{
			ID:             "1",
			UserID:         1,
			CommerceID:     1,
			BranchID:       1,
			CampaignID:     1,
			Amount:         100,
			PointsEarned:   10,
			CashbackEarned: 5,
			CreatedAt:      time.Now(),
		},
		{
			ID:             "2",
			UserID:         1,
			CommerceID:     1,
			BranchID:       1,
			CampaignID:     1,
			Amount:         50,
			PointsEarned:   5,
			CashbackEarned: 2.5,
			CreatedAt:      time.Now(),
		},
	}, nil)

	req, _ := http.NewRequest("GET", "/transactions/1", nil)
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockService.AssertExpectations(t)
}
