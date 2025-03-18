package app

import (
	"github.com/mehdihadeli/go-mediatr"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/app/client"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/app/command/emailverify"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/app/command/passwordchange"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/app/command/passwordreset"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/app/command/passwordresetrequest"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/app/command/passwordset"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/app/command/tokenrefresh"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/app/command/userlogin"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/app/command/usersignup"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/app/repo"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/infrastructure/config"
)

func RegisterCQRSHandlers(
	cfg *config.Config,
	emailService client.EmailServicePort,
	userRepo repo.UserRepo,
	tokenRepo repo.TokenRepo,

) error {

	emailVerifyCommandHandler := emailverify.NewEmailVerifyCommandHandler(cfg, userRepo, tokenRepo, emailService)
	passwordChangeCommandHandler := passwordchange.NewPasswordChangeCommandHandler(userRepo, tokenRepo)
	passwordresetCommandHandler := passwordreset.NewPasswordResetCommandHandler(cfg, userRepo, tokenRepo, emailService)
	passwordsetCommandHandler := passwordset.NewPasswordSetCommandHandler(cfg, userRepo, tokenRepo, emailService)
	passwordResetRequestCommandHandler := passwordresetrequest.NewPasswordResetRequestCommandHandler(cfg, userRepo, tokenRepo, emailService)
	refreshTokenCommandHandler := tokenrefresh.NewTokenRefreshCommandHandler(cfg, userRepo, tokenRepo)
	userLoginCommandHandler := userlogin.NewUserLoginCommandHandler(cfg, userRepo, tokenRepo)
	userSignupCommandHandler := usersignup.NewUserSignupCommandHandler(cfg, userRepo, tokenRepo, emailService)

	if err := mediatr.RegisterRequestHandler(emailVerifyCommandHandler); err != nil {
		return err
	}
	if err := mediatr.RegisterRequestHandler(passwordChangeCommandHandler); err != nil {
		return err
	}
	if err := mediatr.RegisterRequestHandler(passwordresetCommandHandler); err != nil {
		return err
	}
	if err := mediatr.RegisterRequestHandler(passwordsetCommandHandler); err != nil {
		return err
	}
	if err := mediatr.RegisterRequestHandler(passwordResetRequestCommandHandler); err != nil {
		return err
	}
	if err := mediatr.RegisterRequestHandler(refreshTokenCommandHandler); err != nil {
		return err
	}
	if err := mediatr.RegisterRequestHandler(userLoginCommandHandler); err != nil {
		return err
	}
	if err := mediatr.RegisterRequestHandler(userSignupCommandHandler); err != nil {
		return err
	}

	return nil
}
