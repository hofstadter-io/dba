package cmdmodelset

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var createLong = `create a modelset`

func CreateRun(name string, entrypoint string) (err error) {

	return err
}

var CreateCmd = &cobra.Command{

	Use: "create",

	Short: "create a modelset",

	Long: createLong,

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

		var entrypoint string
		entrypoint = "models"

		if 1 < len(args) {

			entrypoint = args[1]

		}

		err = CreateRun(name, entrypoint)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}
