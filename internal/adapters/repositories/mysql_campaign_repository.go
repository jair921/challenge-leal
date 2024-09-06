package repositories

import (
	"database/sql"
	"github.com/jair921/challenge-leal/internal/domain/entities"
	"github.com/jair921/challenge-leal/internal/domain/repositories"
	"strconv"
	"time"
)

type MySQLCampaignRepository struct {
	db *sql.DB
}

func NewMySQLCampaignRepository(db *sql.DB) repositories.CampaignRepository {
	return &MySQLCampaignRepository{db: db}
}

// CreateCampaign inserta una nueva campaña en la base de datos
func (r *MySQLCampaignRepository) CreateCampaign(campaign *entities.Campaign) error {
	query := `
		INSERT INTO campaigns (commerce_id, branch_id, start_date, end_date, multiplier)
		VALUES (?, ?, ?, ?, ?)
	`
	result, err := r.db.Exec(query, campaign.CommerceID, campaign.BranchID, campaign.StartDate, campaign.EndDate, campaign.Multiplier)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	campaign.ID = strconv.Itoa(int(id))
	return err
}

// GetCampaignsByCommerceAndBranch obtiene las campañas por comercio y sucursal
func (r *MySQLCampaignRepository) GetCampaignsByCommerceAndBranch(commerceID, branchID string) ([]*entities.Campaign, error) {
	query := `
		SELECT id, commerce_id, branch_id, start_date, end_date, multiplier
		FROM campaigns
		WHERE commerce_id = ? AND branch_id = ?
	`
	rows, err := r.db.Query(query, commerceID, branchID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var campaigns []*entities.Campaign
	for rows.Next() {
		campaign := &entities.Campaign{}
		var startDate, endDate time.Time

		err := rows.Scan(
			&campaign.ID,
			&campaign.CommerceID,
			&campaign.BranchID,
			&startDate,
			&endDate,
			&campaign.Multiplier,
		)
		if err != nil {
			return nil, err
		}

		campaign.StartDate = startDate
		campaign.EndDate = endDate

		campaigns = append(campaigns, campaign)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return campaigns, nil
}

// GetActiveCampaigns obtiene campañas activas por comercio y sucursal
func (r *MySQLCampaignRepository) GetActiveCampaigns(commerceID, branchID string) ([]*entities.Campaign, error) {
	query := `
		SELECT id, commerce_id, branch_id, start_date, end_date, multiplier
		FROM campaigns
		WHERE commerce_id = ? AND branch_id = ? AND start_date <= ? AND end_date >= ?
	`
	now := time.Now()
	rows, err := r.db.Query(query, commerceID, branchID, now, now)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var campaigns []*entities.Campaign
	for rows.Next() {
		campaign := &entities.Campaign{}
		err := rows.Scan(
			&campaign.ID,
			&campaign.CommerceID,
			&campaign.BranchID,
			&campaign.StartDate,
			&campaign.EndDate,
			&campaign.Multiplier,
		)
		if err != nil {
			return nil, err
		}
		campaigns = append(campaigns, campaign)
	}
	return campaigns, nil
}

// UpdateCampaign actualiza una campaña existente
func (r *MySQLCampaignRepository) UpdateCampaign(campaign *entities.Campaign) error {
	query := `
		UPDATE campaigns SET commerce_id=?, branch_id=?, start_date=?, end_date=?, multiplier=?
		WHERE id=?
	`
	_, err := r.db.Exec(query, campaign.CommerceID, campaign.BranchID, campaign.StartDate, campaign.EndDate, campaign.Multiplier, campaign.ID)
	return err
}

// DeleteCampaign elimina una campaña
func (r *MySQLCampaignRepository) DeleteCampaign(id string) error {
	query := `DELETE FROM campaigns WHERE id=?`
	_, err := r.db.Exec(query, id)
	return err
}

// GetByID obtiene una campaña por su ID
func (r *MySQLCampaignRepository) GetByID(id string) (*entities.Campaign, error) {
	query := `SELECT id, commerce_id, branch_id, start_date, end_date, multiplier FROM campaigns WHERE id=?`
	campaign := &entities.Campaign{}
	err := r.db.QueryRow(query, id).Scan(
		&campaign.ID,
		&campaign.CommerceID,
		&campaign.BranchID,
		&campaign.StartDate,
		&campaign.EndDate,
		&campaign.Multiplier,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return campaign, nil
}
