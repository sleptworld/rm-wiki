package Router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/sleptworld/test/Controller/v1"
)

func EntryRouter(group *gin.RouterGroup)  {

	// All Entry
	group.GET("/Entry",v1.EntryHandler)
	group.POST("/Entry")

	eGroup := group.Group("/Entry")
	{
		// Single Entry api
		singleEntry := "/:id"
		eGroup.GET(singleEntry, v1.EntryHandler)
		eGroup.PUT(singleEntry)
		eGroup.DELETE(singleEntry)
		eGroup.PATCH(singleEntry)

		// Cat
		cGroup := eGroup.Group("/cat")
		{
			catEntry := "/:id"
			cGroup.GET(catEntry)
			cGroup.POST(catEntry)
			cGroup.PUT(catEntry)
			cGroup.DELETE(catEntry)
			cGroup.PATCH(catEntry)
		}
	}
	//
}
