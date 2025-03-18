package response

import (
	"github.com/google/uuid"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/app/command/currentuserdelete"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/app/command/currentuserupdate"
)

type UserResponse struct {
	ID             uuid.UUID `json:"id"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	Email          string    `json:"email"`
	IsVerified     bool      `json:"is_verified"`
	ProfilePicture string    `json:"profile_picture"`
}

type PasswordChangedResponse struct {
	Message string `json:"message"`
}

func NewChangePasswordResponse(message string) *PasswordChangedResponse {
	return &PasswordChangedResponse{Message: message}
}

type UserUpdateResponse struct {
	Message string `json:"message"`
}

func NewUserUpdateResponse(commandResponse *currentuserupdate.CurrentUserUpdateCommandResponse) *UserUpdateResponse {
	return &UserUpdateResponse{
		Message: commandResponse.Message,
	}
}

type UserDeleteResponse struct {
	Message string `json:"message"`
}

func NewUserDeleteResponse(commandResponse *currentuserdelete.CurrentUserDeleteCommandResponse) *UserDeleteResponse {
	return &UserDeleteResponse{
		Message: commandResponse.Message,
	}
}
