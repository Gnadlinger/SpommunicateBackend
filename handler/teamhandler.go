package Handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gnadlinger/SpommunicateBackend/models"
)
type TeamDto struct{
	Persons []models.Person
}
func  GetTeams(c *gin.Context) {
	c.JSON(200, db.Preload("LineUps").Preload("Persons").Find(&[]models.Team{}))
	//c.JSON(200, db.Preload("Functions").Find(&[]models.Person{}))

}
func GetTeamMembers(c *gin.Context){
	var i = c.Param("team")
	c.JSON(200, db.Preload("LineUps").Preload("Persons").Find(&[]models.Team{}).
		Select("Persons").Where("info=?",i))

	/*c.JSON(200, db.Preload("Persons").Table("teams").
		Select("Persons").
		Where("info=?",i).Find(models.Person{}))*/

}
/*func GetMembersByPosition(c *gin.Context){
	var team TeamDto
	persons:=list.New()
	db.Preload("Persons").Find(&[]models.Team{}).
		Select("Persons").Where("info=?",c.Param("team"))

	db.Preload("People").Table("teams").Select("people").
		Where("info=?",c.Params.ByName("id")).Scan(&team)
	for _,element:= range team.Persons {
		if element.PositionType.Name == c.Params.ByName("position") {
			persons.PushFront(element)
		}
	}
	c.JSON(200,persons)
}*/
