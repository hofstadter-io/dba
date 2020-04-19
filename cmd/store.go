package cmd

import (
	"github.com/spf13/cobra"

	"github.com/hofstadter-io/dma/cmd/store"
	/*
		false
		false
		true
		false
		false
		false
	*/)

var storeLong = `create, checkpoint, and migrate your storage engines`

var StoreCmd = &cobra.Command{

	Use: "store",

	Short: "create, checkpoint, and migrate your storage engines",

	Long: storeLong,
}

func init() {
	StoreCmd.AddCommand(cmdstore.RunCmd)
	StoreCmd.AddCommand(cmdstore.ConnCmd)
}
