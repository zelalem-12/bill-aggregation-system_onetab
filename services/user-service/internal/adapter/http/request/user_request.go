package request

import (
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/app/command/currentuserupdate"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/app/command/passwordchange"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/util"
)

type ChangePasswordRequest struct {
	OldPassword     string `json:"old_password" validate:"required"`
	NewPassword     string `json:"new_password" validate:"required"`
	ConfirmPassword string `json:"confirm_password" validate:"required"`
}

func (r *ChangePasswordRequest) Validate() error {

	if r.NewPassword != r.ConfirmPassword {
		return util.CreateError("new password and confirm password does not match")
	}

	return validate.Struct(r)
}
func (r *ChangePasswordRequest) ToCommand() *passwordchange.PasswordChangeCommand {
	return &passwordchange.PasswordChangeCommand{
		OldPassword: r.OldPassword,
		NewPassword: r.NewPassword,
	}
}

type UserUpdateRequest struct {
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	ProfilePicture string `json:"profile_picture"`
}

func (r *UserUpdateRequest) ToCommand() *currentuserupdate.CurrentUserUpdateCommand {
	return &currentuserupdate.CurrentUserUpdateCommand{
		FirstName:      r.FirstName,
		LastName:       r.LastName,
		ProfilePicture: r.ProfilePicture,
	}
}
