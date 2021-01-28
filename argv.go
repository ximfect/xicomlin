package xicomlin

import (
	"errors"
	"strings"
)

// NargV represents the value of a named argument in the Arguments struct
type NargV struct {
	IsBool    bool
	StringVal string
	BoolVal   bool
}

// Arguments represents an Action's arguments, whether requested by the tool or
// passed to the Exec function.
type Arguments struct {
	Pargs []string
	Nargs map[string]NargV
}

// EmptyArgv returns an empty Arguments object
func EmptyArgv() *Arguments {
	return &Arguments{[]string{}, make(map[string]NargV)}
}

// ParseArgv parses a string slice into an Arguments object
func ParseArgv(src []string) (*Arguments, error) {
	var (
		pargs []string
		nargs = make(map[string]NargV)

		inNargs bool
		hasKey  bool
		key     string
	)

	for _, s := range src {
		if !inNargs {
			if strings.HasPrefix(s, "--") {
				inNargs = true
				key = s[2:]
				hasKey = true
			} else {
				pargs = append(pargs, s)
			}
		} else {
			if strings.HasPrefix(s, "--") {
				if !hasKey {
					key = s[2:]
					hasKey = true
				} else {
					if key[0] == '!' {
						nargs[key[1:]] = NargV{true, "", false}
					} else {
						nargs[key] = NargV{true, "", true}
					}
					key = s[2:]
				}
			} else {
				if !hasKey {
					return nil, errors.New(
						"expected key, but got `" + s + "` instead")
				}
				nargs[key] = NargV{false, s, true}
				hasKey = false
			}
		}
	}

	if inNargs && hasKey {
		if key[0] == '!' {
			nargs[key[1:]] = NargV{true, "", false}
		} else {
			nargs[key] = NargV{true, "", true}
		}
	}

	return &Arguments{pargs, nargs}, nil
}

// FormatArgv formats the given Arguments, as seen in the example HelpAction
func FormatArgv(a *Arguments) string {
	pargs := []string{}

	for _, parg := range a.Pargs {
		pargs = append(pargs, "<"+parg+">")
	}

	nargs := []string{}

	for name, narg := range a.Nargs {
		if narg.IsBool {
			nargs = append(nargs, "[--"+name+"]")
		} else {
			if narg.BoolVal {
				nargs = append(nargs, "<--"+name+" "+narg.StringVal+">")
			} else {
				nargs = append(nargs, "[--"+name+" "+narg.StringVal+"]")
			}
		}
	}

	return strings.Join(pargs, " ") + " " + strings.Join(nargs, " ")
}
