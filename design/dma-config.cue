package design

import (
	"github.com/hofstadter-io/hofmod-cuefig/schema"
)

#DmaConfig: schema.#Config & {
  Name: "dma"
  Entrypoint: "dma/meta.cue"

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
