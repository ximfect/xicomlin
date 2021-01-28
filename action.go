package xicomlin

import (
	// imports for HelpAction, not used outside
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Action represents a Tool action
type Action interface {
	GetDesc() string
	GetArgs() *Arguments
	Exec(t *Tool, a *Arguments) error
}

// HelpAction is an (example) help action showing the Tool's name, version,
// description and all the avilable actions.
type HelpAction struct{}

// GetDesc returns this action's description
func (x *HelpAction) GetDesc() string {
	return "Shows this."
}

// GetArgs returns this action's arguments
func (x *HelpAction) GetArgs() *Arguments {
	a := EmptyArgv()
	a.Pargs = []string{"action?"}
	return a
}

// Exec runs the action with the given arguments
func (x *HelpAction) Exec(t *Tool, a *Arguments) error {
	switch len(a.Pargs) {
	case 0:
		title := t.Name + " " + t.Version
		desc := Linebreak(t.Desc, 80)
		actions := []string{}

		for name, action := range t.Actions {
			nameCol := PadLeft(
				More(name+" "+FormatArgv(action.GetArgs()), 29), 30)
			descCol := ": " + PadLeft(More(action.GetDesc(), 48), 50)
			actions = append(actions, nameCol+descCol)
		}

		fmt.Println(
			title + "\n\n" + desc + "\n\n" + strings.Join(actions, "\n") + "\n")

	case 1:
		execPath, err := os.Executable()
		if err != nil {
			return nil
		}
		execName := filepath.Base(execPath)
		var exec string
		if strings.HasSuffix(execName, ".exe") {
			execSplit := strings.Split(execName, ".")
			exec = strings.Join(execSplit[:len(execSplit)-1], ".")
		} else {
			exec = execName
		}
		name := strings.ToLower(a.Pargs[0])
		target, ok := t.Actions[name]
		if !ok {
			return errors.New("could not find action `" + name + "`")
		}

		title := Linebreak(
			exec+" "+name+" "+FormatArgv(target.GetArgs()), 80)
		desc := Linebreak(target.GetDesc(), 90)

		fmt.Println(title + "\n\n" + desc)

	default:
		return errors.New("too many parameters")
	}
	return nil
}
