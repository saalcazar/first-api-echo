// Crea la interface que deben implementar los sitemas de almacenamiento de nuestra API
package handler

import "github.com/saalcazar/first-api-echo/model"

//Storage. que quiera trabajar con mi handler debe implementar esta interface
type Storage interface {
	Create(person *model.Person) error
	Update(ID int, person *model.Person) error
	Delete(ID int) error
	GetByID(ID int) (model.Person, error)
	GetAll() (model.Persons, error)
}
