package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	// "github.com/spf13/viper"

	"github.com/hofstadter-io/dma/cmd/dma/pflags"

	"github.com/hofstadter-io/dma/lib/runtime"
)

var dmaLong = `Your Data Model Assistant`

func init() {

	RootCmd.PersistentFlags().StringVarP(&pflags.RootConfigPflag, "config", "c", "", "The path to a dma config file")
}

func RootPersistentPreRun(args []string) (err error) {

	runtime.Init()

	return err
}

func RootPersistentPostRun(args []string) (err error) {

	WaitPrintUpdateAvail()

	return err
}

var RootCmd = &cobra.Command{

	Use: "dma",

	Short: "Your Data Model Assistant",

	Long: dmaLong,

	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		var err error

		// Argument Parsing

		err = RootPersistentPreRun(args)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},

	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		var err error

		// Argument Parsing

		err = RootPersistentPostRun(args)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.AddCommand(InitCmd)
	RootCmd.AddCommand(ConfigCmd)
	RootCmd.AddCommand(UiCmd)
	RootCmd.AddCommand(ModelsetCmd)
	RootCmd.AddCommand(StoreCmd)
	RootCmd.AddCommand(ImportCmd)
	RootCmd.AddCommand(ExportCmd)
	RootCmd.AddCommand(GenerateCmd)
}

func initConfig() {

}