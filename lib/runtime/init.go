package runtime

import (
	"fmt"

	"cuelang.org/go/cue/errors"

	"github.com/hofstadter-io/dma/lib/cuefig"
)

var DMAConfig DmaConfig

func Init() {

	val, _ := cuefig.LoadDmaDefault(&DMAConfig)
	DMAConfig.CueValue = val

	// Get top level struct from cuelang
	S, err := val.Struct()
	if err != nil {
		// fmt.Println("STRUCT ERR", err)
		es := errors.Errors(err)
		for _, e := range es {
			fmt.Println(e)
		}
		return
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

}
