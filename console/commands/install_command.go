package commands

import (
	"strings"

	"github.com/goravel/framework/contracts/console"
	"github.com/goravel/framework/contracts/console/command"
	"github.com/goravel/framework/support/color"
	"github.com/goravel/framework/support/file"
	"github.com/goravel/framework/support/str"

	apiStack "github.com/kkumar-gcc/gale/stubs/api"
	"github.com/kkumar-gcc/gale/support"
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
			&command.StringFlag{
				Name:    "stack",
				Aliases: []string{"s"},
				Usage:   "The stack that should be installed",
			},
			&command.StringFlag{
				Name:    "module",
				Aliases: []string{"m"},
				Usage:   "Module name of the current project",
			},
			&command.BoolFlag{
				Name:    "force",
				Aliases: []string{"f"},
				Usage:   "Forces install even if the directory already exists",
			},
		},
	}
}

// Handle Execute the console command.
func (receiver *InstallCommand) Handle(ctx console.Context) (err error) {
	stack := ctx.Option("stack")
	//force := ctx.Option("force")
	if stack == "" {
		stack, err = ctx.Choice("Which stack would you like to install?", []console.Choice{
			{"API", true, support.APIStack},
		})
		if err != nil {
			ctx.Error(err.Error())
			return nil
		}
	}

	stack = str.Of(stack).Lower().String()

	color.Infoln("Installing " + stack + " stack...")
	if stack == support.APIStack {
		return receiver.installApiStack(ctx)
	}

	ctx.Error("Invalid stack. Supported stacks is [api]")
	return nil
}

func (receiver *InstallCommand) installApiStack(ctx console.Context) error {
	module := ctx.Option("module")
	if module == "" {
		module = support.GoravelModulePath
	}

	driver := ctx.Option("driver")
	if driver == "" {
		driver = support.MysqlDriver
	}

	driver = str.Of(driver).Lower().String()

	kernel := apiStack.NewKernel(driver)
	stubs := kernel.Stubs()

	for name, stub := range stubs {
		ctx.Info("Creating " + name)

		if err := file.Create(name, receiver.populateStub(stub.Stub(), support.GoravelModulePath)); err != nil {
			ctx.Error(err.Error())
			return nil
		}
	}
	return nil
}

// populateStub Populate the place-holders in the command stub.
func (receiver *InstallCommand) populateStub(stub string, packageName string) string {
	return strings.ReplaceAll(stub, support.APIStackPath, packageName)
}
