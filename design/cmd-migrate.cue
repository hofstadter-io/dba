package design

import (
	"github.com/hofstadter-io/hofmod-cli/schema"
)

MigrateCommand :: schema.Command & {
  Name:    "migrate"
  Usage:   "migrate"
  Short:   "view, prepare, and make migrations to your data model."
  Long:    Short
},

