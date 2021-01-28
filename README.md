# xicomlin
A CLI library.

## Example
```go
package main

import (
    "os"

    "github.com/ximfect/xicomlin"
)

func main() {
    // create our tool
    tool := ximcomlin.NewTool("My Tool", "1.0.0", "Made using xicomlin :)")
    // add help action
    tool.AddAction("help", &ximcomlin.HelpAction{})
    // run action from args
    name := os.Args[0]
    argv := os.Args[1:]
    tool.RunAction(name, argv)
}
```