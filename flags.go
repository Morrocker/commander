package commander

import (
	"fmt"

	"github.com/spf13/pflag"
)

func String(name string, shorthand string, defVal string, usage string) *string {
	var newString *string

	if defVal != "" {
		usage = fmt.Sprintf("%s [ default: %s ]", usage, defVal)
	}
	newString = pflag.StringP(name, shorthand, defVal, usage)
	return newString
}
