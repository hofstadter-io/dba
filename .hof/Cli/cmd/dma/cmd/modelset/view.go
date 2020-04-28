package cmdmodelset

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var viewLong = `view modelset information`

func ViewRun(name string) (err error) {

	return err
}

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

		err = ViewRun(name)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}
