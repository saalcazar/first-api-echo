package model

type Comunity struct {
	Name string
}

type Communities []Comunity

type Person struct {
	Name        string
	Age         uint8
	Communities Communities
}

type Persons []Person
