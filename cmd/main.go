package main

import (
	"log"
	"net/http"

	"github.com/saalcazar/first-api/handler"
	"github.com/saalcazar/first-api/storage"
)

func main() {
	store := storage.NewMemory()
	mux := http.NewServeMux()

	//Grupo de rutas de person

	handler.RoutePerson(mux, &store)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Printf("Error en el servidor: %v \n", err)
	}
}
