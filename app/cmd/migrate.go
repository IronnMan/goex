package cmd

import (
	"github.com/spf13/cobra"
	"goex/database/migrations"
	"goex/pkg/migrate"
)

var CmdMigrate = &cobra.Command{
	Use:   "migrate",
	Short: "Run database migration",
}

var CmdMigrateUp = &cobra.Command{
	Use:   "up",
	Short: "Run unmigrated migrations",
	Run:   runUp,
}

func runUp(cmd *cobra.Command, args []string) {
	migrator().Up()
}

func migrator() *migrate.Migrator {
	migrations.Initialize()
	return migrate.NewMigrator()
}
