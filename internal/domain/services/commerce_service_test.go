package services

import (
	"errors"
	"testing"

	"github.com/jair921/challenge-leal/internal/domain/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockCommerceRepository es el mock del repositorio de comercios
type MockCommerceRepository struct {
	mock.Mock
}

func (m *MockCommerceRepository) Create(commerce *entities.Commerce) error {
	args := m.Called(commerce)
	return args.Error(0)
}

func (m *MockCommerceRepository) GetByID(id string) (*entities.Commerce, error) {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*entities.Commerce), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockCommerceRepository) List() ([]*entities.Commerce, error) {
	args := m.Called()
	if args.Get(0) != nil {
		return args.Get(0).([]*entities.Commerce), args.Error(1)
	}
	return nil, args.Error(1)
}

// Test para la creación de comercios
func TestCreateCommerce(t *testing.T) {
	mockRepo := new(MockCommerceRepository)
	service := NewCommerceService(mockRepo)

	// Test case: successful commerce creation
	commerce := &entities.Commerce{ID: "1", Name: "Commerce A"}
	mockRepo.On("Create", commerce).Return(nil)

	err := service.CreateCommerce(commerce)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)

	// Test case: missing name
	invalidCommerce := &entities.Commerce{ID: "2"}
	err = service.CreateCommerce(invalidCommerce)
	assert.EqualError(t, err, "name is required")
}

// Test para obtener un comercio por ID
func TestGetCommerceByID(t *testing.T) {
	mockRepo := new(MockCommerceRepository)
	service := NewCommerceService(mockRepo)

	// Test case: successful commerce retrieval
	commerce := &entities.Commerce{ID: "1", Name: "Commerce A"}
	mockRepo.On("GetByID", "1").Return(commerce, nil)

	result, err := service.GetCommerceByID("1")
	assert.NoError(t, err)
	assert.Equal(t, commerce, result)
	mockRepo.AssertExpectations(t)

	// Test case: commerce not found
	mockRepo.On("GetByID", "2").Return(nil, errors.New("commerce not found"))

	result, err = service.GetCommerceByID("2")
	assert.Nil(t, result)
	assert.EqualError(t, err, "commerce not found")
	mockRepo.AssertExpectations(t)
}

// Test para listar todos los comercios
func TestListCommerces(t *testing.T) {
	mockRepo := new(MockCommerceRepository)
	service := NewCommerceService(mockRepo)

	// Test case: successful listing of comerces
	commerces := []*entities.Commerce{
		{ID: "1", Name: "Commerce A"},
		{ID: "2", Name: "Commerce B"},
	}
	mockRepo.On("List").Return(commerces, nil)

	result, err := service.ListCommerces()
	assert.NoError(t, err)
	assert.Equal(t, commerces, result)
	mockRepo.AssertExpectations(t)
}
