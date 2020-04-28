package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var exportLong = `export your data model to various formats`

func ExportRun(args []string) (err error) {

	return err
}

var ExportCmd = &cobra.Command{

	Use: "export",

	Short: "export your data model to various formats",

	Long: exportLong,

	Run: func(cmd *cobra.Command, args []string) {
		var err error

		// Argument Parsing

		err = ExportRun(args)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}
