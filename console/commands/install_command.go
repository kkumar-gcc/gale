package commands

import (
	"github.com/goravel/framework/contracts/console"
	"github.com/goravel/framework/contracts/console/command"
)

type InstallCommand struct {
}

func (receiver *InstallCommand) Signature() string {
	return "install"
}

func (receiver *InstallCommand) Description() string {
	return "Install the components of the Gale package"
}

// Extend The console command extend.
func (receiver *InstallCommand) Extend() command.Extend {
	return command.Extend{
		Flags: []command.Flag{
			&command.BoolFlag{
				Name:    "force",
				Aliases: []string{"f"},
				Usage:   "Forces install even if the directory already exists",
			},
		},
	}
}

// Handle Execute the console command.
func (receiver *InstallCommand) Handle(ctx console.Context) error {
	// Do something
	return nil
}
