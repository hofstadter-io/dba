package design

import (
	"github.com/hofstadter-io/hofmod-cli/schema"
)

#GenerateCommand: schema.#Command & {
  Name:    "generate"
  Usage:   "generate"
  Aliases: ["gen"]
  Short:   "generate libraries and utilies for your data model"
  Long:    Short
},

