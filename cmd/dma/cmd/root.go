package cmd

import (
	"fmt"
	"os"

	"strings"

	"github.com/spf13/cobra"

	// "github.com/spf13/viper"

	"github.com/hofstadter-io/dma/cmd/dma/ga"

	"github.com/hofstadter-io/dma/cmd/dma/pflags"

	"github.com/hofstadter-io/dma/lib/runtime"
)

var dmaLong = `Your Data Model Assistant`

func init() {

	RootCmd.PersistentFlags().StringVarP(&pflags.RootConfigPflag, "config", "c", "", "The path to a dma config file")
	RootCmd.PersistentFlags().StringVarP(&pflags.RootCredsPflag, "creds", "C", "", "The path to a dma creds file")
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

	PreRun: func(cmd *cobra.Command, args []string) {

		ga.SendGaEvent("root", strings.Join(args, "/"), 0)
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

	hf := RootCmd.HelpFunc()
	f := func(cmd *cobra.Command, args []string) {
		if RootCmd.Name() == cmd.Name() {
			as := strings.Join(args, "/")
			ga.SendGaEvent("root/help", as, 0)
		}
		hf(cmd, args)
	}
	RootCmd.SetHelpFunc(f)

	cobra.OnInitialize(initConfig)
	RootCmd.AddCommand(InitCmd)
	RootCmd.AddCommand(ConfigCmd)
	RootCmd.AddCommand(UiCmd)
	RootCmd.AddCommand(ModelsetCmd)
	RootCmd.AddCommand(StoreCmd)
	RootCmd.AddCommand(ImportCmd)
	RootCmd.AddCommand(ExportCmd)
	RootCmd.AddCommand(GenerateCmd)
	RootCmd.AddCommand(HackCmd)
}

func initConfig() {

}
