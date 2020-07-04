package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/julioc98/pismo/internal/app/transaction"
)

// TransactionHandler Interface
type TransactionHandler interface {
	Add(w http.ResponseWriter, r *http.Request)
	FindByID(w http.ResponseWriter, r *http.Request)
}

type transactionHandler struct {
	TransactionService transaction.Service
}

// NewTransactionHandler Create a new handler
func NewTransactionHandler(ts transaction.Service) TransactionHandler {
	return &transactionHandler{
		TransactionService: ts,
	}
}

// Add a Transaction
func (th *transactionHandler) Add(w http.ResponseWriter, r *http.Request) {
	var req transaction.Transaction
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = th.TransactionService.Check(&req)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := th.TransactionService.Create(&req)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf(`{ "id": %d }`, id)))
}

// FindByID a Transaction
func (th *transactionHandler) FindByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	Transaction, err := th.TransactionService.Get(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	response, err := json.Marshal(Transaction)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)

}
