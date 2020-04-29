package models

import (
  "github.com/hofstadter-io/dma/schema"
)

@modelset()
Modelset: schema.#Modelset & {
  Tags: ["foo", "goo"]
}
