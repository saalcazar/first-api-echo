package handler

import (
	"log"
	"net/http"

	"github.com/saalcazar/first-api-echo/middleware"
)

// Manejador de la ruta
func RoutePerson(mux *http.ServeMux, storage Storage) {
	h := newPerson(storage)
	log.Println("Servidor iniciado en el puerto 8080")
	mux.HandleFunc("/v1/persons/create", middleware.Log(middleware.Authentication(h.create)))
	mux.HandleFunc("/v1/persons/update", middleware.Log(h.update))
	mux.HandleFunc("/v1/persons/delete", middleware.Log(h.delete))
	mux.HandleFunc("/v1/persons/getall", middleware.Log(h.getAll))

}

// RouteLogin
func RouteLogin(mux *http.ServeMux, storage Storage) {
	h := newLogin(storage)
	mux.HandleFunc("/v1/login", h.login)
}
