package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"goex/bootstrap"
	btsConfig "goex/config"
	"goex/pkg/config"
)

func init() {
	// Load the configuration information in the config directory
	btsConfig.Initialize()
}

func main() {

	// Configuration initialization, relying on the command line --env parameter
	var env string
	flag.StringVar(&env, "env", "", "Load the .env file, such as --env=testing loads the .env.testing file")
	flag.Parse()
	config.InitConfig(env)

	// Initialize Logger
	bootstrap.SetupLogger()

	router := gin.New()

	// Initialize DB
	bootstrap.SetupDB()

	// Initialize route binding
	bootstrap.SetupRoute(router)

	err := router.Run(":" + config.Get("app.port"))
	if err != nil {
		fmt.Printf(err.Error())
	}
}
