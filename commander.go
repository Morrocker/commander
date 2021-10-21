package commander

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/pflag"
)

func ParseCommands() {
	x := pflag.Args()
	y := pflag.NArg()
	z := pflag.NFlag()

	spew.Dump(x)
	spew.Dump(y)
	spew.Dump(z)
}
