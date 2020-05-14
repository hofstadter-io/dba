module github.com/hofstadter-io/dma

go 1.14

require (
	cuelang.org/go v0.1.3-0.20200424192631-12927e83d318
	github.com/go-openapi/strfmt v0.19.5 // indirect
	github.com/google/uuid v1.1.1
	github.com/hofstadter-io/hof v0.3.8
	github.com/hofstadter-io/structural v0.1.3 // indirect
	github.com/hofstadter-io/yagu v0.0.3
	github.com/jedib0t/go-pretty v4.3.0+incompatible // indirect
	github.com/mattn/go-runewidth v0.0.9 // indirect
	github.com/parnurzeal/gorequest v0.2.16
	github.com/spf13/cobra v1.0.0
)

replace cuelang.org/go => ../../../cue/cue

replace github.com/hofstadter-io/structural => ../structural
