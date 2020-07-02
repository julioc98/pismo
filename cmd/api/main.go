package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/julioc98/pismo/cmd/api/handler"
	"github.com/julioc98/pismo/cmd/api/router"
	"github.com/julioc98/pismo/internal/app/account"
	"github.com/julioc98/pismo/internal/app/transaction"
	db "github.com/julioc98/pismo/internal/db"
	"github.com/julioc98/pismo/pkg/env"
	"github.com/julioc98/pismo/pkg/middleware"
)

func handlerHi(w http.ResponseWriter, r *http.Request) {
	msg := "Ola, Seja bem vindo ao Pismo!!"
	log.Println(msg)
	w.Write([]byte(msg))
}

func main() {

	conn := db.Conn()
	defer conn.Close()
	db.Migrate(conn)

	r := mux.NewRouter()
	r.Use(middleware.Logging)

	accountRep := account.NewPostgresRepository(conn)
	accountService := account.NewService(accountRep)
	accountHandler := handler.NewAccountHandler(accountService)

	router.SetAccountRoutes(accountHandler, r.PathPrefix("/accounts").Subrouter())

	transactionRep := transaction.NewPostgresRepository(conn)
	transactionService := transaction.NewService(transactionRep, accountService)
	transactionHandler := handler.NewTransactionHandler(transactionService)
	router.SetTransactionRoutes(transactionHandler, r.PathPrefix("/transactions").Subrouter())

	r.HandleFunc("/", handlerHi)
	http.Handle("/", r)

	port := env.Get("PORT", "5001")
	log.Printf(`%s listening on port: %s `, env.Get("APP", "pismo"), port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
