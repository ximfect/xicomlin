package xicomlin_test

import (
	"errors"
	"testing"

	"github.com/ximfect/xicomlin"
)

type testAction struct{}

func (x *testAction) GetDesc() string {
	return "hello. this is an action. it does what actions do."
}

func (x *testAction) GetArgs() *xicomlin.Arguments {
	pargs := []string{"type"}
	nargs := make(map[string]xicomlin.NargV)
	nargs["name"] = xicomlin.NargV{false, "\"project name\"", false}
	return &xicomlin.Arguments{pargs, nargs}
}

func (x *testAction) Exec(t *xicomlin.Tool, a *xicomlin.Arguments) error {
	return errors.New("hi")
}

func TestHelp(t *testing.T) {
	tool := xicomlin.NewTool(
		"xicomlin test tool",
		"action-test",
		"this is a tool for testing purposes. "+
			"this description is unnecessarily long, "+
			"in order to test whether the help action "+
			"can properly split this into separate lines.")
	tool.AddAction("help", &xicomlin.HelpAction{})
	tool.AddAction("new", &testAction{})
	tool.RunAction("new", []string{})
}
