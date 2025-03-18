package handler

import (
	"context"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/mehdihadeli/go-mediatr"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/adapter/http/request"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/adapter/http/response"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/app/command/emailverify"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/app/command/passwordreset"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/app/command/passwordresetrequest"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/app/command/passwordset"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/app/command/tokenrefresh"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/app/command/userlogin"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/app/command/usersignup"
)

type AuthHandler struct {
	Handler
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{
		Handler: NewHandler(),
	}
}

// SignupUserHandler godoc
// @Summary User signup
// @Description This endpoint allows users to sign up by providing their details like first name, last name, email, etc.
// @Tags Authentication
// @Accept json
// @Produce json
// @Param body body request.RegisterRequest true "User signup details"
// @Success 201 {object} response.UserSignupResponse
// @Failure 400 {object} response.ErrorResponse
// @Router /auth/signup [post]
func (handler *AuthHandler) SignupUserHandler(c echo.Context) error {
	signupRequest := &request.RegisterRequest{}
	signupResponse := &response.UserSignupResponse{}

	if err := handler.BindAndValidate(c, signupRequest); err != nil {
		return echo.ErrBadRequest
	}

	if err := signupRequest.Validate(); err != nil {
		return echo.ErrBadRequest
	}

	command := signupRequest.ToCommand()

	if err := command.Validate(); err != nil {
		return echo.ErrBadRequest
	}

	result, err := mediatr.Send[*usersignup.UserSignupCommand, *usersignup.UserSignupCommandResponse](context.Background(), command)

	if err != nil {
		return err
	}

	signupResponse.FromCommand(result)

	return c.JSON(201, signupResponse)
}

// LoginUserHandler godoc
// @Summary User login
// @Description Allows users to log in using their email and password, returning access and refresh tokens.
// @Tags Authentication
// @Accept json
// @Produce json
// @Param body body request.LoginRequest true "User login details"
// @Success 200 {object} response.UserLoginResponse
// @Failure 400 {object} response.ErrorResponse
// @Router /auth/login [post]
func (handler *AuthHandler) LoginUserHandler(c echo.Context) error {
	userLoginRequest := &request.LoginRequest{}
	userLoginResponse := &response.UserLoginResponse{}

	if err := handler.BindAndValidate(c, userLoginRequest); err != nil {
		return echo.ErrBadRequest
	}

	if err := userLoginRequest.Validate(); err != nil {
		return echo.ErrBadRequest
	}

	command := userLoginRequest.ToCommand()

	if err := command.Validate(); err != nil {
		return echo.ErrBadRequest
	}

	result, err := mediatr.Send[*userlogin.UserLoginCommand, *userlogin.UserLoginCommandResponse](context.Background(), command)

	if err != nil {
		return err
	}

	userLoginResponse.FromCommand(result)

	return c.JSON(200, userLoginResponse)
}

// RefreshTokenHandler godoc
// @Summary Refresh access token
// @Description Allows users to refresh their access token by providing a valid refresh token.
// @Tags Authentication
// @Accept json
// @Produce json
// @Success 200 {object} response.TokenRefreshResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 401 {object} response.ErrorResponse
// @Router /auth/refresh-token [post]
func (handler *AuthHandler) RefreshTokenHandler(c echo.Context) error {
	UserID := c.Get("user_id").(uuid.UUID)
	Token := c.Get("token").(string)

	tokenRefreshResponse := &response.TokenRefreshResponse{}

	command := &tokenrefresh.TokenRefreshCommand{
		UserID:       UserID,
		RefreshToken: Token,
	}

	if err := command.Validate(); err != nil {
		return echo.ErrBadRequest
	}

	result, err := mediatr.Send[*tokenrefresh.TokenRefreshCommand, *tokenrefresh.TokenRefreshCommandResponse](context.Background(), command)
	if err != nil {
		return err
	}
	tokenRefreshResponse.FromCommand(result)

	return c.JSON(200, tokenRefreshResponse)
}

// RequestPasswordResetHandler godoc
// @Summary Request password reset
// @Description Allows users to request a password reset by providing their email.
// @Tags Authentication
// @Accept json
// @Produce json
// @Param body body request.PasswordResetRequestRequest true "Password reset request details"
// @Success 200 {object} response.PasswordResetRequestResponse
// @Failure 400 {object} response.ErrorResponse
// @Router /auth/request-password-reset [post]
func (handler *AuthHandler) RequestPasswordResetHandler(c echo.Context) error {
	passwordResetRequestRequest := &request.PasswordResetRequestRequest{}
	passwordResetRequestResponse := &response.PasswordResetRequestResponse{}

	if err := handler.BindAndValidate(c, passwordResetRequestRequest); err != nil {
		return echo.ErrBadRequest
	}

	if err := passwordResetRequestRequest.Validate(); err != nil {
		return echo.ErrBadRequest
	}

	command := passwordResetRequestRequest.ToCommand()

	if err := command.Validate(); err != nil {
		return echo.ErrBadRequest
	}

	result, err := mediatr.Send[*passwordresetrequest.PasswordResetRequestCommand, *passwordresetrequest.PasswordResetRequestCommandResponse](context.Background(), command)
	if err != nil {
		return err
	}

	passwordResetRequestResponse.FromCommand(result)

	return c.JSON(200, passwordResetRequestResponse)
}

