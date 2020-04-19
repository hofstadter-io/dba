package cmd

import (
	"github.com/spf13/cobra"

	"github.com/hofstadter-io/dma/cmd/migrate"
	/*
		false
		false
		true
		false
		false
		false
	*/)

var migrateLong = `view, prepare, and make migrations to your data model.`

var MigrateCmd = &cobra.Command{

	Use: "migrate",

	Short: "view, prepare, and make migrations to your data model.",

	Long: migrateLong,
}

func init() {
	MigrateCmd.AddCommand(cmdmigrate.CreateCmd)
	MigrateCmd.AddCommand(cmdmigrate.ViewCmd)
	MigrateCmd.AddCommand(cmdmigrate.ListCmd)
	MigrateCmd.AddCommand(cmdmigrate.StatusCmd)
	MigrateCmd.AddCommand(cmdmigrate.DiffCmd)
	MigrateCmd.AddCommand(cmdmigrate.TestCmd)
	MigrateCmd.AddCommand(cmdmigrate.DeleteCmd)
}
