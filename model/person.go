package model

type Comunity struct {
	Name string `json:"name"`
}

type Communities []Comunity

type Person struct {
	Name        string      `json:"name"`
	Age         uint8       `json:"age"`
	Communities Communities `json:"communities"`
}

type Persons []Person
