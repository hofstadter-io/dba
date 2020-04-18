package commands

import (

	// hello... something might need to go here

	"github.com/spf13/cobra"
)

var configLong = `view and set, global and local config values`

var ConfigCmd = &cobra.Command{

	Use: "config",

	Short: "view and set, global and local config values",

	Long: configLong,

	Run: func(cmd *cobra.Command, args []string) {

		// Argument Parsing

		// Default body

		fmt.Println("dba config")

	},
}
