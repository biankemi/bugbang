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
		frontend.GET("/sendCode", controllers.SendCode)
		frontend.GET("/login", controllers.Login)
		frontend.POST("/user/edit/:id", controllers.UserEdit)
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
