package Router

import "github.com/gin-gonic/gin"

func CommentsRouter(group gin.RouterGroup)  {
	cGroup := group.Group("/comments")
	{
		cGroup.GET("/")

		cGroup.GET("/:entryid")

		// user comments
		cGroup.GET("/user/:name")
		cGroup.POST("/user/:name")
		cGroup.PATCH("/user/:name")
		cGroup.DELETE("/user/:name")
	}
}
