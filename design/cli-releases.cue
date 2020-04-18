package design

import (
	"github.com/hofstadter-io/hofmod-cli/schema"
)

CliReleases :: schema.GoReleaser & {
  Author:   "Hofstadter, Inc"
  Homepage: "https://github.com/hofstadter-io/dba"

  GitHub: {
    Owner: "hofstadter-io"
    Repo:  "dba"
  }

  Docker: {
    Maintainer: "Hofstadter, Inc <support@hofstadter.io>"
    Repo: "hofstadter"
  }
}


