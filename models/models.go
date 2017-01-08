package models

import (
	//"github.com/jinzhu/gorm"
	"time"
)
type Participation struct {
	ID int
	username int
	confirmation bool
	rejection string
	Person Person
	Date Date
}
type Person struct {
	ID int
	TeamId int
	FirstName string
	LastName string
	Mobil string
	Email string
	Birthday time.Time
	Username string
	Password string
	PositionType PositionType
	Functions []Function
	Participations []Participation
	Team Team
}
type Function struct {
	ID int
	PersonId int
	Info string
	Person Person
	FunctionType FunctionType
	Squads []Squad
}
type Squad struct {
	ID int
	Name string
	Info string
	Team Team
	Function []Function
	Dates []Date
	LineUp LineUp
}
type LineUp struct {
	ID int
	Version int
	TeamId int
	Team Team
	//PositionX []int
	//PositionY []int
	Image string
}
type Team struct {
	ID int
	PersonId int
	LineUps []LineUp
	Persons []Person
	Info string
}
type Date struct {
	ID int
	Name string
	Place string
	Day time.Time
	Squad Squad
	DateType DateType

}
type DateType struct {
	ID int
	Name string
	Squads []Squad
}
type FunctionType struct {
	ID int
	Function []Function
	Name string
	Short string
}
type PositionType struct {
	ID int
	Participations []Participation
	Name string
	Short string
}

