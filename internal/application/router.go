package application

import (
	"github.com/gorilla/mux"
)

func SetupRouter(deps *Dependencies) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	//campaignController := controllers.NewCampaignController(deps.CampaignRepository, deps.Logger)
	//transactionController := controllers.NewTransactionController(deps.CommerceRepository, deps.Logger)
	//
	//router.HandleFunc("/api/v1/campaigns", campaignController.CreateCampaign).Methods("POST")
	//router.HandleFunc("/api/v1/campaigns/{commerce_id}/{branch_id}", campaignController.GetCampaigns).Methods("GET")
	//router.HandleFunc("/api/v1/transactions", transactionController.AccumulatePoints).Methods("POST")

	return router
}
