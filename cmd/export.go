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

var exportLong = `export your data model to various formats`

var ExportCmd = &cobra.Command{

	Use: "export",

	Short: "export your data model to various formats",

	Long: exportLong,

	Run: func(cmd *cobra.Command, args []string) {
		var err error

		// Argument Parsing

		err = libcmd.ExportRun()

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}
