package design

import (
	"github.com/hofstadter-io/hofmod-cli/schema"
)

#CliPflags: [...schema.#Flag] & [
  {
    Name:    "config"
    Type:    "string"
    Default: ""
    Help:    "The path to a dma config file"
    Long:    "config"
    Short:   "c"
  },
  {
    Name:    "context"
    Long:    "context"
    Short:   "C"
    Type:    "string"
    Default: ""
    Help:    "the context to use during this hof execution"
  },
]


