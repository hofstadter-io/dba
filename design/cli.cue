package design

import (
  "github.com/hofstadter-io/hofmod-cli/schema"
)

Outdir :: "./"

_LibImport :: [
	schema.Import & {Path: CLI.Package + "/lib"},
]

CLI :: schema.Cli & {
	Name:    "dma"
	Package: "github.com/hofstadter-io/dma"

	Usage: "dma"
	Short: "Your Data Model Assistant"
	Long:  Short

	Releases: CliReleases

	OmitRun: true

  Pflags: CliPflags

	Commands: [
    // meta
    InitCommand,
    ConfigCommand,
    UiCommand,

    // main
    MigrateCommand,
    StoreCommand,

    // outer
    ImportCommand,
    ExportCommand,
    GenerateCommand,
	]
}
