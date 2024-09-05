package services

import (
	"errors"
	"github.com/jair921/challenge-leal/internal/domain/entities"
	"github.com/jair921/challenge-leal/internal/domain/repositories"
)

type CampaignService struct {
	campaignRepo repositories.CampaignRepository
}

func NewCampaignService(repo repositories.CampaignRepository) *CampaignService {
	return &CampaignService{campaignRepo: repo}
}

func (s *CampaignService) CreateCampaign(campaign *entities.Campaign) error {
	if campaign.CommerceID == "" || campaign.BranchID == "" {
		return errors.New("invalid campaign: missing commerce or branch ID")
	}
	return s.campaignRepo.CreateCampaign(campaign)
}

func (s *CampaignService) GetCampaignsByCommerceAndBranch(commerceID, branchID string) ([]*entities.Campaign, error) {
	if commerceID == "" || branchID == "" {
		return nil, errors.New("commerceID and branchID are required")
	}
	return s.campaignRepo.GetCampaignsByCommerceAndBranch(commerceID, branchID)
}
