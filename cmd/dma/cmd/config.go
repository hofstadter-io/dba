package cmd

import (
	"strings"

	"github.com/spf13/cobra"

	"github.com/hofstadter-io/dma/cmd/dma/cmd/config"

	"github.com/hofstadter-io/dma/cmd/dma/ga"
)

var configLong = `view and set, global and local config values`

var ConfigCmd = &cobra.Command{

	Use: "config",

	Short: "view and set, global and local config values",

	Long: configLong,

	PreRun: func(cmd *cobra.Command, args []string) {

		cs := strings.Fields(cmd.CommandPath())
		c := strings.Join(cs[1:], "/")
		ga.SendGaEvent(c, "<omit>", 0)

	},
}

func init() {
	hf := ConfigCmd.HelpFunc()
	f := func(cmd *cobra.Command, args []string) {
		cs := strings.Fields(cmd.CommandPath())
		c := strings.Join(cs[1:], "/")
		ga.SendGaEvent(c+"/help", "<omit>", 0)
		hf(cmd, args)
	}
	ConfigCmd.SetHelpFunc(f)
	ConfigCmd.AddCommand(cmdconfig.ListCmd)
	ConfigCmd.AddCommand(cmdconfig.GetCmd)
	ConfigCmd.AddCommand(cmdconfig.SetCmd)
	ConfigCmd.AddCommand(cmdconfig.UseCmd)
}
