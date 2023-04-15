package storage

import (
	"fmt"

	"github.com/saalcazar/first-api-echo/model"
)

// Memory
type Memory struct {
	currentID int
	Persons   map[int]model.Person
}

// New Memory
func NewMemory() Memory {
	persons := make(map[int]model.Person) //Inicialización del mapa
	return Memory{
		currentID: 0,
		Persons:   persons,
	}
}

//CRUD

// Create
func (m *Memory) Create(person *model.Person) error {
	if person == nil {
		return model.ErrPersonCanNotBeNil
	}
	m.currentID++
	m.Persons[m.currentID] = *person
	return nil
}

// Update actualiza una persona en el slice de memoria
func (m *Memory) Update(ID int, person *model.Person) error {
	if person == nil {
		return model.ErrPersonCanNotBeNil
	}
	if _, ok := m.Persons[ID]; !ok {
		return fmt.Errorf("ID: %d: %v", ID, model.ErrIDPersonDoesNotExists)
	}
	m.Persons[ID] = *person

	return nil
}

// Delete
func (m *Memory) Delete(ID int) error {
	if _, ok := m.Persons[ID]; !ok {
		return fmt.Errorf("ID: %d: %v", ID, model.ErrIDPersonDoesNotExists)
	}
	delete(m.Persons, ID)
	return nil
}

// GetByID
func (m *Memory) GetByID(ID int) (model.Person, error) {
	person, ok := m.Persons[ID]
	if !ok {
		return person, fmt.Errorf("ID: %d: %v", ID, model.ErrIDPersonDoesNotExists)
	}
	return person, nil
}

// GetAll
func (m *Memory) GetAll() (model.Persons, error) {
	var result model.Persons
	for _, v := range m.Persons {
		result = append(result, v)
	}
	return result, nil
}
