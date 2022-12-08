package cmd

import (
	"github.com/driver005/commerce/config/db/migrations"
	"github.com/spf13/cobra"
)

func NewMigrateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "migrate",
		Short: "Various migration helpers",
	}

	return cmd
}

func NewRunMigrationCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run",
		Short: "Run database migration",
		Long: `Run this command on a fresh SQL installation and when you upgrade to a new minor version.
It is recommended to run this command close to the SQL instance (e.g. same subnet) instead of over the public internet.
This decreases risk of failure and decreases time required.
You can read in the database URL using the -e flag, for example:
	export DSN=...
	hydra migrate sql -e
### WARNING ###
Before running this command on an existing database, create a back up!`,
		Run: func(cmd *cobra.Command, args []string) {
			migrations.Migrate()
		},
	}

	return cmd
}
