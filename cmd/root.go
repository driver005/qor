/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ecommerce",
		Short: "Run and manage qur ecommerce",
	}

	AddCommandRecursive(cmd)
	return cmd
}

func AddCommandRecursive(parent *cobra.Command) {
	migrateCmd := NewMigrateCmd()
	migrateCmd.AddCommand(NewRunMigrationCmd())

	serveCmd := NewServeCmd()
	serveCmd.AddCommand(NewServeAllCmd())

	parent.AddCommand(
		migrateCmd,
		serveCmd,
		NewVersionCommand(),
	)
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := NewRootCmd().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
