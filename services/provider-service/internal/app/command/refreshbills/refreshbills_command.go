package refreshbills

import "github.com/google/uuid"

type RefreshBillsCommand struct {
	UserID uuid.UUID `json:"user_id"`
}

type RefreshBillsCommandResponse struct {
	Message string `json:"message"`
}
