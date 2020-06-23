package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/julioc98/pismo/internal/app/account"
)

// AccountHandler Interface
type AccountHandler interface {
	Add(w http.ResponseWriter, r *http.Request)
	FindByID(w http.ResponseWriter, r *http.Request)
}

type accountHandler struct {
	accountService account.Service
}

// NewAccountHandler Create a new handler
func NewAccountHandler(accountService account.Service) AccountHandler {
	return &accountHandler{
		accountService: accountService,
	}
}

// Add a Account
func (ah *accountHandler) Add(w http.ResponseWriter, r *http.Request) {
	var req account.Account
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := ah.accountService.Create(&req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf(`{ "account_id": %d }`, id)))
}

// FindByID a Account
func (ah *accountHandler) FindByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	account, err := ah.accountService.Get(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	response, err := json.Marshal(account)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)

}
