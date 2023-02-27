package make

import (
	"fmt"
	"github.com/spf13/cobra"
	"goex/pkg/console"
)

var CmdMakeCMD = &cobra.Command{
	Use:   "cmd",
	Short: "Create a command, should be snake_case, example: make cmd backup_database",
	Run:   runMakeCMD,
	Args:  cobra.ExactArgs(1),
}

func runMakeCMD(cmd *cobra.Command, args []string) {
	model := makeModelFromString(args[0])

	filePath := fmt.Sprintf("app/cmd/%s.go", model.PackageName)

	createFileFromStub(filePath, "cmd", model)

	console.Success("command name:" + model.PackageName)
	console.Success("command variable name: cmd.Cmd" + model.StructName)
	console.Warning("please edit main.go's app.Commands slice to register command")
}
