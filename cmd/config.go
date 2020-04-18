package cmd

import (
	"github.com/spf13/cobra"

	"github.com/hofstadter-io/dba/cmd/config"
	/*
		false
		false
		true
		false
		false
		false
	*/)

var configLong = `view and set, global and local config values`

var ConfigCmd = &cobra.Command{

	Use: "config",

	Short: "view and set, global and local config values",

	Long: configLong,
}

func init() {
	ConfigCmd.AddCommand(cmdconfig.ListCmd)
	ConfigCmd.AddCommand(cmdconfig.GetCmd)
	ConfigCmd.AddCommand(cmdconfig.SetCmd)
	ConfigCmd.AddCommand(cmdconfig.UseCmd)
}
