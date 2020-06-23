package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/julioc98/pismo/pkg/env"
	"github.com/julioc98/pismo/pkg/middleware"
)

func handlerHi(w http.ResponseWriter, r *http.Request) {
	msg := "Ola, Seja bem vindo ao Pismo!!"
	log.Println(msg)
	w.Write([]byte(msg))
}

func main() {

	r := mux.NewRouter()
	r.Use(middleware.Logging)

	r.HandleFunc("/", handlerHi)
	http.Handle("/", r)

	port := env.Get("PORT", "5001")
	log.Printf(`%s listening on port: %s `, env.Get("APP", "pismo"), port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
