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

var listLong = `list configurations`

var ListCmd = &cobra.Command{

	Use: "list",

	Short: "list configurations",

	Long: listLong,

	Run: func(cmd *cobra.Command, args []string) {
		var err error

		// Argument Parsing

		err = libcmdconfig.ListRun()

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}
