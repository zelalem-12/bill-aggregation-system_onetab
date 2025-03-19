package billsbyprovider

import (
	"context"

	"github.com/zelalem-12/bill-aggregation-system_onetab/bill-service/internal/app/repo"
	"github.com/zelalem-12/bill-aggregation-system_onetab/bill-service/internal/app/service"
	"github.com/zelalem-12/bill-aggregation-system_onetab/bill-service/internal/domain"
	"github.com/zelalem-12/bill-aggregation-system_onetab/bill-service/internal/util"
)

type GetBillsByProviderQueryHandler struct {
	BillRepo repo.BillRepo
}

func NewGetBillsByProviderQueryHandler(billRepo repo.BillRepo) *GetBillsByProviderQueryHandler {
	return &GetBillsByProviderQueryHandler{
		BillRepo: billRepo,
	}
}

func (h *GetBillsByProviderQueryHandler) Handle(ctx context.Context, query *GetBillsByProviderQuery) (*GetBillsByProviderQueryResponse, error) {

	var result []*domain.Bill

	if query.ProviderID != nil {
		bills, err := h.BillRepo.FindByProviderID(ctx, query.UserID, *query.ProviderID)
		if err != nil {
			return nil, err
		}
		result = bills

	} else if query.ProviderName != nil {
		bills, err := h.BillRepo.FindByProviderName(ctx, query.UserID, *query.ProviderName)
		if err != nil {
			return nil, err
		}
		result = bills
	} else {
		return nil, util.CreateError("providerID or ProviderName is required")
	}

	var providerName string
	var billResponses []Bill

	for _, bill := range result {

		providerName = bill.GetProviderName()
		billID, err := service.ToUUID(bill.GetID())
		if err != nil {
			return nil, err
		}

		billResponses = append(billResponses, Bill{
			ID:      billID,
			Amount:  bill.GetAmount(),
			DueDate: bill.GetDueDate().Format("2006-01-02"),
			Status:  bill.GetStatus().String(),
		})
	}

	return &GetBillsByProviderQueryResponse{
		Provider: providerName,
		Bills:    billResponses,
	}, nil
}
