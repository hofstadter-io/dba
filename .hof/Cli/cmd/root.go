package cmd

import (
	"github.com/spf13/cobra"
	// "github.com/spf13/viper"
)

var dbaLong = `Your Database Assistant`

var (
	RootConfigPflag  string
	RootContextPflag string
)

func init() {

	RootCmd.PersistentFlags().StringVarP(&RootConfigPflag, "config", "c", "", "The path to a dba config file")

	RootCmd.PersistentFlags().StringVarP(&RootContextPflag, "context", "C", "", "the context to use during this hof execution")

}

var RootCmd = &cobra.Command{

	Use: "dba",

	Short: "Your Database Assistant",

	Long: dbaLong,
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
