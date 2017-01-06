package models

import (
	"github.com/jinzhu/gorm"
	"time"
)
type Participation struct {
	gorm.Model
	username int
	confirmation bool
	rejection string
	Person Person
	Date Date
}
type Person struct {
	ID int
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
	gorm.Model
	PersonId int
	Info string
	Person Person
	FunctionType FunctionType
	Squads []Squad
}
type Squad struct {
	gorm.Model
	Name string
	Info string
	Team Team
	Function []Function
	Dates []Date
	LineUp LineUp
}
type LineUp struct {
	gorm.Model
	Version int
	//PositionX []int
	//PositionY []int
	Image string
}
type Team struct {
	gorm.Model
	LineUps []LineUp
	Persons []Person
	Info string
}
type Date struct {
	gorm.Model
	Name string
	Place string
	Day time.Time
	Squad Squad
	DateType DateType

}
type DateType struct {
	gorm.Model
	Name string
	Squads []Squad
}
type FunctionType struct {
	gorm.Model
	Function []Function
	Name string
	Short string
}
type PositionType struct {
	gorm.Model
	Participations []Participation
	Name string
	Short string
}

