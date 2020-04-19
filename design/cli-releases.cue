package design

import (
	"github.com/hofstadter-io/hofmod-cli/schema"
)

CliReleases :: schema.GoReleaser & {
  Author:   "Hofstadter, Inc"
  Homepage: "https://github.com/hofstadter-io/dma"

  GitHub: {
    Owner: "hofstadter-io"
    Repo:  "dma"
  }

  Docker: {
    Maintainer: "Hofstadter, Inc <support@hofstadter.io>"
    Repo: "hofstadter"
  }
}


