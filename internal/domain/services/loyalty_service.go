package services

import (
	"errors"
	"github.com/jair921/challenge-leal/internal/domain/entities"
	"github.com/jair921/challenge-leal/internal/domain/repositories"
)

type LoyaltyService struct {
	transactionRepo repositories.TransactionRepository
	userRepo        repositories.UserRepository
}

func NewLoyaltyService(transRepo repositories.TransactionRepository, userRepo repositories.UserRepository) *LoyaltyService {
	return &LoyaltyService{transactionRepo: transRepo, userRepo: userRepo}
}

func (s *LoyaltyService) AccumulatePoints(transaction *entities.Transaction) error {
	if transaction.Amount <= 0 {
		return errors.New("invalid transaction amount")
	}

	// Logica para calcular puntos o cashback
	transaction.PointsEarned = transaction.Amount * 0.05   // Ejemplo: 5% en puntos
	transaction.CashbackEarned = transaction.Amount * 0.03 // Ejemplo: 3% en cashback

	return s.transactionRepo.AccumulatePoints(transaction)
}

func (s *LoyaltyService) RedeemPoints(userID string, points float64) error {
	user, err := s.userRepo.GetUserByID(userID)
	if err != nil {
		return err
	}

	if user.Points < points {
		return errors.New("not enough points to redeem")
	}

	user.Points -= points
	return s.userRepo.UpdateUser(user)
}
