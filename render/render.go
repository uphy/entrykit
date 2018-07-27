package render

import (
	"io/ioutil"

	"github.com/progrium/entrykit"
	"github.com/uphy/sigil"

	_ "github.com/uphy/sigil/builtin"
)

func init() {
	entrykit.Cmds["render"] = Run
}

func Run(config *entrykit.Config) {
	defer entrykit.Exec(config.Exec)
	for name, target := range config.Tasks {
		template := target + ".tmpl"
		data, err := ioutil.ReadFile(template)
		if err != nil {
			entrykit.Error(err)
		}
		render, err := sigil.Execute(data, nil, name)
		if err != nil {
			entrykit.Error(err)
		}
		// todo: use same filemode as template
		err = ioutil.WriteFile(target, render.Bytes(), 0644)
		if err != nil {
			entrykit.Error(err)
		}
	}
}
