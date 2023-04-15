package handler

import (
	"github.com/labstack/echo"
	"github.com/saalcazar/first-api-echo/middleware"
)

// Manejador de la ruta
func RoutePerson(e *echo.Echo, storage Storage) {
	h := newPerson(storage)
	person := e.Group("/v1/persons")
	person.Use(middleware.Authentication)
	person.POST("", h.create)
	person.PUT("/:id", h.update)
	person.GET("", h.getAll)
	person.GET("/:id", h.getByID)
	person.DELETE("/:id", h.delete)
}

// RouteLogin
func RouteLogin(e *echo.Echo, storage Storage) {
	h := newLogin(storage)
	e.POST("/v1/login", h.login)
}
