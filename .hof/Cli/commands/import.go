package commands

import (

	// hello... something might need to go here

	"github.com/spf13/cobra"
)

var importLong = `Import and create a data model from a multitude of sources such as
SQL, NoSQL, object storage, and buckets.`

var ImportCmd = &cobra.Command{

	Use: "import",

	Short: "import and create a data model from a multitude of sources",

	Long: importLong,

	Run: func(cmd *cobra.Command, args []string) {

		// Argument Parsing

		// Default body

		fmt.Println("dba import")

	},
}
