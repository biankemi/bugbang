package main

import (
	"github.com/biankemi/bugbang/routers"
)

func main() {
	r := routers.SetupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8888")
}
