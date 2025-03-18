package response

import (
	"github.com/google/uuid"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/app/command/emailverify"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/app/command/passwordreset"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/app/command/passwordresetrequest"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/app/command/passwordset"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/app/command/tokenrefresh"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/app/command/userlogin"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/app/command/usersignup"
)

type UserSignupResponse struct {
	UserID  uuid.UUID `json:"user_id"`
	Message string    `json:"message"`
}

func (u *UserSignupResponse) FromCommand(command *usersignup.UserSignupCommandResponse) {
	u.UserID = command.UserID
	u.Message = command.Message
}

type UserLoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (u *UserLoginResponse) FromCommand(command *userlogin.UserLoginCommandResponse) {
	u.AccessToken = command.AccessToken
	u.RefreshToken = command.RefreshToken
}

type TokenRefreshResponse struct {
	AccessToken string `json:"access_token"`
}

func (t *TokenRefreshResponse) FromCommand(command *tokenrefresh.TokenRefreshCommandResponse) {
	t.AccessToken = command.AccessToken
}

type PasswordResetRequestResponse struct {
	Message string `json:"message"`
}

func (p *PasswordResetRequestResponse) FromCommand(command *passwordresetrequest.PasswordResetRequestCommandResponse) {
	p.Message = command.Message
}

type PasswordResetResponse struct {
	Message string `json:"message"`
}

func (p *PasswordResetResponse) FromCommand(command *passwordreset.PasswordResetCommandResponse) {
	p.Message = command.Message
}

type EmailVerifyResponse struct {
	Message string `json:"message"`
}

func (e *EmailVerifyResponse) FromCommand(command *emailverify.EmailVerifyCommandResponse) {
	e.Message = command.Message
}

type PasswordSetResponse struct {
	Message string `json:"message"`
}

func (p *PasswordSetResponse) FromCommand(command *passwordset.PasswordSetCommandResponse) {
	p.Message = command.Message
}
