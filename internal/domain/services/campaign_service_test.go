package services

import (
	"errors"
	"testing"
	"time"

	"github.com/jair921/challenge-leal/internal/domain/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockCampaignRepository es el mock del repositorio de campañas
type MockCampaignRepository struct {
	mock.Mock
}

func (m *MockCampaignRepository) CreateCampaign(campaign *entities.Campaign) error {
	args := m.Called(campaign)
	return args.Error(0)
}

func (m *MockCampaignRepository) GetByID(id string) (*entities.Campaign, error) {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*entities.Campaign), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockCampaignRepository) GetCampaignsByCommerceAndBranch(commerceID, branchID string) ([]*entities.Campaign, error) {
	args := m.Called(commerceID, branchID)
	if args.Get(0) != nil {
		return args.Get(0).([]*entities.Campaign), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockCampaignRepository) GetActiveCampaigns(commerceID, branchID string) ([]*entities.Campaign, error) {
	args := m.Called(commerceID, branchID)
	if args.Get(0) != nil {
		return args.Get(0).([]*entities.Campaign), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockCampaignRepository) UpdateCampaign(campaign *entities.Campaign) error {
	args := m.Called(campaign)
	return args.Error(0)
}

func (m *MockCampaignRepository) DeleteCampaign(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

// Test para la creación de campañas
func TestCreateCampaign(t *testing.T) {
	mockRepo := new(MockCampaignRepository)
	service := NewCampaignService(mockRepo)

	// Test case: successful campaign creation
	campaign := &entities.Campaign{
		ID:         "1",
		CommerceID: 123,
		BranchID:   456,
		StartDate:  time.Now(),
		EndDate:    time.Now().Add(24 * time.Hour),
		Multiplier: 1.5,
	}
	mockRepo.On("CreateCampaign", campaign).Return(nil)

	err := service.CreateCampaign(campaign)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)

	// Test case: start date after end date
	invalidCampaign := &entities.Campaign{
		ID:         "2",
		CommerceID: 123,
		BranchID:   456,
		StartDate:  time.Now().Add(24 * time.Hour),
		EndDate:    time.Now(),
		Multiplier: 1.5,
	}
	err = service.CreateCampaign(invalidCampaign)
	assert.EqualError(t, err, "start date cannot be after end date")
}

// Test para obtener campañas por ID
func TestGetCampaignByID(t *testing.T) {
	mockRepo := new(MockCampaignRepository)
	service := NewCampaignService(mockRepo)

	// Test case: successful campaign retrieval
	campaign := &entities.Campaign{ID: "1", CommerceID: 123, BranchID: 456}
	mockRepo.On("GetByID", "1").Return(campaign, nil)

	result, err := service.GetByID("1")
	assert.NoError(t, err)
	assert.Equal(t, campaign, result)
	mockRepo.AssertExpectations(t)

	// Test case: campaign not found
	mockRepo.On("GetByID", "2").Return(nil, errors.New("campaign not found"))

	result, err = service.GetByID("2")
	assert.Nil(t, result)
	assert.EqualError(t, err, "campaign not found")
	mockRepo.AssertExpectations(t)
}

// Test para obtener campañas por comercio y sucursal
func TestGetCampaignsByCommerceAndBranch(t *testing.T) {
	mockRepo := new(MockCampaignRepository)
	service := NewCampaignService(mockRepo)

	// Test case: successful campaigns retrieval
	campaigns := []*entities.Campaign{
		{ID: "1", CommerceID: 123, BranchID: 456},
		{ID: "2", CommerceID: 123, BranchID: 456},
	}
	mockRepo.On("GetCampaignsByCommerceAndBranch", "123", "456").Return(campaigns, nil)

	result, err := service.GetCampaignsByCommerceAndBranch("123", "456")
	assert.NoError(t, err)
	assert.Equal(t, campaigns, result)
	mockRepo.AssertExpectations(t)
}

// Test para la actualización de campañas
func TestUpdateCampaign(t *testing.T) {
	mockRepo := new(MockCampaignRepository)
	service := NewCampaignService(mockRepo)

	// Test case: successful campaign update
	campaign := &entities.Campaign{
		ID:         "1",
		CommerceID: 123,
		BranchID:   456,
		StartDate:  time.Now(),
		EndDate:    time.Now().Add(24 * time.Hour),
		Multiplier: 2.0,
	}
	mockRepo.On("UpdateCampaign", campaign).Return(nil)

	err := service.UpdateCampaign(campaign)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)

	// Test case: start date after end date
	invalidCampaign := &entities.Campaign{
		ID:         "1",
		CommerceID: 123,
		BranchID:   456,
		StartDate:  time.Now().Add(24 * time.Hour),
		EndDate:    time.Now(),
		Multiplier: 2.0,
	}
	err = service.UpdateCampaign(invalidCampaign)
	assert.EqualError(t, err, "start date cannot be after end date")
}

// Test para eliminar campañas
func TestDeleteCampaign(t *testing.T) {
	mockRepo := new(MockCampaignRepository)
	service := NewCampaignService(mockRepo)

	// Test case: successful campaign deletion
	mockRepo.On("DeleteCampaign", "1").Return(nil)

	err := service.DeleteCampaign("1")
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)

	// Test case: campaign not found during deletion
	mockRepo.On("DeleteCampaign", "2").Return(errors.New("campaign not found"))

	err = service.DeleteCampaign("2")
	assert.EqualError(t, err, "campaign not found")
	mockRepo.AssertExpectations(t)
}
