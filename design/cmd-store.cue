package design

import (
	"github.com/hofstadter-io/hofmod-cli/schema"
)

StoreCommand :: schema.Command & {
  Name:    "store"
  Usage:   "store"
  Short:   "create, checkpoint, and migrate your storage engines"
  Long:    Short
},
