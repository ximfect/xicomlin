package xicomlin_test

import (
	"fmt"
	"testing"

	"github.com/ximfect/xicomlin"
)

func TestArgv(t *testing.T) {
	args := []string{"new", "python", "--name", "test", "--!venv"}

	fmt.Println(xicomlin.ParseArgv(args))
}
