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

  Updates: {
    CheckURL: "https://updates.dmatool.com/check"
    DevCheckURL: "http://localhost:8080/check"
  }

	OmitRun: true

  PersistentPrerun: true
  PersistentPostrun: true

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
