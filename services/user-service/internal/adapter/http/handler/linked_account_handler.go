package handler

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mehdihadeli/go-mediatr"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/adapter/http/request"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/adapter/http/response"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/app/command/linkaccount"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/app/command/unlinkaccount"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/util"
)

type LinkedAccountHandler struct {
	Handler
}

func NewLinkedAccountHandler() *LinkedAccountHandler {
	return &LinkedAccountHandler{
		Handler: NewHandler(),
	}
}

// LinkAccountHandler godoc
// @Summary Link a utility account
// @Description Links a utility account to the user’s profile.
// @Tags Accounts
// @Accept json
// @Produce json
// @Param body body request.LinkAccountRequest true "Utility account details"
// @Success 201 {object} response.LinkAccountResponse
// @Failure 400 {object} response.ErrorResponse
// @Router /accounts/link [post]
func (handler *LinkedAccountHandler) LinkAccountHandler(c echo.Context) error {

	user := c.Get("user").(util.BasicUserInfo)

	linkAccountRequest := &request.LinkAccountRequest{}
	linkAccountResponse := &response.LinkAccountResponse{}

	if err := handler.BindAndValidate(c, linkAccountRequest); err != nil {
		return echo.ErrBadRequest
	}

	command := linkAccountRequest.ToCommand()
	command.UserID = user.UserID

	if err := command.Validate(); err != nil {
		return echo.ErrBadRequest
	}

	result, err := mediatr.Send[*linkaccount.LinkAccountCommand, *linkaccount.LinkAccountCommandResponse](context.Background(), command)
	if err != nil {
		return err
	}

	linkAccountResponse.FromCommand(result)

	return c.JSON(http.StatusCreated, linkAccountResponse)
}

// UnlinkAccountHandler godoc
// @Summary Unlink a utility account
// @Description Removes a linked utility account from the user’s profile.
// @Tags Accounts
// @Accept json
// @Produce json
// @Param account_id path string true "Linked account ID"
// @Success 200 {object} response.UnlinkAccountResponse
// @Failure 400 {object} response.ErrorResponse
// @Router /accounts/{account_id} [delete]
func (handler *LinkedAccountHandler) UnlinkAccountHandler(c echo.Context) error {

	request := &request.UnlinkAccountRequest{}
	if err := handler.BindAndValidate(c, request); err != nil {
		return err
	}

	if err := request.Validate(); err != nil {
		return err
	}

	command := &unlinkaccount.UnlinkAccountCommand{
		AccountID: request.AccountID,
	}

	result, err := mediatr.Send[*unlinkaccount.UnlinkAccountCommand, *unlinkaccount.UnlinkAccountCommandResponse](context.Background(), command)
	if err != nil {
		return err
	}

	response := response.UnlinkAccountResponse{}
	response.FromCommand(result)

	return c.JSON(http.StatusOK, response)
}
