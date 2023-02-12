package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goex/bootstrap"
)

func main() {

	router := gin.New()

	bootstrap.SetupRoute(router)

	err := router.Run(":3000")
	if err != nil {
		fmt.Printf(err.Error())
	}
}
