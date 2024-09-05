package repositories

import (
	"github.com/jair921/challenge-leal/internal/domain/entities"
)

// CampaignRepository define las operaciones de persistencia para las campañas.
type CampaignRepository interface {
	CreateCampaign(campaign *entities.Campaign) error
	GetCampaignsByCommerceAndBranch(commerceID, branchID string) ([]*entities.Campaign, error)
}
