package repositories

import (
	"github.com/jair921/challenge-leal/internal/domain/entities"
)

// CampaignRepository define las operaciones de persistencia para las campañas.
type CampaignRepository interface {
	CreateCampaign(campaign *entities.Campaign) error
	GetActiveCampaigns(commerceID, branchID string) ([]*entities.Campaign, error)
	GetByID(id string) (*entities.Campaign, error)
	GetCampaignsByCommerceAndBranch(commerceID, branchID string) ([]*entities.Campaign, error)
	UpdateCampaign(campaign *entities.Campaign) error
	DeleteCampaign(id string) error
}
