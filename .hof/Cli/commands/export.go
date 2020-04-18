package commands

import (

	// hello... something might need to go here

	"github.com/spf13/cobra"
)

var exportLong = `export your data model to various formats`

var ExportCmd = &cobra.Command{

	Use: "export",

	Short: "export your data model to various formats",

	Long: exportLong,

	Run: func(cmd *cobra.Command, args []string) {

		// Argument Parsing

		// Default body

		fmt.Println("dba export")

	},
}
