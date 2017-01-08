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
	c.JSON(200, db.Find([]models.LineUp{}))
}
func  GetParticipations(c *gin.Context) {
	c.JSON(200, db.Find([]models.Participation{}))
}
func  GetSquads(c *gin.Context) {
	c.JSON(200, db.Find([]models.Squad{}))
}


