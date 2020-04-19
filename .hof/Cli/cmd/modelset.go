package cmd

import (
	"github.com/spf13/cobra"

	"github.com/hofstadter-io/dma/cmd/modelset"
	/*
		false
		false
		true
		false
		false
		false
	*/)

var modelsetLong = `create, view, migrate, and understand your data models.`

var ModelsetCmd = &cobra.Command{

	Use: "modelset",

	Aliases: []string{
		"model",
		"m",
	},

	Short: "create, view, migrate, and understand your data models.",

	Long: modelsetLong,
}

func init() {
	ModelsetCmd.AddCommand(cmdmodelset.CreateCmd)
	ModelsetCmd.AddCommand(cmdmodelset.ViewCmd)
	ModelsetCmd.AddCommand(cmdmodelset.ListCmd)
	ModelsetCmd.AddCommand(cmdmodelset.StatusCmd)
	ModelsetCmd.AddCommand(cmdmodelset.GraphCmd)
	ModelsetCmd.AddCommand(cmdmodelset.DiffCmd)
	ModelsetCmd.AddCommand(cmdmodelset.MigrateCmd)
	ModelsetCmd.AddCommand(cmdmodelset.TestCmd)
	ModelsetCmd.AddCommand(cmdmodelset.DeleteCmd)
}
