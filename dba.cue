package dba

import (
	"github.com/hofstadter-io/hofmod-cli:cli"
	"github.com/hofstadter-io/dba/design"
)

HofGenCli: cli.HofGenerator & {
  Outdir: "./"
  Cli: design.CLI
}
