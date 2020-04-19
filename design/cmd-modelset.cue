package design

import (
	"github.com/hofstadter-io/hofmod-cli/schema"
)

ModelCommand :: schema.Command & {
  Name:    "modelset"
  Usage:   "modelset"
  Aliases: ["model", "m"]
  Short:   "create, view, migrate, and understand your data models."
  Long:    Short

  OmitRun: true

  Commands: [...schema.Command] & [
    {
      Name:  "create"
      Usage: "create"
      Short: "create a modelset"
      Long:  Short

      Args: [...schema.Arg] & [
        {
          Name:     "name"
          Type:     "string"
          Required: true
          Help:     "modelset name"
        },
      ]
    },
    {
      Name:  "view"
      Usage: "view"
      Short: "view modelset information"
      Long:  Short
      Args: [...schema.Arg] & [
        {
          Name:     "name"
          Type:     "string"
          Required: true
          Help:     "modelset name"
        },
      ]
    },
    {
      Name:  "list"
      Usage: "list"
      Short: "list the known modelsets"
      Long:  Short
    },
    {
      Name:  "status"
      Usage: "status"
      Short: "show the current status for a modelset"
      Long:  Short
      Args: [...schema.Arg] & [
        {
          Name:     "name"
          Type:     "string"
          Required: true
          Help:     "modelset name"
        },
      ]
    },
    {
      Name:  "graph"
      Usage: "graph"
      Short: "show the relationship graph for a modelset"
      Long:  Short
      Args: [...schema.Arg] & [
        {
          Name:     "name"
          Type:     "string"
          Required: true
          Help:     "modelset name"
        },
      ]
    },
    {
      Name:  "diff"
      Usage: "diff"
      Short: "show the current diff for a modelset"
      Long:  Short
      Args: [...schema.Arg] & [
        {
          Name:     "name"
          Type:     "string"
          Required: true
          Help:     "modelset name"
        },
      ]
    },
    {
      Name:  "migrate"
      Usage: "migrate"
      Short: "create the next migration for a modelset"
      Long:  Short
      Args: [...schema.Arg] & [
        {
          Name:     "name"
          Type:     "string"
          Required: true
          Help:     "modelset name"
        },
      ]
    },
    {
      Name:  "test"
      Usage: "test"
      Short: "test the current migration and diff for a modelset"
      Long:  Short
      Args: [...schema.Arg] & [
        {
          Name:     "name"
          Type:     "string"
          Required: true
          Help:     "modelset name"
        },
      ]
    },
    {
      Name:  "delete"
      Usage: "delete"
      Short: "delete a modelset permentantly"
      Long:  Short
      Args: [...schema.Arg] & [
        {
          Name:     "name"
          Type:     "string"
          Required: true
          Help:     "modelset name"
        },
      ]
    },

  ]
},

