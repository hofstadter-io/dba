package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var importLong = `Import and create a data model from a multitude of sources such as
SQL, NoSQL, object storage, and buckets.`

func ImportRun(args []string) (err error) {

	return err
}

var ImportCmd = &cobra.Command{

	Use: "import",

	Short: "import and create a data model from a multitude of sources",

	Long: importLong,

	Run: func(cmd *cobra.Command, args []string) {
		var err error

		// Argument Parsing

		err = ImportRun(args)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}
