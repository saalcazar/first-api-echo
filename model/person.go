package model

type Comunity struct {
	Name string
}

type Comunities []Comunity

type Person struct {
	Name       string
	Age        uint8
	Comunities Comunities
}

type Persons []Person
