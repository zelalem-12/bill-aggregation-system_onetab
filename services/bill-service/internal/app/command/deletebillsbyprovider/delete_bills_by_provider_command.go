package deletebillsbyprovider

import "github.com/google/uuid"

type DeleteBillsByProviderCommand struct {
	UserID     uuid.UUID
	ProviderID uuid.UUID
}

type DeleteBillsByProviderCommandResponse struct {
	Message string
}
