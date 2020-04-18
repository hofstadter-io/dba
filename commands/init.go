package commands

import (

	// hello... something might need to go here

	"github.com/spf13/cobra"
)

var initLong = `init the current directory for dba usage.`

var InitCmd = &cobra.Command{

	Use: "init",

	Short: "init the current directory for dba usage.",

	Long: initLong,

	Run: func(cmd *cobra.Command, args []string) {

		// Argument Parsing

		// Default body

		fmt.Println("dba init")

	},
}
