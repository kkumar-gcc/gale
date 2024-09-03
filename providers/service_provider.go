package providers

import (
	"github.com/goravel/framework/contracts/console"
	"github.com/goravel/framework/contracts/foundation"

	"github.com/kkumar-gcc/gale/console/commands"
)

type ServiceProvider struct{}

func (receiver *ServiceProvider) Register(app foundation.Application) {

}

func (receiver *ServiceProvider) Boot(app foundation.Application) {
	app.Commands([]console.Command{
		&commands.InstallCommand{},
	})
}
