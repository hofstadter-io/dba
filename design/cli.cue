package design

import (
  "github.com/hofstadter-io/hofmod-cli/schema"
)

Outdir :: "./" @reln("anystring")

_LibImport :: [
	schema.Import & {Path: CLI.Package + "/lib"},
]

CLI :: schema.Cli & {
	Name:    "dba"
	Package: "github.com/hofstadter-io/dba"

	Usage: "hof"
	Short: "Polyglot Code Gereration Framework"
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
