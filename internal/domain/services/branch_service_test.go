package services

import (
	"errors"
	"github.com/jair921/challenge-leal/internal/domain/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

// MockBranchRepository es un mock del repositorio de sucursales
type MockBranchRepository struct {
	mock.Mock
}

func (m *MockBranchRepository) Create(branch *entities.Branch) error {
	args := m.Called(branch)
	return args.Error(0)
}

func (m *MockBranchRepository) GetByID(id string) (*entities.Branch, error) {
	args := m.Called(id)
	return args.Get(0).(*entities.Branch), args.Error(1)
}

func (m *MockBranchRepository) List() ([]*entities.Branch, error) {
	args := m.Called()
	return args.Get(0).([]*entities.Branch), args.Error(1)
}

func TestCreateBranch(t *testing.T) {
	mockRepo := new(MockBranchRepository)
	service := NewBranchService(mockRepo)

	// Test case: successful branch creation
	branch := &entities.Branch{Name: "Branch A", CommerceID: "123"}
	mockRepo.On("Create", branch).Return(nil)

	err := service.CreateBranch(branch)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)

	// Test case: missing required fields
	invalidBranch := &entities.Branch{}
	err = service.CreateBranch(invalidBranch)
	assert.EqualError(t, err, "name and commerce ID are required")
}

func TestGetBranchByID(t *testing.T) {
	mockRepo := new(MockBranchRepository)
	service := NewBranchService(mockRepo)

	// Test case: successful branch retrieval
	mockBranch := &entities.Branch{ID: "1", Name: "Branch A", CommerceID: "123"}
	mockRepo.On("GetByID", "1").Return(mockBranch, nil)

	branch, err := service.GetBranchByID("1")
	assert.NoError(t, err)
	assert.Equal(t, mockBranch, branch)
	mockRepo.AssertExpectations(t)

	// Test case: branch not found
	mockRepo.On("GetByID", "2").Return((*entities.Branch)(nil), errors.New("branch not found")) // Aquí se asegura que retorna nil correctamente

	branch, err = service.GetBranchByID("2")
	assert.Nil(t, branch)
	assert.EqualError(t, err, "branch not found")
	mockRepo.AssertExpectations(t)
}

func TestListBranches(t *testing.T) {
	mockRepo := new(MockBranchRepository)
	service := NewBranchService(mockRepo)

	// Test case: successful branch listing
	mockBranches := []*entities.Branch{
		{ID: "1", Name: "Branch A", CommerceID: "123"},
		{ID: "2", Name: "Branch B", CommerceID: "456"},
	}

	// El primer caso: se simula una respuesta exitosa.
	mockRepo.On("List").Return(mockBranches, nil).Once()

	branches, err := service.ListBranches()
	assert.NoError(t, err)
	assert.Equal(t, mockBranches, branches)
	mockRepo.AssertExpectations(t)
}
