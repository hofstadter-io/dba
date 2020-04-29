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
]


