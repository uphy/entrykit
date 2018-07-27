package main

import (
	"fmt"
	"os"

	"github.com/uphy/entrykit"

	_ "github.com/uphy/entrykit/codep"
	_ "github.com/uphy/entrykit/prehook"
	_ "github.com/uphy/entrykit/render"
	_ "github.com/uphy/entrykit/switch"
)

var Version string

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "-v", "--version":
			fmt.Println(Version)
			os.Exit(0)
		case "--symlink":
			entrykit.Symlink()
			os.Exit(0)
		}
	}
	entrykit.RunCmd()
}
