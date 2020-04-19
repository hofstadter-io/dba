package libcmdmodelset

import (
	"fmt"

	"github.com/hofstadter-io/dma/lib/cuefig"
)

func CreateRun(name string) (err error) {

	// TODO, check to see if it exists
	C := cuefig.Dma

	msets := C["modelsets"].(map[string]interface{})

	_, ok := msets[name]
	if ok {
		return fmt.Errorf("modelset with name %q already exists", name)
	}

	return err
}
