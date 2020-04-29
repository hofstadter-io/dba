package runtime

import (
	"fmt"

	"cuelang.org/go/cue"

	"github.com/hofstadter-io/dma/cmd/dma/pflags"
	"github.com/hofstadter-io/dma/lib/cuefig"
	"github.com/hofstadter-io/dma/lib/types"
)

var rt *Runtime

func init() {
	rt = NewRuntime()
}

func Init() error {
	r := NewRuntime()
	err := r.Init()
	if err != nil {
		return err
	}

	rt = r

	return nil
}

func GetRuntime() *Runtime {
	return rt
}

type Runtime struct {
	DmaConfig *types.DmaConfig
	DmaCueVal cue.Value
}

func NewRuntime() *Runtime {
	return &Runtime{}
}

func (R *Runtime) Init() error {

	R.DmaConfig = &types.DmaConfig{}
	var val cue.Value
	var err error
	if pflags.RootConfigPflag != "" {
		val, err = cuefig.LoadDmaConfig(pflags.RootConfigPflag, R.DmaConfig)
	} else {
		val, err = cuefig.LoadDmaDefault(R.DmaConfig)
	}

	if err != nil {
		return err
	}
	R.DmaCueVal = val
	return nil
}

func (R *Runtime) Print() error {
	// Get top level struct from cuelang
	S, err := R.DmaCueVal.Struct()
	if err != nil {
		return err
	}

	iter := S.Fields()
	for iter.Next() {

		label := iter.Label()
		value := iter.Value()
		fmt.Println("  -", label, value)
		for attrKey, attrVal := range value.Attributes() {
			fmt.Println("  --", attrKey)
			for i := 0; i < 5; i++ {
				str, err := attrVal.String(i)
				if err != nil {
					break
				}
				fmt.Println("  ---", str)
			}
		}
	}

	return nil
}
