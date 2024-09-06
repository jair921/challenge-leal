package services

import (
	"errors"
	"github.com/jair921/challenge-leal/internal/domain/entities"
	"github.com/jair921/challenge-leal/internal/domain/repositories"
)

type CampaignService interface {
	CreateCampaign(campaign *entities.Campaign) error
	GetByID(id string) (*entities.Campaign, error)
	GetCampaignsByCommerceAndBranch(commerceID, branchID string) ([]*entities.Campaign, error)
	GetActiveCampaigns(commerceID, branchID string) ([]*entities.Campaign, error)
	UpdateCampaign(campaign *entities.Campaign) error
	DeleteCampaign(id string) error
}

type campaignService struct {
	campaignRepo repositories.CampaignRepository
}

// NewCampaignService crea un nuevo servicio de campañas
func NewCampaignService(campaignRepo repositories.CampaignRepository) CampaignService {
	return &campaignService{campaignRepo: campaignRepo}
}

// CreateCampaign crea una nueva campaña
func (s *campaignService) CreateCampaign(campaign *entities.Campaign) error {
	if campaign.StartDate.After(campaign.EndDate) {
		return errors.New("start date cannot be after end date")
	}
	return s.campaignRepo.CreateCampaign(campaign)
}

// GetByID obtiene una campaña por su ID
func (s *campaignService) GetByID(id string) (*entities.Campaign, error) {
	return s.campaignRepo.GetByID(id)
}

// GetCampaignsByCommerceAndBranch obtiene las campañas de un comercio y una sucursal
func (s *campaignService) GetCampaignsByCommerceAndBranch(commerceID, branchID string) ([]*entities.Campaign, error) {
	return s.campaignRepo.GetCampaignsByCommerceAndBranch(commerceID, branchID)
}

// GetActiveCampaigns obtiene las campañas activas
func (s *campaignService) GetActiveCampaigns(commerceID, branchID string) ([]*entities.Campaign, error) {
	return s.campaignRepo.GetActiveCampaigns(commerceID, branchID)
}

// UpdateCampaign actualiza una campaña existente
func (s *campaignService) UpdateCampaign(campaign *entities.Campaign) error {
	if campaign.StartDate.After(campaign.EndDate) {
		return errors.New("start date cannot be after end date")
	}
	return s.campaignRepo.UpdateCampaign(campaign)
}

// DeleteCampaign elimina una campaña por su ID
func (s *campaignService) DeleteCampaign(id string) error {
	return s.campaignRepo.DeleteCampaign(id)
}
