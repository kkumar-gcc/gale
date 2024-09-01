package routes

type AuthRoute struct{}

func (r *AuthRoute) Stub() string {
	return `package routes

import (
	"github.com/goravel/framework/facades"

	"github.com/kkumar-gcc/gale/stubs/api/app/http/controllers/auth"
	authServices "github.com/kkumar-gcc/gale/stubs/api/app/services/auth"
)

func Auth() {
	route := facades.Route()
	userService := authServices.NewUserImpl()
	hashService := authServices.NewHashImpl()
	mailService := authServices.NewMailImpl()
	passwordResetService := authServices.NewPasswordResetImpl()

	loginController := auth.NewLoginController(userService, hashService)
	registerController := auth.NewRegisterController(userService, hashService, mailService)
	newPasswordController := auth.NewPasswordController(userService, passwordResetService)
	forgotPasswordController := auth.NewForgotPasswordController(userService, passwordResetService, mailService)
	verifyEmailController := auth.NewVerifyEmailController()

	route.Middleware().Post("/login", loginController.Store)
	route.Post("/register", registerController.Store)
	route.Post("/forgot-password", forgotPasswordController.Store)
	route.Post("/reset-password", newPasswordController.Store)
	route.Get("/verify-email/{id}/{hash}", verifyEmailController.Store)
	route.Post("/logout", loginController.Destroy)
}
`
}
