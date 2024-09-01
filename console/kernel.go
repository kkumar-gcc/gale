package console

import (
	"github.com/goravel/framework/contracts/console"

	"github.com/kkumar-gcc/gale/console/commands"
)

type Kernel struct {
}

func (kernel *Kernel) Commands() []console.Command {
	return []console.Command{
		&commands.InstallCommand{},
	}
}
