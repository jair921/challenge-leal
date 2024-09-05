package entities

import "time"

// Transaction representa una transacción donde se acumulan puntos o cashback
type Transaction struct {
	ID             string    `json:"id"`
	UserID         string    `json:"user_id"`
	CommerceID     string    `json:"commerce_id"`
	BranchID       string    `json:"branch_id"`
	Amount         float64   `json:"amount"`
	PointsEarned   float64   `json:"points_earned"`
	CashbackEarned float64   `json:"cashback_earned"`
	CreatedAt      time.Time `json:"created_at"`
}
