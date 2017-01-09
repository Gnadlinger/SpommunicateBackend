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
		v1.GET("/user/:username", Handler.GetUser)
		v1.GET("/teams", Handler.GetTeams)
		v1.POST("/addTeamToPlayer",Handler.AddTeamToPlayer)
		/*v1.GET("/members/:team", Handler.GetTeamMembers)
		v1.GET("/lineups/:team", Handler.GetLineUps)
		v1.GET("/participations/:username", Handler.GetParticipations)

		v1.GET("/addteamtoplayer", Handler.AddTeamToPlayer)
		/*v1.GET("/functions", handler.GetUsers)
		v1.GET("/lineups", handler.GetUsers)
		v1.GET("/squads", handler.GetUsers)
		v1.GET("/teams", Handler.GetTeams)*/
	}

	r.Run(":8080")
}
