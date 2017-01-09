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
	db.AutoMigrate(&models.Participation{}, &models.DateType{}, &models.Date{}, &models.Function{},
		&models.FunctionType{},
		&models.LineUp{}, &models.Person{}, &models.Team{})

	participation:= models.Participation{Confirmation:true,Rejection:"Sorry"}
	lineup:= models.LineUp{
		Version:1,
		Image:"FOo",
	}

	team := models.Team{
		Info:"FCB",
		LineUps:[]models.LineUp{lineup},
	}
	team1 := models.Team{
		Info:"BVB",
	}
	user := models.Person{
		FirstName:"Isa",
		LastName:"istcool",
		Functions:[]models.Function{{Info: "Trainer"}, {Info: "CEO"}},
		Username:"Isarovic",
		Participations:[]models.Participation{participation},
	}

	user1 := models.Person{FirstName: "Thomas", LastName: "abc", Birthday: time.Now(),Username:"Foo"}
	user2 := models.Person{FirstName: "Hermann", LastName: "abc", Birthday: time.Now(),Username:"Sonny"}
	user3 := models.Person{FirstName: "Hey", LastName: "waslos", Birthday: time.Now(),Username:"Black"}


	/*db.Create(&team)
	db.Create(&participation)*/
	db.Create(&user)
	db.Create(&team)
	db.Create(&team1)
	db.Create(&user1)
	db.Create(&user2)
	db.Create(&user3)
	db.Create(participation)
}
