package cmd

import (
	"fmt"
	"os"

	"strings"

	"github.com/spf13/cobra"

	"github.com/hofstadter-io/dma/cmd/dma/ga"
)

var uiLong = `run dma's local web ui`

func UiRun(args []string) (err error) {

	return err
}

var UiCmd = &cobra.Command{

	Use: "ui",

	Short: "run dma's local web ui",

	Long: uiLong,

	PreRun: func(cmd *cobra.Command, args []string) {

		cs := strings.Fields(cmd.CommandPath())
		c := strings.Join(cs[1:], "/")
		ga.SendGaEvent(c, strings.Join(args, "/"), 0)

	},

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
