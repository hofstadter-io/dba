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

var generateLong = `generate libraries and utilies for your data model`

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

		err = libcmd.GenerateRun()

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}
