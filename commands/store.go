package commands

import (

	// hello... something might need to go here

	"github.com/spf13/cobra"
)

var storeLong = `create, checkpoint, and migrate your storage engines`

var StoreCmd = &cobra.Command{

	Use: "store",

	Short: "create, checkpoint, and migrate your storage engines",

	Long: storeLong,

	Run: func(cmd *cobra.Command, args []string) {

		// Argument Parsing

		// Default body

		fmt.Println("dba store")

	},
}
