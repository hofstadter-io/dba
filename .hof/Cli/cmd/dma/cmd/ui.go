package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var uiLong = `run dma's local web ui`

func UiRun(args []string) (err error) {

	return err
}

var UiCmd = &cobra.Command{

	Use: "ui",

	Short: "run dma's local web ui",

	Long: uiLong,

	Run: func(cmd *cobra.Command, args []string) {
		var err error

		// Argument Parsing

		err = UiRun(args)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}
