package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/saalcazar/first-api-echo/model"
)

// No exportados por que se van a llamar desde la ruta y nadie debe conocerlos
// Crea una estructura Person que tiene dentro de sus campos un storage
type person struct {
	storage Storage
}

func newPerson(storage Storage) person {
	return person{storage}
}

// CREATE
func (p *person) create(c echo.Context) error {
	data := model.Person{}
	err := c.Bind(&data)
	if err != nil {
		response := newResponse(Error, "La persona no tiene una estructura correcta", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	err = p.storage.Create(&data)
	if err != nil {
		response := newResponse(Error, "Hubo un problema al crear la persona", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := newResponse(Message, "Si se pudo crear la persona", nil)
	return c.JSON(http.StatusCreated, response)

}

// Update
func (p *person) update(c echo.Context) error {

	//Query param
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil && ID < 0 {
		response := newResponse(Error, "El ID debe ser un número entero positivo", nil)
		c.JSON(http.StatusBadRequest, response)
	}

	data := model.Person{}
	err = c.Bind(&data)
	if err != nil {
		response := newResponse(Error, "La persona no tiene una estructura correcta", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	err = p.storage.Update(ID, &data)
	if err != nil {
		response := newResponse(Error, "Hubo un problema al crear a la persona", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := newResponse(Message, "Si se pudo actualizar a la persona", nil)
	return c.JSON(http.StatusOK, response)
}

// Delete
func (p *person) delete(c echo.Context) error {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil && ID < 0 {
		response := newResponse(Error, "El ID debe ser un número entero positivo", nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	err = p.storage.Delete(ID)
	if errors.Is(err, model.ErrIDPersonDoesNotExists) {
		response := newResponse(Error, "El ID de la persona no existe", nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	if err != nil {
		response := newResponse(Error, "Hubo un problema al eliminar a la persona", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}
	response := newResponse(Message, "OK", nil)
	return c.JSON(http.StatusOK, response)
}

// GET BY ID
func (p *person) getByID(c echo.Context) error {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response := newResponse(Error, "El ID debe ser un número entero positivo", nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	data, err := p.storage.GetByID(ID)
	if errors.Is(err, model.ErrIDPersonDoesNotExists) {
		response := newResponse(Error, "El ID de la persona no existe", nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	if err != nil {
		response := newResponse(Error, "Hubo un problema al cargar a la persona", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}
	response := newResponse(Message, "OK", data)
	return c.JSON(http.StatusOK, response)

}

// GET ALL
func (p *person) getAll(c echo.Context) error {
	data, err := p.storage.GetAll()
	if err != nil {
		response := newResponse(Error, "Hubo un rpoblema al obtener a todas las personas", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := newResponse(Message, "OK", data)
	return c.JSON(http.StatusOK, response)

}
