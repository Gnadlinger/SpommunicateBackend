package Handler
import (
	"github.com/gnadlinger/Spommunicate1/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"time"
)
var db, err = gorm.Open("mysql", "Gnadlinger:admin@/spommunicatedb?charset=utf8&parseTime=True&loc=Local")
func InitDb() {
	defer db.Close()

	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.DropTable(&models.Person{})
	// Migrate the schema
	db.AutoMigrate(&models.Participation{}, &models.DateType{}, &models.Date{}, &models.Function{}, &models.FunctionType{},
		&models.LineUp{}, &models.Person{}, &models.Squad{}, &models.Team{})
	function:=[]models.Function{{Info:"Trainer"},{Info:"Platzwart"}}
	user := models.Person{FirstName: "Jinzhu", LastName: "Foo", Birthday: time.Now()}
	user1 := models.Person{FirstName: "Thomas", LastName: "abc", Birthday: time.Now()}
	user2 := models.Person{FirstName: "Thomas", LastName: "abc", Birthday: time.Now(),Functions:function}
	user3 := models.Person{FirstName: "Hey", LastName: "waslos", Birthday: time.Now()}
	db.Create(&user)
	db.Create(&user1)
	db.Create(&user2)
	db.Create(&user3)
}
func  GetDateTypes(c *gin.Context) {
	c.JSON(200, db.Find([]models.DateType{}))
}
func  GetDates(c *gin.Context) {
	c.JSON(200, db.Find([]models.Date{}))
}
func  GetFunctionTypes(c *gin.Context) {
	c.JSON(200, db.Find([]models.FunctionType{}))
}
func  GetFunctions(c *gin.Context) {
	c.JSON(200, db.Find([]models.Function{}))
}
func  GetLineUps(c *gin.Context) {
	c.JSON(200, db.Find([]models.LineUp{}))
}
func  GetParticipations(c *gin.Context) {
	c.JSON(200, db.Find([]models.Participation{}))
}
func  GetSquads(c *gin.Context) {
	c.JSON(200, db.Find([]models.Squad{}))
}


