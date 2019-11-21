package routers

import (
	"net/http"

	"github.com/biankemi/bugbang/web/controllers"
	"github.com/gin-gonic/gin"
)

// SetupRouter sss
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// 前台
	frontend := router.Group("/v1")
	{
		user := frontend.Group("/user")
		{
			user.GET("/sendCode", controllers.SendCode)
			user.GET("/login", controllers.Login)
			user.POST("/edit/:id", controllers.UserEdit)
		}

		relation := frontend.Group("/relation")
		{
			relation.GET("/list", controllers.RelationList)
		}

		userRelation := frontend.Group("/user-relation")
		{
			userRelation.GET("/list/:user_id", controllers.UserRelationList)
			userRelation.POST("/create", controllers.UserRelationCreate)
			userRelation.POST("/edit", controllers.UserRelationEdit)
			userRelation.POST("/delete", controllers.UserRelationDel)
		}

		broadcast := frontend.Group("/post")
		{
			broadcast.GET("/list", controllers.BroadcastList)
			broadcast.POST("/create", controllers.BroadcastCreate)
			broadcast.POST("/update/:id", controllers.BroadcastEdit)
			broadcast.POST("/like", controllers.DoLike)
			broadcast.POST("/comment", controllers.DoComment)
			broadcast.POST("/reward", controllers.DoReward)
		}

	}

	router.GET("/someJSON", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "GO语言",
			"tag":  "<br>",
		}

		// 输出 : {"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
		c.AsciiJSON(http.StatusOK, data)
	})

	return router
}
