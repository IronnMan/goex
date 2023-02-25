package main

import (
	"flag"
	"fmt"
	"github.com/spf13/cobra"
	"goex/app/cmd"
	"goex/bootstrap"
	btsConfig "goex/config"
	"goex/pkg/config"
	"goex/pkg/console"
	"os"
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

	var rootCmd = &cobra.Command{
		Use:   config.GetString("app.name"),
		Short: "A social news website and forum project",
		Long:  `Default will run "serve" command, you can use "-h" flag to see all subcommands`,

		PersistentPreRun: func(command *cobra.Command, args []string) {

			config.InitConfig(cmd.Env)

			// Initialize Logger
			bootstrap.SetupLogger()

			// Initialize DB
			bootstrap.SetupDB()

			// Initialize Redis
			bootstrap.SetupRedis()
		},
	}

	rootCmd.AddCommand(
		cmd.CmdServe,
		cmd.CmdKey,
	)

	cmd.RegisterDefaultCmd(rootCmd, cmd.CmdServe)

	cmd.RegisterGlobalFlags(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		console.Exit(fmt.Sprintf("Failed to run app with %v: %s", os.Args, err.Error()))
	}
}
