package main

import (
	"log"
	"net/http"

	"github.com/saalcazar/first-api-echo/authorization"
	"github.com/saalcazar/first-api-echo/handler"
	"github.com/saalcazar/first-api-echo/storage"
)

func main() {
	err := authorization.LoadFiles("certificates/app.rsa", "certificates/app.rsa.pub")
	if err != nil {
		log.Fatalf("no se pudo cargar los certificados: %v", err)
	}
	store := storage.NewMemory()
	mux := http.NewServeMux()

	//Grupo de rutas de person

	handler.RoutePerson(mux, &store)
	handler.RouteLogin(mux, &store)

	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Printf("Error en el servidor: %v \n", err)
	}
}
