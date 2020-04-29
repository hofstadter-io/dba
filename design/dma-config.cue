package design

import (
	"github.com/hofstadter-io/hofmod-cuefig/schema"
)

#DmaConfig: schema.#Config & {
  Name: "dma"
  Entrypoint: "\(#CLI.ConfigDir)/config.cue"

  ConfigSchema: {
    models: {
      name: "string"
    }
    stores: {
      name: "string"
      type: "string"
    }
  }
}
