package runtime

import (
	"cuelang.org/go/cue"

	"github.com/hofstadter-io/dma/lib/modelset"
	"github.com/hofstadter-io/dma/lib/store"
)

type DmaConfig struct {
	Modelsets map[string]modelset.Modelset
	Stores    map[string]store.Store

	CueValue  cue.Value
}
