package Handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gnadlinger/SpommunicateBackend/models"
	"container/list"
)
type Team struct{
	Persons []models.Person
}
func  GetTeams(c *gin.Context) {
	c.JSON(200, db.Find([]models.Team{}))
}
func GetTeamMembers(c *gin.Context){
	c.JSON(200, db.Table("teams").
		Select("people").
		Where("info=?",c.Params.ByName("id")))

}
func GetMembersByPosition(c *gin.Context){
	db.Table("teams").Select("people").Where("positiontype=?",c.Params.ByName("position"))
	db.Table("teams").
		Select("people").
		Where("info=?",c.Params.ByName("id"))
	var team Team
	persons:=list.New()

	db.Table("teams").Select("people").Where("info=?",c.Params.ByName("id")).Scan(&team)
	for _,element:= range team.Persons {
		if element.PositionType.Name == c.Params.ByName("position") {
			persons.PushFront(element)
		}
	}
	c.JSON(200,persons)
}
