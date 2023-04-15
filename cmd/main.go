package main

import (
	"log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

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

	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	handler.RouteLogin(e, &store)
	handler.RoutePerson(e, &store)

	err = e.Start(":8080")
	if err != nil {
		log.Printf("Error en el servidor: %v \n", err)
	}
}
