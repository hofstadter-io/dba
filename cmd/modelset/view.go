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

var viewLong = `view modelset information`

var ViewCmd = &cobra.Command{

	Use: "view",

	Short: "view modelset information",

	Long: viewLong,

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

		err = libcmdmodelset.ViewRun(name)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}
