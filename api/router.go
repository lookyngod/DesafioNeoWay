package main

import (
	"api/config"
	"api/rest"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	config.Carregar()

	router := mux.NewRouter()
	r := router.PathPrefix("/api").Subrouter()

	r.HandleFunc("/cpfcnpj", (rest.CpfCnpjHandler))
	r.HandleFunc("/getall", (rest.GetAllHandler))
	r.HandleFunc("/deletedados", (rest.DeleteHandler))
	r.HandleFunc("/database", (rest.CreateTableHandler))
	http.Handle("/", router)

	log.Printf("Padronizando para porta %d", config.Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil); err != nil {
		log.Fatal(err)
	}
}
