package design

import (
	"github.com/hofstadter-io/hofmod-cli/schema"
)

#CliPflags: [...schema.#Flag] & [
  {
    Name:    "config"
    Long:    "config"
    Short:   "c"
    Type:    "string"
    Default: ""
    Help:    "The path to a dma config file"
  },
  {
    Name:    "creds"
    Long:    "creds"
    Short:   "C"
    Type:    "string"
    Default: ""
    Help:    "The path to a dma creds file"
  },
]


