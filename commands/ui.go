package commands

import (

	// hello... something might need to go here

	"github.com/spf13/cobra"
)

var uiLong = `run dba's local web ui`

var UiCmd = &cobra.Command{

	Use: "ui",

	Short: "run dba's local web ui",

	Long: uiLong,

	Run: func(cmd *cobra.Command, args []string) {

		// Argument Parsing

		// Default body

		fmt.Println("dba ui")

	},
}
