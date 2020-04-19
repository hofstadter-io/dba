package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	/*
		false
		false
		false
		false
		false
		true
	*/

	"github.com/hofstadter-io/dma/lib/cmd"
)

var importLong = `Import and create a data model from a multitude of sources such as
SQL, NoSQL, object storage, and buckets.`

var ImportCmd = &cobra.Command{

	Use: "import",

	Short: "import and create a data model from a multitude of sources",

	Long: importLong,

	Run: func(cmd *cobra.Command, args []string) {
		var err error

		// Argument Parsing

		err = libcmd.ImportRun()

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}
