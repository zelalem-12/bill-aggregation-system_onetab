package client

import (
	"github.com/google/uuid"
)

type RemoveBillsResponse struct {
	Message string `json:"message"`
}

type BillServiceClient interface {
	RemoveUnlinkedProviderBills(userId, providerId uuid.UUID) (*RemoveBillsResponse, error)
}
