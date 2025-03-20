package app

import (
	"github.com/mehdihadeli/go-mediatr"
	"github.com/zelalem-12/bill-aggregation-system_onetab/bill-service/internal/app/command/createbill"
	"github.com/zelalem-12/bill-aggregation-system_onetab/bill-service/internal/app/command/deletebill"
	"github.com/zelalem-12/bill-aggregation-system_onetab/bill-service/internal/app/command/deletebillsbyprovider"
	"github.com/zelalem-12/bill-aggregation-system_onetab/bill-service/internal/app/command/markbillaspaid"
	"github.com/zelalem-12/bill-aggregation-system_onetab/bill-service/internal/app/query/aggregatedbills"
	"github.com/zelalem-12/bill-aggregation-system_onetab/bill-service/internal/app/query/billcategories"
	"github.com/zelalem-12/bill-aggregation-system_onetab/bill-service/internal/app/query/billpaymenthistory"
	"github.com/zelalem-12/bill-aggregation-system_onetab/bill-service/internal/app/query/billsbyprovider"
	"github.com/zelalem-12/bill-aggregation-system_onetab/bill-service/internal/app/query/billsummary"
	"github.com/zelalem-12/bill-aggregation-system_onetab/bill-service/internal/app/query/overduebills"
	"github.com/zelalem-12/bill-aggregation-system_onetab/bill-service/internal/app/query/spendingtrends"
	"github.com/zelalem-12/bill-aggregation-system_onetab/bill-service/internal/app/repo"
	"github.com/zelalem-12/bill-aggregation-system_onetab/bill-service/internal/infrastructure/config"
)

func RegisterCQRSHandlers(
	cfg *config.Config,
	billRepo repo.BillRepo,
) error {

	markBillAsPaidQueryHandler := markbillaspaid.NewMarkBillAsPaidQueryHandler(billRepo)
	deleteBillQueryHandler := deletebill.NewDeleteBillQueryHandler(billRepo)
	createBillCommandHandler := createbill.NewCreateBillCommandHandler(billRepo)
	deleteBillsByProviderCommandHandler := deletebillsbyprovider.NewDeleteBillsByProviderCommandHandler(billRepo)

	aggregatedBillsQueryHandler := aggregatedbills.NewGetAggregatedBillsQueryHandler(billRepo)
	billsByProviderQueryHandler := billsbyprovider.NewGetBillsByProviderQueryHandler(billRepo)
	overdueBillsQueryHandler := overduebills.NewGetOverdueBillsQueryHandler(billRepo)
	categorySpendingQueryHandler := billcategories.NewGetCategorySpendingQueryHandler(billRepo)
	paymentHistoryQueryHandler := billpaymenthistory.NewGetBillPaymentHistoryQueryHandler(billRepo)
	billSummaryQueryHandler := billsummary.NewGetBillSummaryQueryHandler(billRepo)
	monthlySpendingTrendsQueryHandler := spendingtrends.NewGetMonthlySpendingTrendsQueryHandler(billRepo)

	// Commands
	if err := mediatr.RegisterRequestHandler(markBillAsPaidQueryHandler); err != nil {
		return err
	}
	if err := mediatr.RegisterRequestHandler(deleteBillQueryHandler); err != nil {
		return err
	}
	if err := mediatr.RegisterRequestHandler(createBillCommandHandler); err != nil {
		return err
	}
	if err := mediatr.RegisterRequestHandler(deleteBillsByProviderCommandHandler); err != nil {
		return err
	}

	// Queries

	if err := mediatr.RegisterRequestHandler(aggregatedBillsQueryHandler); err != nil {
		return err
	}
	if err := mediatr.RegisterRequestHandler(billsByProviderQueryHandler); err != nil {
		return err
	}
	if err := mediatr.RegisterRequestHandler(overdueBillsQueryHandler); err != nil {
		return err
	}
	if err := mediatr.RegisterRequestHandler(categorySpendingQueryHandler); err != nil {
		return err
	}
	if err := mediatr.RegisterRequestHandler(paymentHistoryQueryHandler); err != nil {
		return err
	}
	if err := mediatr.RegisterRequestHandler(billSummaryQueryHandler); err != nil {
		return err
	}
	if err := mediatr.RegisterRequestHandler(monthlySpendingTrendsQueryHandler); err != nil {
		return err
	}

	return nil
}
