package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
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
