package cmdmodelset

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var migrateLong = `create the next migration for a modelset`

func MigrateRun(name string) (err error) {

	return err
}

var MigrateCmd = &cobra.Command{

	Use: "migrate",

	Short: "create the next migration for a modelset",

	Long: migrateLong,

	Run: func(cmd *cobra.Command, args []string) {
		var err error

		// Argument Parsing

		if 0 >= len(args) {
			fmt.Println("missing required argument: 'Name'")
			cmd.Usage()
			os.Exit(1)
		}

		var name string

		if 0 < len(args) {

			name = args[0]

		}

		err = MigrateRun(name)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}
