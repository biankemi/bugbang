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
			userRelation.GET("/list", controllers.UserRelationList)
			userRelation.GET("/create", controllers.UserRelationCreate)
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
