package api

import (
	"github.com/kkumar-gcc/gale/contracts"
	"github.com/kkumar-gcc/gale/stubs/api/app/constants"
	controllers "github.com/kkumar-gcc/gale/stubs/api/app/http/controllers/auth"
	requests "github.com/kkumar-gcc/gale/stubs/api/app/http/requests/auth"
	"github.com/kkumar-gcc/gale/stubs/api/app/models"
	services "github.com/kkumar-gcc/gale/stubs/api/app/services/auth"
	"github.com/kkumar-gcc/gale/stubs/api/app/utils"
	"github.com/kkumar-gcc/gale/stubs/api/routes"
)

type Kernel struct {
}

var (
	ControllerPath = "app/http/controllers/auth"
	RequestPath    = "app/http/requests/auth"
	ServicePath    = "app/services/auth"
	ConstantsPath  = "app/constants"
	RoutePath      = "routes"
	UtilsPath      = "app/utils"
	ModelPath      = "app/models"
)

func (kernel *Kernel) Stubs() map[string]contracts.TemplateStub {
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
