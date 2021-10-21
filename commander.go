package commander

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/pflag"
)

func ParseCommands() {
	pflag.Parse()
	spew.Dump(pflag.Args())

}
