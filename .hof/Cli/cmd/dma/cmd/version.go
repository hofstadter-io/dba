package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/spf13/cobra"
)

var (
	Version = "Local"
	Commit  = "Dirty"

	BuildDate = "Unknown"
	GoVersion = "Unknown"
	BuildOS   = "Unknown"
	BuildArch = "Unknown"
)

func init() {

	if BuildDate == "Unknown" {
		BuildDate = time.Now().String()
		GoVersion = "run 'go version', you should have been the one who built this"
		BuildOS = runtime.GOOS
		BuildArch = runtime.GOARCH
	}
}

const versionMessage = `
Version:     v%s
Commit:      %s

BuildDate:   %s
GoVersion:   %s
OS / Arch:   %s %s
`

var VersionLong = `Print the build version for dma`

var VersionCmd = &cobra.Command{

	Use: "version",

	Aliases: []string{
		"ver",
	},

	Short: "print the version",

	Long: VersionLong,

	Run: func(cmd *cobra.Command, args []string) {

		s, e := os.UserConfigDir()
		fmt.Printf("dma ConfigDir %q %v\n", filepath.Join(s, "dma"), e)

		fmt.Printf(
			versionMessage,
			Version,
			Commit,
			BuildDate,
			GoVersion,
			BuildOS,
			BuildArch,
		)
	},
}

func init() {
	RootCmd.AddCommand(VersionCmd)
}
