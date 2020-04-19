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

var statusLong = `show the current status for a datastore`

var StatusCmd = &cobra.Command{

	Use: "status",

	Short: "show the current status for a datastore",

	Long: statusLong,

	Run: func(cmd *cobra.Command, args []string) {
		var err error

		// Argument Parsing

		if 0 >= len(args) {
			fmt.Println("missing required argument: 'Name'")
			cmd.Usage()
			os.Exit(1)
		}

		var name string

		if 0 < len(args) {

			name = args[0]

		}

		err = libcmdmigrate.StatusRun(name)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}
