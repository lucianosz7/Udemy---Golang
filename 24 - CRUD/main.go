package main

import (
	"crud/servidor"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/users", servidor.CriarUsuario).Methods(http.MethodPost)
	router.HandleFunc("/users/{id}", servidor.BuscarUsuario).Methods(http.MethodGet)
	router.HandleFunc("/users", servidor.BuscarUsuarios).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", servidor.AtualizarUsuario).Methods(http.MethodPut)
	router.HandleFunc("/users/{id}", servidor.DeletarUsuario).Methods(http.MethodDelete)

	fmt.Println("Listening port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
