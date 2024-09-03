package api

import (
	"github.com/goravel/framework/support/collect"

	"github.com/kkumar-gcc/gale/contracts"
	"github.com/kkumar-gcc/gale/stubs/api/app/constants"
	controllers "github.com/kkumar-gcc/gale/stubs/api/app/http/controllers/auth"
	requests "github.com/kkumar-gcc/gale/stubs/api/app/http/requests/auth"
	"github.com/kkumar-gcc/gale/stubs/api/app/models"
	services "github.com/kkumar-gcc/gale/stubs/api/app/services/auth"
	"github.com/kkumar-gcc/gale/stubs/api/app/utils"
	"github.com/kkumar-gcc/gale/stubs/api/database/migrations"
	"github.com/kkumar-gcc/gale/stubs/api/routes"
	"github.com/kkumar-gcc/gale/support"
)

type Kernel struct {
	driver string
}

func NewKernel(driver string) *Kernel {
	return &Kernel{driver: driver}
}

var (
	ControllerPath = "app/http/controllers/auth"
	RequestPath    = "app/http/requests/auth"
	ServicePath    = "app/services/auth"
	ConstantsPath  = "app/constants"
	RoutePath      = "routes"
	UtilsPath      = "app/utils"
	ModelPath      = "app/models"
	MigrationPath  = "database/migrations"
)

func (kernel *Kernel) Stubs() map[string]contracts.TemplateStub {
	var stubs map[string]contracts.TemplateStub

	switch kernel.driver {
	case support.MysqlDriver:
		stubs = kernel.mysqlStubs()
	case support.PostgresDriver:
		stubs = kernel.postgresStubs()
	case support.SqliteDriver:
		stubs = kernel.sqliteStubs()
	default:
		stubs = map[string]contracts.TemplateStub{}
	}

	return collect.Merge(stubs, kernel.commonStubs())
}

func (kernel *Kernel) commonStubs() map[string]contracts.TemplateStub {
	return map[string]contracts.TemplateStub{
		ControllerPath + "/forgot_password_controller.go": &controllers.ForgotPasswordController{},
		ControllerPath + "/login_controller.go":           &controllers.LoginController{},
		ControllerPath + "/password_controller.go":        &controllers.PasswordController{},
		ControllerPath + "/register_controller.go":        &controllers.RegisterController{},
		ControllerPath + "/verify_email_controller.go":    &controllers.VerifyEmailController{},
		RequestPath + "/login_request.go":                 &requests.LoginRequest{},
		RequestPath + "/register_request.go":              &requests.RegisterRequest{},
		ServicePath + "/hash.go":                          &services.HashService{},
		ServicePath + "/mail.go":                          &services.MailService{},
		ServicePath + "/password_reset.go":                &services.PasswordResetService{},
		ServicePath + "/user.go":                          &services.UserService{},
		ConstantsPath + "/messages.go":                    &constants.Message{},
		RoutePath + "/auth.go":                            &routes.AuthRoute{},
		UtilsPath + "/auth.go":                            &utils.AuthUtils{},
		UtilsPath + "/json_response.go":                   &utils.JsonResponse{},
		ModelPath + "/user.go":                            &models.User{},
		ModelPath + "/password_reset_token.go":            &models.PasswordResetToken{},
	}
}

func (kernel *Kernel) mysqlStubs() map[string]contracts.TemplateStub {
	return map[string]contracts.TemplateStub{
		MigrationPath + "/20240831171525_create_users_table.down.sql":           &migrations.MysqlUserDownStub{},
		MigrationPath + "/20240831171525_create_users_table.up.sql":             &migrations.MysqlUserUpStub{},
		MigrationPath + "/20240831171525_create_password_resets_table.down.sql": &migrations.MysqlPasswordResetDownStub{},
		MigrationPath + "/20240831171525_create_password_resets_table.up.sql":   &migrations.MysqlPasswordResetUpStub{},
	}
}

func (kernel *Kernel) postgresStubs() map[string]contracts.TemplateStub {
	return map[string]contracts.TemplateStub{
		MigrationPath + "/20240831171525_create_users_table.down.sql":           &migrations.PostgresUserDownStub{},
		MigrationPath + "/20240831171525_create_users_table.up.sql":             &migrations.PostgresUserUpStub{},
		MigrationPath + "/20240831171525_create_password_resets_table.down.sql": &migrations.PostgresPasswordResetDownStub{},
		MigrationPath + "/20240831171525_create_password_resets_table.up.sql":   &migrations.PostgresPasswordResetUpStub{},
	}
}

func (kernel *Kernel) sqliteStubs() map[string]contracts.TemplateStub {
	return map[string]contracts.TemplateStub{
		MigrationPath + "/20240831171525_create_users_table.down.sql":           &migrations.SqliteUserDownStub{},
		MigrationPath + "/20240831171525_create_users_table.up.sql":             &migrations.SqliteUserUpStub{},
		MigrationPath + "/20240831171525_create_password_resets_table.down.sql": &migrations.SqlitePasswordResetDownStub{},
		MigrationPath + "/20240831171525_create_password_resets_table.up.sql":   &migrations.SqlitePasswordResetUpStub{},
	}
}
