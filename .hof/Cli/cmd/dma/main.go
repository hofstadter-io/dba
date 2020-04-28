package main

import (
	"fmt"
	"os"

	"github.com/hofstadter-io/dma/cmd/dma/cmd"
)

func main() {

	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
