package design

import (
  "github.com/hofstadter-io/hofmod-cli/schema"
)

#Outdir: "./"
#Module: "github.com/hofstadter-io/dma"

#_LibImport: [
	schema.#Import & {Path: #Module + "/lib"},
]

#CLI: schema.#Cli & {
	Name:    "dma"
	Package: "\(#Module)/cmd/dma"

	Usage: "dma"
	Short: "Your Data Model Assistant"
	Long:  Short

	Releases: #CliReleases

	OmitRun: true

  PersistentPrerun: true

  Pflags: #CliPflags

	Commands: [
    // meta
    #InitCommand,
    #ConfigCommand,
    #UiCommand,

    // main
    #ModelCommand,
    #StoreCommand,

    // outer
    #ImportCommand,
    #ExportCommand,
    #GenerateCommand,
	]
}
