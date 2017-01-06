package Handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gnadlinger/Spommunicate1/models"
)

func  GetUsers(c *gin.Context) {
	c.JSON(200, db.Find(&[]models.Person{}))
}