package app

import (
	"database/sql"
	"github.com/jair921/challenge-leal/internal/adapters/repositories"
	"github.com/jair921/challenge-leal/internal/domain/services"
)

// Dependencies agrupa todas las dependencias de la aplicación
type Dependencies struct {
	CampaignService    services.CampaignService
	CommerceService    services.CommerceService
	BranchService      services.BranchService
	TransactionService services.TransactionService
	UserService        services.UserService
}

// SetupDependencies inicializa y configura las dependencias de la aplicación
func SetupDependencies(db *sql.DB) *Dependencies {
	campaignRepo := repositories.NewMySQLCampaignRepository(db)
	commerceRepo := repositories.NewMySQLCommerceRepository(db)
	branchRepo := repositories.NewMySQLBranchRepository(db)
	transactionRepo := repositories.NewMySQLTransactionRepository(db)
	userRepo := repositories.NewMySQLUserRepository(db)

	// Inicializar servicios
	campaignService := services.NewCampaignService(campaignRepo)
	commerceService := services.NewCommerceService(commerceRepo)
	branchService := services.NewBranchService(branchRepo)
	transactionService := services.NewTransactionService(transactionRepo, campaignRepo, userRepo)
	userService := services.NewUserService(userRepo)

	return &Dependencies{
		CampaignService:    campaignService,
		CommerceService:    commerceService,
		BranchService:      branchService,
		TransactionService: transactionService,
		UserService:        userService,
	}
}
