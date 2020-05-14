package cmdmodelset

import (
	"fmt"
	"os"

	"strings"

	"github.com/spf13/cobra"

	"github.com/hofstadter-io/dma/cmd/dma/ga"
)

var deleteLong = `delete a modelset permentantly`

func DeleteRun(name string) (err error) {

	return err
}

var DeleteCmd = &cobra.Command{

	Use: "delete",

	Short: "delete a modelset permentantly",

	Long: deleteLong,

	PreRun: func(cmd *cobra.Command, args []string) {

		cs := strings.Fields(cmd.CommandPath())
		c := strings.Join(cs[1:], "/")
		ga.SendGaEvent(c, "<omit>", 0)

	},

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

		err = DeleteRun(name)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}
