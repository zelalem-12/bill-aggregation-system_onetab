package usersignup

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/app/client"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/app/repo"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/app/service"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/domain"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/infrastructure/config"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/util"
)

type UserSignupCommandHandler struct {
	config       *config.Config
	userRepo     repo.UserRepo
	tokenRepo    repo.TokenRepo
	emailService client.EmailServicePort
}

func NewUserSignupCommandHandler(
	config *config.Config,
	userRepo repo.UserRepo,
	tokenRepo repo.TokenRepo,
	emailService client.EmailServicePort,
) *UserSignupCommandHandler {
	return &UserSignupCommandHandler{
		config:       config,
		userRepo:     userRepo,
		tokenRepo:    tokenRepo,
		emailService: emailService,
	}
}

func (h *UserSignupCommandHandler) Handle(ctx context.Context, command *UserSignupCommand) (*UserSignupCommandResponse, error) {

	createdUser, err := h.userRepo.Save(ctx, domain.NewUser(command.FirstName, command.LastName, command.Email))
	if err != nil {
		fmt.Printf("UserSignupCommandHandler : Error saving user: %v", err)
		return nil, err
	}

	userID, err := uuid.Parse(createdUser.GetID())
	if err != nil {
		fmt.Printf("UserSignupCommandHandler : Error parsing user id: %v", err)
		return nil, err
	}

	go func() {
		emailVerificationToken, err := util.GenerateNonAccessToken(h.config.ACCESS_TOKEN_KEY, userID, h.config.ACCESS_TOKEN_EXPIRY)
		if err != nil {
			fmt.Printf("Error Generating verification token :%v", err)
		}

		//TODO use  cache only the reset token instead of saving it to the database
		_, err = h.tokenRepo.Save(ctx, domain.NewToken(emailVerificationToken, createdUser.GetID()))
		if err != nil {
			fmt.Printf("Error saving verification token :%v", err)
		}

		contentData := service.NewEmailContentData(
			"Email Verification",
			"Email Verification",
			createdUser.GetFirstName(),
			createdUser.GetLastName(),
			fmt.Sprintf("%s/verify-email?token=%s", h.config.FRONTEND_URL, emailVerificationToken),
		)

		err = h.emailService.SendEmail(
			h.config.SENDER_EMAIL,
			[]string{createdUser.GetEmail()}, nil, nil,
			"Verify Your Email",
			"email_verification_request",
			contentData,
		)
		if err != nil {
			fmt.Printf("Error sending email verificaiton email: %v", err)
		}
	}()

	return &UserSignupCommandResponse{
		UserID:  userID,
		Message: "User created successfully, please verify your email",
	}, nil
}
