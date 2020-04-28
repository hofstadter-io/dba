package design

import (
	"github.com/hofstadter-io/hofmod-cli/schema"
)

// Name arg
#nameArg: schema.#Arg & {
	Name:     "name"
	Type:     "string"
	Required: true
	Help:     "A name from /[a-zA-Z][a-zA-Z0-9_]*"
}

// Identifyier (name or id)
#identArg: schema.#Arg & {
	Name:     "ident"
	Type:     "string"
	Required: true
	Help:     "A name or id"
}

// input arg
#inputArg: schema.#Arg & {
	Name:     "input"
	Type:     "string"
	Required: true
}

#contextArg: schema.#Arg & {
	Name: "context"
	Type: "string"
	Help: "A context name"
}

