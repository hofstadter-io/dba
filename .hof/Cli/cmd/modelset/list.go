package cmdmodelset

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

	"github.com/hofstadter-io/dma/lib/cmd/modelset"
)

var listLong = `list the known modelsets`

var ListCmd = &cobra.Command{

	Use: "list",

	Short: "list the known modelsets",

	Long: listLong,

	Run: func(cmd *cobra.Command, args []string) {
		var err error

		// Argument Parsing

		err = libcmdmodelset.ListRun()

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}
