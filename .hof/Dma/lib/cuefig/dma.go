package cuefig

// Name: dma

import (
	"fmt"
	"os"
	"strings"

	"cuelang.org/go/cue"
	// "cuelang.org/go/cue/build"
	"cuelang.org/go/cue/errors"
	"cuelang.org/go/cue/load"

	"github.com/hofstadter-io/hof/lib/util"
)

var DmaFilepath string

func init() {
	cfgdir, err := os.UserConfigDir()
	if err == nil {
		DmaFilepath = filepath.Join(cfgdir, "dma/config.cue")
	} else {
		DmaFilepath = "dma/config.cue"
	}
}

func LoadDmaDefault(cfg interface{}) (cue.Value, error) {
	return LoadDmaConfig(DmaFilepath, cfg)
}

func LoadDmaConfig(entry string, cfg interface{}) (val cue.Value, err error) {

	_, err = os.Lstat(DmaFilepath)
	if err != nil {
		if _, ok := err.(*os.PathError); !ok && (strings.Contains(err.Error(), "file does not exist") || strings.Contains(err.Error(), "no such file")) {
			// error is worse than non-existant
			return val, err
		}
		// otherwise, does not exist, so we should init?
		// XXX want to let applications decide how to handle this
		return val, err
	}

	var errs []error

	CueRT := &cue.Runtime{}
	BIS := load.Instances([]string{entry}, nil)
	for _, bi := range BIS {

		if bi.Err != nil {
			// fmt.Println("BI ERR", bi.Err, bi.Incomplete, bi.DepsErrors)
			es := errors.Errors(bi.Err)
			for _, e := range es {
				errs = append(errs, e.(error))
			}
			continue
		}

		// Build the Instance
		I, err := CueRT.Build(bi)
		if err != nil {
			es := errors.Errors(err)
			// fmt.Println("BUILD ERR", es, I)
			for _, e := range es {
				errs = append(errs, e.(error))
			}
			continue
		}

		// Get top level value from cuelang
		V := I.Value()

		err = V.Decode(&cfg)
		if err != nil {
			errs = append(errs, err)
			continue
		}

		val = V

	}

	if len(errs) > 0 {
		for _, e := range errs {
			util.PrintCueError(e)
		}
		return val, fmt.Errorf("Errors while reading DMA config file")
	}

	return val, nil
}
