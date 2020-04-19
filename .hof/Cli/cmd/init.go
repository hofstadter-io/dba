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

var initLong = `init the current directory for dma usage.`

var InitCmd = &cobra.Command{

	Use: "init",

	Short: "init the current directory for dma usage.",

	Long: initLong,

	Run: func(cmd *cobra.Command, args []string) {
		var err error

		// Argument Parsing

		if 0 >= len(args) {
			fmt.Println("missing required argument: 'Dstype'")
			cmd.Usage()
			os.Exit(1)
		}

		var dstype string

		if 0 < len(args) {

			dstype = args[0]

		}

		if 1 >= len(args) {
			fmt.Println("missing required argument: 'Name'")
			cmd.Usage()
			os.Exit(1)
		}

		var name string

		if 1 < len(args) {

			name = args[1]

		}

		err = libcmd.InitRun(dstype, name)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}
