package response

import (
	"github.com/google/uuid"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/app/command/linkaccount"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/app/command/unlinkaccount"
)

type LinkAccountResponse struct {
	AccountID uuid.UUID `json:"account_id"`
	Message   string    `json:"message"`
}

func (r *LinkAccountResponse) FromCommand(cmd *linkaccount.LinkAccountCommandResponse) {
	r.AccountID = cmd.AccountID
	r.Message = cmd.Message
}

type UnlinkAccountResponse struct {
	AccountID uuid.UUID `json:"account_id"`
	Message   string    `json:"message"`
}

// FromCommand maps command response to HTTP response
func (r *UnlinkAccountResponse) FromCommand(response *unlinkaccount.UnlinkAccountCommandResponse) {
	r.AccountID = response.AccountID
	r.Message = response.Message
}
