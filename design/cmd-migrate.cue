package design

import (
	"github.com/hofstadter-io/hofmod-cli/schema"
)

MigrateCommand :: schema.Command & {
  Name:    "migrate"
  Usage:   "migrate"
  Short:   "view, prepare, and make migrations to your data model."
  Long:    Short

  OmitRun: true

  Commands: [...schema.Command] & [
    {
      Name:  "create"
      Usage: "create"
      Short: "create a migreation for a datastore"
      Long:  Short

      Args: [...schema.Arg] & [
        {
          Name:     "name"
          Type:     "string"
          Required: true
          Help:     "datastore name"
        },
      ]
    },
    {
      Name:  "view"
      Usage: "view"
      Short: "view migration information for a datastore"
      Long:  Short
      Args: [...schema.Arg] & [
        {
          Name:     "name"
          Type:     "string"
          Required: true
          Help:     "datastore name"
        },
      ]
    },
    {
      Name:  "list"
      Usage: "list"
      Short: "list the known datastores and migration briefs"
      Long:  Short
    },
    {
      Name:  "status"
      Usage: "status"
      Short: "show the current status for a datastore"
      Long:  Short
      Args: [...schema.Arg] & [
        {
          Name:     "name"
          Type:     "string"
          Required: true
          Help:     "datastore name"
        },
      ]
    },
    {
      Name:  "diff"
      Usage: "diff"
      Short: "show the current diff for a datastore"
      Long:  Short
      Args: [...schema.Arg] & [
        {
          Name:     "name"
          Type:     "string"
          Required: true
          Help:     "datastore name"
        },
      ]
    },
    {
      Name:  "test"
      Usage: "test"
      Short: "test the current migration and diff for a datastore"
      Long:  Short
      Args: [...schema.Arg] & [
        {
          Name:     "name"
          Type:     "string"
          Required: true
          Help:     "datastore name"
        },
      ]
    },
    {
      Name:  "delete"
      Usage: "delete"
      Short: "delete a datastore permentantly"
      Long:  Short
      Args: [...schema.Arg] & [
        {
          Name:     "name"
          Type:     "string"
          Required: true
          Help:     "datastore name"
        },
      ]
    },

  ]
},

