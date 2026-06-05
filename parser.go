package mon

import (
	"os"
	"strings"
)

var (
	flagMap  map[string]*flag
	tokenIdx int
)

func init() {
	flagMap = make(map[string]*flag)
	tokenIdx = 1
}

func fillFlagMap(c *Command) {
	for _, f := range c.flags {
		for _, l := range f.longNames {
			if l != "" {
				flagMap["--"+l] = &f
			}
		}

		if f.shortName != "" {
			flagMap["-"+f.shortName] = &f
		}
	}
}

func isLongFlag(s string) bool {
	return len(s) > 2 && strings.HasPrefix(s, "--")
}

func isShortFlag(s string) bool {
	return len(s) > 1 && !isLongFlag(s) && strings.HasPrefix(s, "-")
}

func isFlag(s string) bool {
	return isLongFlag(s) || isShortFlag(s)
}

func (c *Command) parseFlags() {
	fillFlagMap(c)

	for tokenIdx < len(os.Args) {
		token := os.Args[tokenIdx]

		if isFlag(token) {
			f, found := flagMap[token]
			if !found {
				// Warning: Unrecognized flag
				continue
			}

			if f.requiresVal {
				if tokenIdx+1 < len(os.Args) && !isFlag(os.Args[tokenIdx+1]) {
					f.setValue(os.Args[tokenIdx+1])
					f.isValueSet = true
				} else {
					// Error: No value supplied for flag
				}
			} else {
				f.setValue("true")
				f.isValueSet = true
			}
		} else {
			for _, f := range c.flags {
				if f.isRequired && !f.isValueSet {
					// Error: No value supplied for Required Flag
				}
			}

			// TODO: Subcommand
		}
	}
}
