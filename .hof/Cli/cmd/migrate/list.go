package cmdmigrate

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

	"github.com/hofstadter-io/dma/lib/cmd/migrate"
)

var listLong = `list the known datastores and migration briefs`

var ListCmd = &cobra.Command{

	Use: "list",

	Short: "list the known datastores and migration briefs",

	Long: listLong,

	Run: func(cmd *cobra.Command, args []string) {
		var err error

		// Argument Parsing

		err = libcmdmigrate.ListRun()

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}
