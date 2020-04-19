package cmd

import (
	"github.com/spf13/cobra"
	// "github.com/spf13/viper"
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
}

func init() {
	RootCmd.AddCommand(InitCmd)
	RootCmd.AddCommand(ConfigCmd)
	RootCmd.AddCommand(UiCmd)
	RootCmd.AddCommand(MigrateCmd)
	RootCmd.AddCommand(StoreCmd)
	RootCmd.AddCommand(ImportCmd)
	RootCmd.AddCommand(ExportCmd)
	RootCmd.AddCommand(GenerateCmd)
}