// ResetPasswordHandler godoc
// @Summary Reset user password
// @Description Allows users to reset their password by providing a new password and reset token.
// @Tags Authentication
// @Accept json
// @Produce json
// @Param body body request.PasswordResetRequest true "Password reset details"
// @Success 200 {object} response.PasswordResetResponse
// @Failure 400 {object} response.ErrorResponse
// @Router /auth/reset-password [post]
func (handler *AuthHandler) ResetPasswordHandler(c echo.Context) error {
	userId := c.Get("user_id").(uuid.UUID)
	passwordResetToken := c.Get("token").(string)

	passwordResetRequest := &request.PasswordResetRequest{}
	passwordResetResponse := &response.PasswordResetResponse{}

	if err := handler.BindAndValidate(c, passwordResetRequest); err != nil {
		return echo.ErrBadRequest
	}

	if err := passwordResetRequest.Validate(); err != nil {
		return echo.ErrBadRequest
	}

	command := &passwordreset.PasswordResetCommand{
		UserID:     userId,
		Password:   passwordResetRequest.Password,
		ResetToken: passwordResetToken,
	}

	if err := command.Validate(); err != nil {
		return echo.ErrBadRequest
	}

	result, err := mediatr.Send[*passwordreset.PasswordResetCommand, *passwordreset.PasswordResetCommandResponse](context.Background(), command)
	if err != nil {
		return err
	}

	passwordResetResponse.FromCommand(result)

	return c.JSON(200, passwordResetResponse)
}

// VerifyEmailHandler godoc
// @Summary Verify user email
// @Description Allows users to verify their email by providing a verification token.
// @Tags Authentication
// @Accept json
// @Produce json
// @Success 200 {object} response.EmailVerifyResponse
// @Failure 400 {object} response.ErrorResponse
// @Router /auth/verify-email [post]
func (handler *AuthHandler) VerifyEmailHandler(c echo.Context) error {
	userId := c.Get("user_id").(uuid.UUID)
	passwordResetToken := c.Get("token").(string)

	emailVerifyResponse := &response.EmailVerifyResponse{}

	command := &emailverify.EmailVerifyCommand{
		UserID: userId,
		Token:  passwordResetToken,
	}

	if err := command.Validate(); err != nil {
		return echo.ErrBadRequest
	}

	result, err := mediatr.Send[*emailverify.EmailVerifyCommand, *emailverify.EmailVerifyCommandResponse](context.Background(), command)
	if err != nil {
		return err
	}

	emailVerifyResponse.FromCommand(result)

	return c.JSON(200, emailVerifyResponse)
}

// ResetPasswordHandler godoc
// @Summary Reset user password
// @Description Allows users to reset their password by providing a new password and set token.
// @Tags Authentication
// @Accept json
// @Produce json
// @Param body body request.PasswordSetRequest true "Password set details"
// @Success 200 {object} response.PasswordSetResponse
// @Failure 400 {object} response.ErrorResponse
// @Router /auth/set-password [post]
func (handler *AuthHandler) SetPasswordHandler(c echo.Context) error {
	userId := c.Get("user_id").(uuid.UUID)
	passwordSetToken := c.Get("token").(string)

	passwordSetRequest := &request.PasswordSetRequest{}
	passwordSetResponse := &response.PasswordSetResponse{}

	if err := handler.BindAndValidate(c, passwordSetRequest); err != nil {
		return echo.ErrBadRequest
	}

	if err := passwordSetRequest.Validate(); err != nil {
		return echo.ErrBadRequest
	}

	command := &passwordset.PasswordSetCommand{
		UserID:   userId,
		Password: passwordSetRequest.Password,
		SetToken: passwordSetToken,
	}

	if err := command.Validate(); err != nil {
		return echo.ErrBadRequest
	}

	result, err := mediatr.Send[*passwordset.PasswordSetCommand, *passwordset.PasswordSetCommandResponse](context.Background(), command)
	if err != nil {
		return err
	}

	passwordSetResponse.FromCommand(result)

	return c.JSON(200, passwordSetResponse)
}
