package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/hofstadter-io/dma/cmd/dma/ga"
)

var hackLong = `development command`

func HackRun(args []string) (err error) {

	return err
}

var HackCmd = &cobra.Command{

	Use: "hack ...",

	Short: "development command",

	Long: hackLong,

	PreRun: func(cmd *cobra.Command, args []string) {

		cs := strings.Fields(cmd.CommandPath())
		c := strings.Join(cs[1:], "/")
		ga.SendGaEvent(c, "<omit>", 0)

	},

	Run: func(cmd *cobra.Command, args []string) {
		var err error

		// Argument ParsingG
		err = HackRun(args)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}