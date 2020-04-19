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

		err = libcmd.InitRun()

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}
