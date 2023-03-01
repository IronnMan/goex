package cmd

import (
	"github.com/spf13/cobra"
	"goex/database/migrations"
	"goex/pkg/migrate"
)

func init() {
	CmdMigrate.AddCommand(
		CmdMigrateUp,
		CmdMigrateRollback,
		CmdMigrateRefresh,
		CmdMigrateReset,
	)
}

var CmdMigrate = &cobra.Command{
	Use:   "migrate",
	Short: "Run database migration",
}

var CmdMigrateUp = &cobra.Command{
	Use:   "up",
	Short: "Run unmigrated migrations",
	Run:   runUp,
}

var CmdMigrateRollback = &cobra.Command{
	Use:     "down",
	Aliases: []string{"rollback"},
	Short:   "Reverse the up command",
	Run:     runDown,
}

var CmdMigrateReset = &cobra.Command{
	Use:   "reset",
	Short: "Rollback all database migrations",
	Run:   runReset,
}

var CmdMigrateRefresh = &cobra.Command{
	Use:   "refresh",
	Short: "Reset and re-run all migrations",
	Run:   runRefresh,
}

func runRefresh(cmd *cobra.Command, args []string) {
	migrator().Refresh()
}

func runReset(cmd *cobra.Command, args []string) {
	migrator().Reset()
}

func runDown(cmd *cobra.Command, args []string) {
	migrator().Rollback()
}

func runUp(cmd *cobra.Command, args []string) {
	migrator().Up()
}

func migrator() *migrate.Migrator {
	migrations.Initialize()
	return migrate.NewMigrator()
}
