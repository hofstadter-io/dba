package dma

import (
	"github.com/hofstadter-io/hofmod-cli:cli"
	"github.com/hofstadter-io/dma/design"
)

HofGenCli: cli.HofGenerator & {
  Outdir: "./"
  Cli: design.CLI
}
