package router

import (
	"github.com/gorilla/mux"
	"github.com/julioc98/pismo/cmd/api/handler"
)

// SetAccountRoutes add routes from Account
func SetAccountRoutes(ah handler.AccountHandler, r *mux.Router) {
	r.HandleFunc("", ah.Add).Methods("POST")
	r.HandleFunc("/{id:[0-9]+}", ah.FindByID).Methods("GET")
}
