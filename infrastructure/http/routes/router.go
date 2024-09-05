package routes

import (
	"github.com/gorilla/mux"
	"github.com/jair921/challenge-leal/internal/dependencies"
)

func NewRouter(deps *dependencies.Dependencies) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	//router.HandleFunc("/api/v1/campaigns", controllers.CreateCampaign).Methods("POST")
	//router.HandleFunc("/api/v1/campaigns/{commerce_id}/{branch_id}", controllers.GetCampaigns).Methods("GET")
	//router.HandleFunc("/api/v1/transactions", controllers.AccumulatePoints).Methods("POST")

	return router
}
