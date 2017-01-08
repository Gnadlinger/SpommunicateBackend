package Handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gnadlinger/SpommunicateBackend/models"
)

func  GetUsers(c *gin.Context) {
	//db.Preload("Functions").First(&models.Person{},1)
	c.JSON(200, db.Preload("Functions").Find(&[]models.Person{}))
}