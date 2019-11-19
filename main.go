package main

import (
	"github.com/biankemi/bugbang/router"
)

func main() {
	r := router.SetupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
