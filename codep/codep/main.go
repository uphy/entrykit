package main

import (
	"github.com/uphy/entrykit"
	_ "github.com/uphy/entrykit/codep"
)

var cmd = "codep"

func main() {
	entrykit.Cmds[cmd](
		entrykit.NewConfig(cmd, true))
}
