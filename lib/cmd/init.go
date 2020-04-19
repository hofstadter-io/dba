package libcmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/hofstadter-io/dma/lib/cuefig"
)

const initMetaContents = `
modelsets: {}

stores: {}
`

func InitRun() error {

	// meta filepath
	// check that dma/meta.cue exists, exit early
	_, err := os.Lstat(cuefig.DmaFilepath)
	if err != nil {
		if _, ok := err.(*os.PathError); !ok && ( strings.Contains(err.Error(), "file does not exist") || strings.Contains(err.Error(), "no such file") ) {
			// error is worse than non-existant
			return err
		}
		// otherwise, does not exist, so we should init
	} else {
		return fmt.Errorf("dma already initialized")
	}

	// Mkdir
	err = os.MkdirAll("dma", 0755)
	if err != nil {
		return err
	}

	// create meta.cue file
	err = ioutil.WriteFile(cuefig.DmaFilepath, []byte(initMetaContents), 0644)
	if err != nil {
		return err
	}

	return nil
}
