package design

import (
	"github.com/hofstadter-io/hofmod-cli/schema"
)

ConfigCommand :: schema.Command & {
  Name:    "config"
  Usage:   "config"
  Short:   "view and set, global and local config values"
  Long:    Short
},


