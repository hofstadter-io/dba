package cmdconfig

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

	"github.com/hofstadter-io/dma/lib/cmd/config"
)

var getLong = `print a configuration`

var GetCmd = &cobra.Command{

	Use: "get",

	Short: "print a configuration",

	Long: getLong,

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

		err = libcmdconfig.GetRun(name)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}
