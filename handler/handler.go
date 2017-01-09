package Handler
import (
	"github.com/gnadlinger/SpommunicateBackend/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)
var db, err = gorm.Open("mysql", "Gnadlinger:admin@/spommunicatedb?charset=utf8&parseTime=True&loc=Local")
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
	c.JSON(200,db.
		Find(&[]models.LineUp{}).
		Where("info=?",c.Param("team")))
}
func  GetParticipations(c *gin.Context) {
	//var i = c.Param("username")
	c.JSON(200, db.Preload("Person").Find(&[]models.Participation{}).
		Where("username=?",c.Param("username")))
}
func GetUser(c *gin.Context){
	c.JSON(200, db.Where("username = ?", c.Param("username")).Find(&models.Person{}))
}
func  AddTeamToPlayer(c *gin.Context) {
	username := c.PostForm("username")
	team := c.PostForm("team")
	db.Model(&models.Person{}).Where("username=?",username).
			Update("team",db.Find(&models.Person{}).Where("info=?",team))
}
func  PostPlayerToTeam(c *gin.Context) {
	username := c.PostForm("username")
	team := c.PostForm("team")
	db.Model(&models.Team{}).Where("info=?",team).
	UpdateColumn("Persons",db.Find(&models.Person{}).Where("username=?",username))
}


