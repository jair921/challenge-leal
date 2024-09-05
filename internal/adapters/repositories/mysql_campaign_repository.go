package repositories

import (
	"database/sql"
	"github.com/jair921/challenge-leal/internal/domain/entities"
	"github.com/jair921/challenge-leal/internal/domain/repositories"
)

type MySQLCampaignRepository struct {
	db *sql.DB
}

func NewMySQLCampaignRepository(db *sql.DB) repositories.CampaignRepository {
	return &MySQLCampaignRepository{db: db}
}

func (r *MySQLCampaignRepository) CreateCampaign(campaign *entities.Campaign) error {
	query := `INSERT INTO campaigns (id, commerce_id, branch_id, start_date, end_date, multiplier) 
	          VALUES (?, ?, ?, ?, ?, ?)`
	_, err := r.db.Exec(query, campaign.ID, campaign.CommerceID, campaign.BranchID, campaign.StartDate, campaign.EndDate, campaign.Multiplier)
	return err
}

func (r *MySQLCampaignRepository) GetCampaignsByCommerceAndBranch(commerceID, branchID string) ([]*entities.Campaign, error) {
	query := `SELECT id, commerce_id, branch_id, start_date, end_date, multiplier 
	          FROM campaigns WHERE commerce_id=? AND branch_id=?`
	rows, err := r.db.Query(query, commerceID, branchID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var campaigns []*entities.Campaign
	for rows.Next() {
		campaign := &entities.Campaign{}
		err := rows.Scan(&campaign.ID, &campaign.CommerceID, &campaign.BranchID, &campaign.StartDate, &campaign.EndDate, &campaign.Multiplier)
		if err != nil {
			return nil, err
		}
		campaigns = append(campaigns, campaign)
	}

	return campaigns, nil
}
