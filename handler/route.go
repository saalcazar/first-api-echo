package handler

import (
	"log"
	"net/http"
)

// Manejador de la ruta
func RoutePerson(mux *http.ServeMux, storage Storage) {
	h := newPerson(storage)
	log.Println("Servidor iniciado en el puerto 8080")
	mux.HandleFunc("/v1/persons/create", h.create)
	mux.HandleFunc("/v1/persons/update", h.update)
	mux.HandleFunc("/v1/persons/delete", h.delete)
	mux.HandleFunc("/v1/persons/getall", h.getAll)

}
