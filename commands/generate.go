package commands

import (

	// hello... something might need to go here

	"github.com/spf13/cobra"
)

var generateLong = `generate libraries and utilies for your data model`

var GenerateCmd = &cobra.Command{

	Use: "generate",

	Aliases: []string{
		"gen",
	},

	Short: "generate libraries and utilies for your data model",

	Long: generateLong,

	Run: func(cmd *cobra.Command, args []string) {

		// Argument Parsing

		// Default body

		fmt.Println("dba generate")

	},
}
