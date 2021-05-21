package Router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/sleptworld/test/Controller/v1"
)

func EntryRouter(group *gin.RouterGroup)  {

	// All Entry
	group.GET("/Entry",v1.GETEntry)
	group.POST("/Entry",v1.POSTEntry)

	eGroup := group.Group("/Entry")
	{
		// Single Entry api
		singleEntry := "/:id"
		eGroup.GET(singleEntry,v1.GETEntryByID)

		eGroup.PUT(singleEntry,v1.PUTEntryByID)
		eGroup.DELETE(singleEntry,v1.DELETEEntryByID)
		eGroup.PATCH(singleEntry,v1.PATCHEntryByID)
	}

	// Cat
	eGroup.GET("/cat",v1.GETCat)
	eGroup.POST("/cat",v1.POSTCat)

	cGroup := eGroup.Group("/cat",v1.GETCat)
	{
		singleCat := "/:id"
		cGroup.GET(singleCat,v1.GETCat)
		cGroup.DELETE(singleCat,v1.GETCat)
	}
}
