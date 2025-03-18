package handler

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/mehdihadeli/go-mediatr"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/adapter/http/request"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/adapter/http/response"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/app/command/currentuserdelete"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/app/command/currentuserupdate"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/app/command/passwordchange"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/app/query/currentuser"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/util"
)

type UserHandler struct {
	Handler
}

func NewUserHandler() *UserHandler {
	return &UserHandler{
		Handler: NewHandler(),
	}
}

// ChangePasswordHandler godoc
// @Summary Change user password
// @Description This endpoint allows the user to change their password. The user needs to provide the old password, new password, and a confirmation of the new password.
// @Tags User
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param body body request.ChangePasswordRequest true "Change password request details"
// @Success 200 {object} response.PasswordChangedResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /users/change-password [post]
func (handler *UserHandler) ChangePasswordHandler(c echo.Context) error {

	changePasswordRequest := &request.ChangePasswordRequest{}

	if err := handler.BindAndValidate(c, changePasswordRequest); err != nil {
		return echo.ErrBadRequest
	}

	if err := changePasswordRequest.Validate(); err != nil {
		return echo.ErrBadRequest
	}

	user := c.Get("user").(util.BasicUserInfo)

	command := passwordchange.PasswordChangeCommand{
		UserID:      user.UserID,
		OldPassword: changePasswordRequest.OldPassword,
		NewPassword: changePasswordRequest.NewPassword,
	}

	result, err := mediatr.Send[*passwordchange.PasswordChangeCommand, *passwordchange.PasswordChangeCommandResponse](context.Background(), &command)
	if err != nil {
		return err
	}

	return c.JSON(200, response.NewChangePasswordResponse(result.Message))
}

// GetCurrentUserHandler godoc
// @Summary Get the current authenticated user
// @Description This endpoint returns the details of the currently authenticated user.
// @Tags User
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} response.UserResponse
// @Failure 401 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /users/current [get]
func (handler *UserHandler) GetCurrentUserHandler(c echo.Context) error {
	user := c.Get("user").(util.BasicUserInfo)

	currentUserQuery := currentuser.CurrentUserQuery{
		UserID: user.UserID,
	}

	result, err := mediatr.Send[*currentuser.CurrentUserQuery, *currentuser.CurrentUserQueryResponse](context.Background(), &currentUserQuery)
	if err != nil {
		return err
	}

	response := response.UserResponse{
		ID:             result.ID,
		FirstName:      result.FirstName,
		LastName:       result.LastName,
		Email:          result.Email,
		IsVerified:     result.IsVerified,
		ProfilePicture: result.ProfilePicture,
	}

	return c.JSON(200, response)
}

// UpdateCurrentUserHandler godoc
// @Summary Update current user profile
// @Description This endpoint allows users to update their profile information such as first name, last name, profile picture.
// @Tags User
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param body body request.UserUpdateRequest true "Update user profile information"
// @Success 200 {object} response.UserUpdateResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /users/update [put]
func (handler *UserHandler) UpdateCurrentUserHandler(c echo.Context) error {

	user := c.Get("user").(util.BasicUserInfo)
	request := &request.UserUpdateRequest{}

	if err := handler.BindAndValidate(c, request); err != nil {
		return err
	}

	command := request.ToCommand()
	command.UserID = user.UserID

	result, err := mediatr.Send[*currentuserupdate.CurrentUserUpdateCommand, *currentuserupdate.CurrentUserUpdateCommandResponse](context.Background(), command)
	if err != nil {
		return err
	}
	return c.JSON(200, response.NewUserUpdateResponse(result))
}

// DeleteCurrentUserHandler godoc
// @Summary Delete current user account
// @Description This endpoint allows the user to delete their account permanently.
// @Tags User
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} response.UserDeleteResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /users/delete [delete]
func (handler *UserHandler) DeleteCurrentUserHandler(c echo.Context) error {
	user := c.Get("user").(util.BasicUserInfo)

	command := currentuserdelete.CurrentUserDeleteCommand{
		UserID: user.UserID,
	}

	result, err := mediatr.Send[*currentuserdelete.CurrentUserDeleteCommand, *currentuserdelete.CurrentUserDeleteCommandResponse](context.Background(), &command)
	if err != nil {
		return err
	}

	return c.JSON(200, response.NewUserDeleteResponse(result))
}
