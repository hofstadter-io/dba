package libcmd

import (
	"fmt"

	"github.com/hofstadter-io/dma/lib/cuefig"
)

func RootPersistentPreRun() (err error) {

	cfg, err := cuefig.LoadDefault()
	fmt.Printf("CFG:\n%#+v\n%v\n", cfg, err)

	return err
}
