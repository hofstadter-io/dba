package commands

import (

	// hello... something might need to go here

	"github.com/spf13/cobra"
)

var migrateLong = `view, prepare, and make migrations to your data model.`

var MigrateCmd = &cobra.Command{

	Use: "migrate",

	Short: "view, prepare, and make migrations to your data model.",

	Long: migrateLong,

	Run: func(cmd *cobra.Command, args []string) {

		// Argument Parsing

		// Default body

		fmt.Println("dba migrate")

	},
}
