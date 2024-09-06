package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jair921/challenge-leal/infrastructure/http/controllers"
	"github.com/jair921/challenge-leal/internal/dependencies"
)

// NewRouter crea y devuelve un nuevo router Gin con todas las rutas configuradas.
func NewRouter(deps *app.Dependencies) *gin.Engine {
	router := gin.Default()

	// Controladores
	branchController := controllers.NewBranchController(deps.BranchService)
	commerceController := controllers.NewCommerceController(deps.CommerceService)
	campaignController := controllers.NewCampaignController(deps.CampaignService)
	transactionController := controllers.NewTransactionController(deps.TransactionService)
	userController := controllers.NewUserController(deps.UserService)

	// Rutas para Branch
	router.POST("/branches", branchController.CreateBranch)
	router.GET("/branches/:id", branchController.GetBranchByID)
	router.GET("/branches", branchController.ListBranches)

	// Rutas para Commerce
	router.POST("/commerces", commerceController.CreateCommerce)
	router.GET("/commerces/:id", commerceController.GetCommerceByID)
	router.GET("/commerces", commerceController.ListCommerces)

	// Rutas para Campaign
	router.POST("/campaigns", campaignController.CreateCampaign)
	router.GET("/campaigns/:id", campaignController.GetByID)
	router.GET("/campaigns/all/:commerceID/:branchID", campaignController.GetCampaignsByCommerceAndBranch)
	router.GET("/campaigns/active/:commerceID/:branchID", campaignController.GetActiveCampaigns)
	router.PUT("/campaigns/:id", campaignController.UpdateCampaign)
	router.DELETE("/campaigns/:id", campaignController.DeleteCampaign)

	// Rutas para Transaction
	router.POST("/transactions", transactionController.AccumulatePoints)
	router.GET("/transactions/user/:userID", transactionController.GetTransactionsByUserID)

	// Rutas para User
	router.POST("/users", userController.CreateUser)
	router.GET("/users/:id", userController.GetUserByID)
	router.GET("/users", userController.ListUsers)

	return router
}
