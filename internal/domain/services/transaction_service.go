package services

import (
	"errors"
	"github.com/jair921/challenge-leal/internal/domain/entities"
	"github.com/jair921/challenge-leal/internal/domain/repositories"
	"strconv"
	"time"
)

// TransactionService define los métodos relacionados con la lógica de negocio de transacciones.
type TransactionService interface {
	AccumulatePoints(transaction *entities.Transaction) error
	GetTransactionsByUserID(userID string) ([]*entities.Transaction, error)
}

// transactionService es una implementación concreta de TransactionService
type transactionService struct {
	transactionRepo repositories.TransactionRepository
	campaignRepo    repositories.CampaignRepository
	userRepo        repositories.UserRepository
}

// NewTransactionService crea un nuevo servicio de transacciones
func NewTransactionService(transactionRepo repositories.TransactionRepository, campaignRepo repositories.CampaignRepository, userRepo repositories.UserRepository) TransactionService {
	return &transactionService{
		transactionRepo: transactionRepo,
		campaignRepo:    campaignRepo,
		userRepo:        userRepo,
	}
}

// AccumulatePoints acumula puntos para una transacción
func (s *transactionService) AccumulatePoints(transaction *entities.Transaction) error {
	if transaction.UserID == 0 || transaction.CampaignID == 0 || transaction.Amount == 0 {
		return errors.New("user ID, campaign ID and amount are required")
	}

	// Obtener la campaña activa
	campaign, err := s.campaignRepo.GetByID(strconv.Itoa(transaction.CampaignID))
	if err != nil {
		return err
	}
	if campaign == nil || campaign.EndDate.Before(time.Now()) {
		return errors.New("campaign not found or expired")
	}

	user, err := s.userRepo.GetUserByID(strconv.Itoa(transaction.UserID))
	if err != nil {
		return err
	}

	// Calcular puntos ganados
	pointsEarned := (transaction.Amount / 1000) * campaign.Multiplier
	transaction.PointsEarned = pointsEarned
	transaction.CashbackEarned = pointsEarned

	user.Points += pointsEarned

	err = s.transactionRepo.Create(transaction)

	if transaction.ID != "" {
		s.userRepo.UpdateUser(user)
	}

	return err
}

// GetTransactionsByUserID obtiene transacciones por el ID de usuario
func (s *transactionService) GetTransactionsByUserID(userID string) ([]*entities.Transaction, error) {
	return s.transactionRepo.GetByUserID(userID)
}
