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

var uiLong = `run dma's local web ui`

var UiCmd = &cobra.Command{

	Use: "ui",

	Short: "run dma's local web ui",

	Long: uiLong,

	Run: func(cmd *cobra.Command, args []string) {
		var err error

		// Argument Parsing

		err = libcmd.UiRun()

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}
