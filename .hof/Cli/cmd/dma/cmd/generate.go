package cmd

import (
	"fmt"
	"os"

	"strings"

	"github.com/spf13/cobra"

	"github.com/hofstadter-io/dma/cmd/dma/ga"
)

var generateLong = `generate libraries and utilies for your data model`

func GenerateRun(args []string) (err error) {

	return err
}

var GenerateCmd = &cobra.Command{

	Use: "generate",

	Aliases: []string{
		"gen",
	},

	Short: "generate libraries and utilies for your data model",

	Long: generateLong,

	PreRun: func(cmd *cobra.Command, args []string) {

		cs := strings.Fields(cmd.CommandPath())
		c := strings.Join(cs[1:], "/")
		ga.SendGaEvent(c, strings.Join(args, "/"), 0)

	},

	Run: func(cmd *cobra.Command, args []string) {
		var err error

		// Argument Parsing

		err = GenerateRun(args)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}
