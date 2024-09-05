package entities

import "time"

// Campaign representa una campaña de puntos o cashback para un comercio y sucursal
type Campaign struct {
	ID         string    `json:"id"`
	CommerceID string    `json:"commerce_id"`
	BranchID   string    `json:"branch_id"`
	StartDate  time.Time `json:"start_date"`
	EndDate    time.Time `json:"end_date"`
	Multiplier float64   `json:"multiplier"` // Factor para puntos o cashback
}
