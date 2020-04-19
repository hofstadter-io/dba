package libcmdmodelset

import (
	"fmt"
	"os"
	"strings"

	"github.com/jedib0t/go-pretty/table"

	"github.com/hofstadter-io/dma/lib/cuefig"
)

func ListRun() (err error) {

	t := table.NewWriter()
	t.Style().Options = table.OptionsNoBordersAndSeparators
	t.Style().Box.PaddingLeft = ""
	t.Style().Box.PaddingRight = "  "
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"modelset", "entry", "stores"})

	C := cuefig.Dma

	msets := C["modelsets"].(map[string]interface{})

	for name, set := range msets {
		S := set.(map[string]interface{})
		entry := fmt.Sprintf("%s", S["entry"])
		stores := S["stores"].(map[string]interface{})
		snames := []string{}
		for sn, _ := range stores {
			snames = append(snames, sn)
		}

		t.AppendRows([]table.Row{ {name, entry, strings.Join(snames, ",")}, })
	}

	t.Render()

	return err
}
