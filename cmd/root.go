package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	// "github.com/spf13/viper"

	"github.com/hofstadter-io/dma/lib/cmd"
)

var dmaLong = `Your Data Model Assistant`

var (
	RootConfigPflag  string
	RootContextPflag string
)

func init() {

	RootCmd.PersistentFlags().StringVarP(&RootConfigPflag, "config", "c", "", "The path to a dma config file")

	RootCmd.PersistentFlags().StringVarP(&RootContextPflag, "context", "C", "", "the context to use during this hof execution")

}

var RootCmd = &cobra.Command{

	Use: "dma",

	Short: "Your Data Model Assistant",

	Long: dmaLong,

	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		var err error

		// Argument Parsing

		err = libcmd.RootPersistentPreRun()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	RootCmd.AddCommand(InitCmd)
	RootCmd.AddCommand(ConfigCmd)
	RootCmd.AddCommand(UiCmd)
	RootCmd.AddCommand(ModelsetCmd)
	RootCmd.AddCommand(StoreCmd)
	RootCmd.AddCommand(ImportCmd)
	RootCmd.AddCommand(ExportCmd)
	RootCmd.AddCommand(GenerateCmd)
}