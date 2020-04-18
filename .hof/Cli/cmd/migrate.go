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

var migrateLong = `view, prepare, and make migrations to your data model.`

var MigrateCmd = &cobra.Command{

	Use: "migrate",

	Short: "view, prepare, and make migrations to your data model.",

	Long: migrateLong,

	Run: func(cmd *cobra.Command, args []string) {
		var err error

		// Argument Parsing

		err = libcmd.MigrateRun()

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}
