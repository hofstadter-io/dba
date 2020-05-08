package cmd

import (
	"strings"

	"github.com/spf13/cobra"

	"github.com/hofstadter-io/dma/cmd/dma/cmd/store"

	"github.com/hofstadter-io/dma/cmd/dma/ga"
)

var storeLong = `create, checkpoint, and migrate your storage engines`

var StoreCmd = &cobra.Command{

	Use: "store",

	Aliases: []string{
		"s",
	},

	Short: "create, checkpoint, and migrate your storage engines",

	Long: storeLong,

	PreRun: func(cmd *cobra.Command, args []string) {

		cs := strings.Fields(cmd.CommandPath())
		c := strings.Join(cs[1:], "/")
		ga.SendGaEvent(c, strings.Join(args, "/"), 0)

	},
}

func init() {
	hf := StoreCmd.HelpFunc()
	f := func(cmd *cobra.Command, args []string) {
		cs := strings.Fields(cmd.CommandPath())
		c := strings.Join(cs[1:], "/")
		as := strings.Join(args, "/")
		ga.SendGaEvent(c+"/help", as, 0)
		hf(cmd, args)
	}
	StoreCmd.SetHelpFunc(f)
	StoreCmd.AddCommand(cmdstore.RunCmd)
	StoreCmd.AddCommand(cmdstore.ConnCmd)
}
