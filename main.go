package main

import (
	"github.com/biankemi/bugbang/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := routers.SetupRouter()
	gin.SetMode(gin.ReleaseMode)
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8888")
}
