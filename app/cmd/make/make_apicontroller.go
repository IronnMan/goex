package make

import (
	"fmt"
	"github.com/spf13/cobra"
	"goex/pkg/console"
	"strings"
)

var CmdMakeAPIController = &cobra.Command{
	Use:   "apicontroller",
	Short: "Create api controller, example: make apicontroller v1/user",
	Run:   runMakeAPIController,
	Args:  cobra.ExactArgs(1),
}

func runMakeAPIController(cmd *cobra.Command, args []string) {

	array := strings.Split(args[0], "/")
	if len(array) != 2 {
		console.Exit("api controller name format: v1/user")
	}

	apiVersion, name := array[0], array[1]
	model := makeModelFromString(name)

	filePath := fmt.Sprintf("app/http/controllers/api/%s/%s_controller.go", apiVersion, model.TableName)

	createFileFromStub(filePath, "apicontroller", model)
}
