package main

import (
	"os"

	frameworkconsole "github.com/goravel/framework/console"

	"github.com/kkumar-gcc/gale/console"
	"github.com/kkumar-gcc/gale/support"
)

func main() {
	name := "Gale CLI"
	usage := "An authentication addon for Goravel projects."
	usageText := "gale [global options] command [command options] [arguments...]"

	cliApp := frameworkconsole.NewApplication(name, usage, usageText, support.Version, false)

	kernel := &console.Kernel{}

	cliApp.Register(kernel.Commands())
	cliApp.Run(os.Args, false)
}
