package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var initLong = `init the current directory for dma usage.`

func InitRun(args []string) (err error) {

	return err
}

var InitCmd = &cobra.Command{

	Use: "init",

	Short: "init the current directory for dma usage.",

	Long: initLong,

	Run: func(cmd *cobra.Command, args []string) {
		var err error

		// Argument Parsing

		err = InitRun(args)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}
