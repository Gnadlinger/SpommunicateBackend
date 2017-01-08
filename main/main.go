package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/gnadlinger/SpommunicateBackend/routes"
	"time"
	"github.com/jinzhu/gorm"
	"github.com/gnadlinger/SpommunicateBackend/models"
)
var db, err = gorm.Open("mysql", "Gnadlinger:admin@/spommunicatedb?charset=utf8&parseTime=True&loc=Local")
func main() {
	initDb()
	routes.CreateRoutes()
}
func initDb(){
	defer db.Close()

	if err != nil {
		panic("failed to connect database")
	}

	db.DropTable(&models.Person{},&models.Team{})
	// Migrate the schema
	db.AutoMigrate(&models.Participation{}, &models.DateType{}, &models.Date{}, &models.Function{}, &models.FunctionType{},
		&models.LineUp{}, &models.Person{}, &models.Squad{}, &models.Team{})
	user := models.Person{
		FirstName:"Isa",
		LastName:"istcool",
		Functions:[]models.Function{{Info: "Trainer"}, {Info: "CEO"}},
	}

	user1 := models.Person{FirstName: "Thomas", LastName: "abc", Birthday: time.Now()}
	user2 := models.Person{FirstName: "Thomas", LastName: "abc", Birthday: time.Now()}
	user3 := models.Person{FirstName: "Hey", LastName: "waslos", Birthday: time.Now()}

	lineup:= models.LineUp{
		Version:1,
		Image:"FOo",
	}
	team := models.Team{
		Info:"FCB",
		Persons: []models.Person{user,user1},
		LineUps:[]models.LineUp{lineup},
	}
	db.Create(&team)
	db.Create(&user)
	db.Create(&user1)
	db.Create(&user2)
	db.Create(&user3)
}
