package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gnadlinger/SpommunicateBackend/handler"
)

func CreateRoutes(){
	r := gin.Default()
	v1 := r.Group("api/v1")
	{
		v1.GET("/people", Handler.GetUsers)
		v1.GET("/teams", Handler.GetUsers)
		/*v1.GET("/dates", handler.GetUsers)
		v1.GET("/datetypes", handler.GetUsers)
		v1.GET("/functiontypes", handler.GetUsers)
		v1.GET("/functions", handler.GetUsers)
		v1.GET("/lineups", handler.GetUsers)
		v1.GET("/squads", handler.GetUsers)
		v1.GET("/teams", Handler.GetTeams)*/
	}

	r.Run(":8080")
}
