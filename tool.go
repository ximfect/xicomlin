package xicomlin

import (
	"errors"
	"fmt"
	"strings"
)

// Tool represents a CLI
type Tool struct {
	Actions map[string]Action
	Name    string
	Version string
	Desc    string
}

// NewTool creates and returns a new Tool
func NewTool(name, version, desc string) *Tool {
	return &Tool{
		make(map[string]Action),
		name, version, desc}
}

// AddAction adds an Action to this Tool under the given name
func (t *Tool) AddAction(n string, a Action) bool {
	n = strings.ToLower(n)
	_, e := t.Actions[n]
	if e {
		return false
	}
	t.Actions[n] = a
	return true
}

// RemAction removes the Action with the given name
func (t *Tool) RemAction(n string) {
	n = strings.ToLower(n)
	_, e := t.Actions[n]
	if e {
		delete(t.Actions, n)
	}
}

func (t *Tool) runActionInternal(n string, a []string) error {
	argv, err := ParseArgv(a)
	if err != nil {
		return err
	}
	n = strings.ToLower(n)
	action, ok := t.Actions[n]
	if !ok {
		return errors.New("unknown action `" + n + "`")
	}
	return action.Exec(t, argv)
}

// RunAction will run the given action
func (t *Tool) RunAction(name string, argv []string) {
	err := t.runActionInternal(name, argv)
	if err != nil {
		fmt.Println(name+": ERROR:", err.Error())
	}
}
