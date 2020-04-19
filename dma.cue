package dma

import (
	cuefig "github.com/hofstadter-io/hofmod-cuefig/gen"
	cli "github.com/hofstadter-io/hofmod-cli/gen"
	"github.com/hofstadter-io/dma/design"
)

HofGenCli: cli.HofGenerator & {
  Outdir: "./"
  Cli: design.CLI
}

HofGenDma: cuefig.HofGenerator & {
  Outdir: "./"
  Config: design.DmaConfig
}
