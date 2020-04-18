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

	"github.com/hofstadter-io/dba/lib/cmd"
)

var storeLong = `create, checkpoint, and migrate your storage engines`

var StoreCmd = &cobra.Command{

	Use: "store",

	Short: "create, checkpoint, and migrate your storage engines",

	Long: storeLong,

	Run: func(cmd *cobra.Command, args []string) {
		var err error

		// Argument Parsing

		err = libcmd.StoreRun()

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}
