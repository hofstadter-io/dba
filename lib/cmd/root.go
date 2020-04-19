package libcmd

import (
	"github.com/hofstadter-io/dma/lib/cuefig"
)

func RootPersistentPreRun() (err error) {

	cuefig.LoadDefault()

	return nil
}
