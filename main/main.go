package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	//"github.com/gnadlinger/Spommunicate1/handler"
	"github.com/gnadlinger/Spommunicate1/routes"
	"github.com/gnadlinger/Spommunicate1/models"
	"github.com/jinzhu/gorm"
	"time"
)
var db, err = gorm.Open("mysql", "Gnadlinger:admin@/spommunicatedb?charset=utf8&parseTime=True&loc=Local")
func main() {
	defer db.Close()
//foo
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	//db.Exec("PRAGMA foreign_keys = ON;")
	db.DropTable(&models.Person{})
	db.AutoMigrate(&models.Participation{}, &models.DateType{}, &models.Date{}, &models.Function{}, &models.FunctionType{},
		&models.LineUp{}, &models.Person{}, &models.Squad{}, &models.Team{})
	//function:=[]models.Function{{Info:"Trainer"}}
	user := models.Person{
		FirstName:            "Isa",
		LastName:  "istcool",
		Functions:          []models.Function{{Info: "Trainer"}, {Info: "CEO"}},
	}
	//foo
	user1 := models.Person{FirstName: "Thomas", LastName: "abc", Birthday: time.Now()}
	user2 := models.Person{FirstName: "Johannes", LastName: "Gnadlinger"}
	user3 := models.Person{FirstName: "Hey", LastName: "waslos", Birthday: time.Now()}
	db.Create(&user)
	db.Create(&user1)
	db.Create(&user2)
	db.Create(&user3)
	routes.CreateRoutes()

	var u models.Person
	db.Debug().First(&u).Related(&u.Functions)
}
