package auth

import (
	"github.com/goravel/framework/contracts/http"

	"github.com/kkumar-gcc/gale/stubs/api/app/constants"
	authServices "github.com/kkumar-gcc/gale/stubs/api/app/services/auth"
	"github.com/kkumar-gcc/gale/stubs/api/app/utils"
)

type ForgotPasswordController struct {
	userService          authServices.User
	passwordResetService authServices.PasswordReset
	mailService          authServices.Mail
}

func NewForgotPasswordController(userService authServices.User, passwordResetService authServices.PasswordReset, mailService authServices.Mail) *ForgotPasswordController {
	return &ForgotPasswordController{
		userService:          userService,
		passwordResetService: passwordResetService,
		mailService:          mailService,
	}
}

func (r *ForgotPasswordController) Store(ctx http.Context) http.Response {
	email := ctx.Request().Input("email")
	if email == "" {
		return utils.NewJsonResponse().SetMessage(constants.ErrBadRequest).SetErrors("Email is required").Build(ctx)
	}

	user, err := r.userService.FindByEmail(email)
	if err != nil {
		return utils.NewJsonResponse().SetMessage(constants.ErrUserNotFound).SetErrors(err.Error()).Build(ctx)
	}

	token, err := r.passwordResetService.GenerateToken(user.Email)
	if err != nil {
		return utils.NewJsonResponse().SetMessage(constants.ErrInternalServer).SetErrors(err.Error()).Build(ctx)
	}

	if err := r.mailService.SendPasswordResetEmail(user.Email, token); err != nil {
		return utils.NewJsonResponse().SetMessage(constants.ErrInternalServer).SetErrors(err.Error()).Build(ctx)
	}

	return utils.NewJsonResponse().SetMessage(constants.SuccessPasswordResetEmailSent).Build(ctx)
}
