package cmdmodelset

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var testLong = `test the current migration and diff for a modelset`

func TestRun(name string) (err error) {

	return err
}

var TestCmd = &cobra.Command{

	Use: "test",

	Short: "test the current migration and diff for a modelset",

	Long: testLong,

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

		err = TestRun(name)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}
