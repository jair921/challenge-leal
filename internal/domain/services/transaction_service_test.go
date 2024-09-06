package services

import (
	"errors"
	_ "strconv"
	"testing"
	"time"

	"github.com/jair921/challenge-leal/internal/domain/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockTransactionRepository es el mock del repositorio de transacciones
type MockTransactionRepository struct {
	mock.Mock
}

func (m *MockTransactionRepository) Create(transaction *entities.Transaction) error {
	args := m.Called(transaction)
	return args.Error(0)
}

func (m *MockTransactionRepository) GetByUserID(userID string) ([]*entities.Transaction, error) {
	args := m.Called(userID)
	if args.Get(0) != nil {
		return args.Get(0).([]*entities.Transaction), args.Error(1)
	}
	return nil, args.Error(1)
}

// Test para acumular puntos
func TestAccumulatePoints(t *testing.T) {
	mockTransactionRepo := new(MockTransactionRepository)
	mockCampaignRepo := new(MockCampaignRepository)
	mockUserRepo := new(MockUserRepository)
	service := NewTransactionService(mockTransactionRepo, mockCampaignRepo, mockUserRepo)

	// Test case: successful points accumulation
	campaign := &entities.Campaign{ID: "1", Multiplier: 1, EndDate: time.Now().Add(time.Hour)}
	user := &entities.User{ID: "1", Points: 0}
	transaction := &entities.Transaction{UserID: 1, CampaignID: 1, Amount: 2000}

	mockCampaignRepo.On("GetByID", "1").Return(campaign, nil)
	mockUserRepo.On("GetUserByID", "1").Return(user, nil)
	mockTransactionRepo.On("Create", transaction).Return(nil)
	mockUserRepo.On("UpdateUser", user).Return(nil)

	err := service.AccumulatePoints(transaction)
	assert.NoError(t, err)
	assert.Equal(t, 2.0, transaction.PointsEarned)
	assert.Equal(t, 2.0, transaction.CashbackEarned)
	assert.Equal(t, 2.0, user.Points)
}

// Test para obtener transacciones por ID de usuario
func TestGetTransactionsByUserID(t *testing.T) {
	mockTransactionRepo := new(MockTransactionRepository)
	service := NewTransactionService(mockTransactionRepo, nil, nil)

	// Test case: successful retrieval of transactions
	transactions := []*entities.Transaction{
		{ID: "1", UserID: 1, Amount: 2000},
		{ID: "2", UserID: 1, Amount: 3000},
	}
	mockTransactionRepo.On("GetByUserID", "1").Return(transactions, nil)

	result, err := service.GetTransactionsByUserID("1")
	assert.NoError(t, err)
	assert.Equal(t, transactions, result)
	mockTransactionRepo.AssertExpectations(t)

	// Test case: error in retrieving transactions
	mockTransactionRepo.On("GetByUserID", "2").Return(nil, errors.New("error retrieving transactions"))

	result, err = service.GetTransactionsByUserID("2")
	assert.Nil(t, result)
	assert.EqualError(t, err, "error retrieving transactions")
	mockTransactionRepo.AssertExpectations(t)
}
