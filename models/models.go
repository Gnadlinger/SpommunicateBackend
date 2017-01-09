package models

import (
	//"github.com/jinzhu/gorm"
	"time"
)
type Participation struct {
	ID int
	Username int
	Confirmation bool
	Rejection string
	Person Person
	Date Date
}
type Person struct {
	ID int
	TeamId int
	ParticipationId int
	FunctionId int
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
}
type Squad struct {
	ID int
	Name string
	Info string
	Person []Person
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
	LineUpId int
	LineUps []LineUp
	Persons []Person
	Info string
}
type Date struct {
	ID int
	Name string
	Place string
	Day time.Time
	DateType DateType

}
type DateType struct {
	ID int
	Name string
}
type FunctionType struct {
	ID int
	Function []Function
	Name string
	Short string
}
type PositionType struct {
	ID int
	Name string
	Short string
}



