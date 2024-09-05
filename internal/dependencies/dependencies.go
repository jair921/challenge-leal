package app

import (
	"database/sql"
	"github.com/jair921/challenge-leal/internal/adapters/repositories"
	repositoriesDomain "github.com/jair921/challenge-leal/internal/domain/repositories"
)

// Dependencies agrupa todas las dependencias de la aplicación
type Dependencies struct {
	CampaignRepository    repositoriesDomain.CampaignRepository
	CommerceRepository    repositoriesDomain.CommerceRepository
	BranchRepository      repositoriesDomain.BranchRepository
	TransactionRepository repositoriesDomain.TransactionRepository
	UserRepository        repositoriesDomain.UserRepository
}

// SetupDependencies inicializa y configura las dependencias de la aplicación
func SetupDependencies(db *sql.DB) *Dependencies {
	campaignRepo := repositories.NewMySQLCampaignRepository(db)
	commerceRepo := repositories.NewMySQLCommerceRepository(db)
	branchRepo := repositories.NewMySQLBranchRepository(db)
	transactionRepo := repositories.NewMySQLTransactionRepository(db)
	userRepo := repositories.NewMySQLUserRepository(db)

	return &Dependencies{
		CampaignRepository:    campaignRepo,
		CommerceRepository:    commerceRepo,
		BranchRepository:      branchRepo,
		TransactionRepository: transactionRepo,
		UserRepository:        userRepo,
	}
}
