package Handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gnadlinger/SpommunicateBackend/models"
	"container/list"
)
type TeamDto struct{
	Persons []models.Person
}
func  GetTeams(c *gin.Context) {
	c.JSON(200, db.Preload("LineUps").Preload("Persons").Find(&[]models.Team{}))
	//c.JSON(200, db.Preload("Functions").Find(&[]models.Person{}))

}
func GetTeamMembers(c *gin.Context){
	c.JSON(200, db.Preload("People").Table("teams").
		Select("people").
		Where("info=?",c.Params.ByName("id")))

}
func GetMembersByPosition(c *gin.Context){
	var team TeamDto
	persons:=list.New()

	db.Preload("People").Table("teams").Select("people").
		Where("info=?",c.Params.ByName("id")).Scan(&team)
	for _,element:= range team.Persons {
		if element.PositionType.Name == c.Params.ByName("position") {
			persons.PushFront(element)
		}
	}
	c.JSON(200,persons)
}
