package handler

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/mehdihadeli/go-mediatr"
	"github.com/zelalem-12/bill-aggregation-system_onetab/bill-service/internal/adapter/http/request"
	"github.com/zelalem-12/bill-aggregation-system_onetab/bill-service/internal/adapter/http/response"

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
	"github.com/zelalem-12/bill-aggregation-system_onetab/bill-service/internal/util"
)

type BillHandler struct {
	Handler
}

func NewBillHandler() *BillHandler {
	return &BillHandler{
		Handler: NewHandler(),
	}
}

// GetAggregatedBills godoc
// @Summary Fetch aggregated bills for a user
// @Description Retrieve all bills for a user, aggregated from all linked providers.
// @Tags Bills
// @Accept json
// @Produce json
// @Param user_id path string true "User ID"
// @Success 200 {object} response.GetAggregatedBillsResponse
// @Failure 400 {object} response.ErrorResponse
// @Router /bills [get]
func (handler *BillHandler) GetAggregatedBillsHandler(c echo.Context) error {

	user := c.Get("user").(util.BasicUserInfo)

	responseDTO := &response.GetAggregatedBillsResponse{}

	query := aggregatedbills.GetAggregatedBillsQuery{
		UserID: user.UserID,
	}

	if err := query.Validate(); err != nil {
		return echo.ErrBadRequest
	}

	result, err := mediatr.Send[*aggregatedbills.GetAggregatedBillsQuery, *aggregatedbills.GetAggregatedBillsQueryResponse](context.Background(), &query)
	if err != nil {
		return err
	}

	responseDTO.FromQueryResponse(result)

	return c.JSON(http.StatusOK, responseDTO)
}

// GetBillsByProvider godoc
// @Summary Fetch all bills from a specific provider
// @Description Fetch all bills from a specific provider
// @Tags Bills
// @Accept json
// @Produce json
// @Param provider_name path string true "Provider Name"
// @Success 200 {object} response.GetBillsByProviderResponse
// @Failure 400 {object} response.ErrorResponse
// @Router /bills/{provider_name} [get]
func (handler *BillHandler) GetBillsByProviderHandler(c echo.Context) error {

	user := c.Get("user").(util.BasicUserInfo)

	requestDTO := &request.GetBillsByProviderRequest{}
	responseDTO := response.GetBillsByProviderResponse{}

	if err := handler.BindAndValidate(c, requestDTO); err != nil {
		return echo.ErrBadRequest
	}
	if err := requestDTO.Validate(); err != nil {
		return echo.ErrBadRequest
	}

	query := billsbyprovider.GetBillsByProviderQuery{
		UserID:       user.UserID,
		ProviderName: &requestDTO.ProviderName,
	}

	result, err := mediatr.Send[*billsbyprovider.GetBillsByProviderQuery, *billsbyprovider.GetBillsByProviderQueryResponse](context.Background(), &query)
	if err != nil {
		return err
	}

	responseDTO.FromQueryResponse(result)

	return c.JSON(http.StatusOK, responseDTO)
}

// GetBillsByProviderID godoc
// @Summary Fetch all bills from a specific provider
// @Description Fetch all bills from a specific provider
// @Tags Bills
// @Accept json
// @Produce json
// @Param provider_name path string true "Provider ID"
// @Success 200 {object} response.GetBillsByProviderResponse
// @Failure 400 {object} response.ErrorResponse
// @Router /bills/provider/{provider_Id} [get]
func (handler *BillHandler) GetBillsByProviderIdHandler(c echo.Context) error {

	user := c.Get("user").(util.BasicUserInfo)

	requestDTO := &request.GetBillsByProviderIdRequest{}
	responseDTO := response.GetBillsByProviderResponse{}

	if err := handler.BindAndValidate(c, requestDTO); err != nil {
		return echo.ErrBadRequest
	}
	if err := requestDTO.Validate(); err != nil {
		return echo.ErrBadRequest
	}

	providerID, err := util.ToUUID(requestDTO.ProviderId)
	if err != nil {
		return echo.ErrBadRequest
	}

	query := billsbyprovider.GetBillsByProviderQuery{
		UserID:     user.UserID,
		ProviderID: &providerID,
	}

	result, err := mediatr.Send[*billsbyprovider.GetBillsByProviderQuery, *billsbyprovider.GetBillsByProviderQueryResponse](context.Background(), &query)
	if err != nil {
		return err
	}

	responseDTO.FromQueryResponse(result)

	return c.JSON(http.StatusOK, responseDTO)
}

// MarkBillAsPaid godoc
// @Summary Mark a bill as paid
// @Description Marks a bill as paid in the system
// @Tags Bills
// @Accept json
// @Produce json
// @Param bill_id path string true "Bill ID"
// @Success 200 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Router /bills/{bill_id}/pay [patch]
func (handler *BillHandler) MarkBillAsPaidHandler(c echo.Context) error {

	requestDTO := &request.MarkBillAsPaidRequest{}
	if err := handler.BindAndValidate(c, requestDTO); err != nil {
		return echo.ErrBadRequest
	}

	query := markbillaspaid.MarkBillAsPaidQuery{
		BillID: requestDTO.BillID,
	}

	result, err := mediatr.Send[*markbillaspaid.MarkBillAsPaidQuery, *markbillaspaid.MarkBillAsPaidQueryResponse](context.Background(), &query)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, response.SuccessResponse{Message: result.Message})
}

// DeleteBill godoc
// @Summary Delete a bill
// @Description Removes a bill record (if user deletes a linked provider, associated bills may also be deleted)
// @Tags Bills
// @Accept json
// @Produce json
// @Param bill_id path string true "Bill ID"
// @Success 200 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Router /bills/{bill_id} [delete]
func (handler *BillHandler) DeleteBillHandler(c echo.Context) error {
	requestDTO := &request.DeleteBillRequest{}
	if err := handler.BindAndValidate(c, requestDTO); err != nil {
		return echo.ErrBadRequest
	}

	query := deletebill.DeleteBillQuery{
		BillID: requestDTO.BillID,
	}

	result, err := mediatr.Send[*deletebill.DeleteBillQuery, *deletebill.DeleteBillQueryResponse](context.Background(), &query)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, response.SuccessResponse{Message: result.Message})
}

// CreateBill godoc
// @Summary Create a new bill
// @Description Creates a new bill record.
// @Tags Bills
// @Accept json
// @Produce json
// @Param bill body request.CreateBillRequest true "Bill details"
// @Success 200 {object} response.CreateBillResponse
// @Failure 400 {object} response.ErrorResponse
// @Router /bills [post]
func (handler *BillHandler) CreateBillHandler(c echo.Context) error {

	reqDTO := &request.CreateBillRequest{}
	if err := handler.BindAndValidate(c, reqDTO); err != nil {
		log.Error("Error binding request", err)
		return echo.ErrBadRequest
	}

	cmd, err := reqDTO.ToCommand()
	if err != nil {
		return echo.ErrBadRequest
	}

	cmdResp, err := mediatr.Send[*createbill.CreateBillCommand, *createbill.CreateBillCommandResponse](context.Background(), cmd)
	if err != nil {
		return err
	}

	responseDTO := response.CreateBillResponse{
		ID:      cmdResp.BillID,
		Message: "Bill created successfully",
	}

	return c.JSON(http.StatusOK, responseDTO)
}

func (handler *BillHandler) DeleteBillsByProvider(c echo.Context) error {

	reqDTO := &request.DeleteBillsByProviderRequest{}
	if err := handler.BindAndValidate(c, reqDTO); err != nil {
		return echo.ErrBadRequest
	}

	cmd := reqDTO.ToCommand()

	cmdResp, err := mediatr.Send[*deletebillsbyprovider.DeleteBillsByProviderCommand, *deletebillsbyprovider.DeleteBillsByProviderCommandResponse](context.Background(), cmd)
	if err != nil {
		return err
	}

	resDTO := response.DeleteBillsByProviderResponse{}
	resDTO.FromCommandResponse(cmdResp)

	return c.JSON(http.StatusOK, resDTO)
}

// GetOverdueBills godoc
// @Summary Get overdue bills for a user
// @Description Retrieves all overdue bills for a user along with amount due and due dates.
// @Tags Bills
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} response.GetOverdueBillsResponse
// @Failure 400 {object} response.ErrorResponse
// @Router /bills/overdue [get]
func (handler *BillHandler) GetOverdueBillsHandler(c echo.Context) error {

	user := c.Get("user").(util.BasicUserInfo)
	resDTO := response.GetOverdueBillsResponse{}

	query := overduebills.GetOverdueBillsQuery{UserID: user.UserID}

	result, err := mediatr.Send[*overduebills.GetOverdueBillsQuery, *overduebills.GetOverdueBillsQueryResponse](context.Background(), &query)
	if err != nil {
		return err
	}

	resDTO.FromQueryResponse(result)

	return c.JSON(http.StatusOK, resDTO)
}

// GetCategorySpending godoc
// @Summary Get category spending insights for a user
// @Description Retrieves insights into the total amount spent in each bill category.
// @Tags Bills
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} response.GetCategorySpendingResponse
// @Failure 400 {object} response.ErrorResponse
// @Router /bills/categories [get]
func (handler *BillHandler) GetCategorySpendingHandler(c echo.Context) error {

	user := c.Get("user").(util.BasicUserInfo)
	resDTO := response.GetCategorySpendingResponse{}

	query := billcategories.GetCategorySpendingQuery{UserID: user.UserID}
	result, err := mediatr.Send[*billcategories.GetCategorySpendingQuery, *billcategories.GetCategorySpendingQueryResponse](context.Background(), &query)
	if err != nil {
		return err
	}

	resDTO.FromQueryResponse(result)
	return c.JSON(http.StatusOK, resDTO)
}

// GetBillPaymentHistory godoc
// @Summary Get bill payment history
// @Description Retrieves the payment history of bills for a user, including paid date and due date.
// @Tags Bills
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Success 200 {array} response.PaymentHistoryResponse
// @Failure 400 {object} response.ErrorResponse
// @Router /bills/history [get]
func (handler *BillHandler) GetBillPaymentHistoryHandler(c echo.Context) error {

	user := c.Get("user").(util.BasicUserInfo)
	resDTO := response.GetBillPaymentHistoryResponse{}

	query := billpaymenthistory.GetBillPaymentHistoryQuery{UserID: user.UserID}
	result, err := mediatr.Send[*billpaymenthistory.GetBillPaymentHistoryQuery, *billpaymenthistory.GetBillPaymentHistoryQueryResponse](context.Background(), &query)
	if err != nil {
		return err
	}

	resDTO.FromQueryResponse(result)
	return c.JSON(http.StatusOK, resDTO)
}

// HandlerGetBillSummary godoc
// @Summary Get bill summary for a user
// @Description Retrieves total outstanding, total paid, and overall due amounts for the user.
// @Tags Bills
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} response.GetBillSummaryResponse
// @Failure 400 {object} response.ErrorResponse
// @Router /bills/summary [get]
func (handler *BillHandler) HandlerGetBillSummary(c echo.Context) error {
	user := c.Get("user").(util.BasicUserInfo)
	resDTO := response.GetBillSummaryResponse{}

	query := billsummary.GetBillSummaryQuery{UserID: user.UserID}
	result, err := mediatr.Send[*billsummary.GetBillSummaryQuery, *billsummary.GetBillSummaryQueryResponse](context.Background(), &query)
	if err != nil {
		return err
	}

	resDTO.FromQueryResponse(result)
	return c.JSON(http.StatusOK, resDTO)
}

// HandlerGetMonthlySpendingTrends godoc
// @Summary Get monthly spending trends for a user
// @Description Retrieves spending trends over time, grouped by month.
// @Tags Bills
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} response.GetMonthlySpendingTrendsResponse
// @Failure 400 {object} response.ErrorResponse
// @Router /bills/summary/trends [get]
func (handler *BillHandler) HandlerGetMonthlySpendingTrends(c echo.Context) error {
	user := c.Get("user").(util.BasicUserInfo)
	resDTO := response.GetMonthlySpendingTrendsResponse{}

	query := spendingtrends.GetMonthlySpendingTrendsQuery{UserID: user.UserID}
	result, err := mediatr.Send[*spendingtrends.GetMonthlySpendingTrendsQuery, *spendingtrends.GetMonthlySpendingTrendsQueryResponse](context.Background(), &query)
	if err != nil {
		return err
	}
	resDTO.FromQueryResponse(result)
	return c.JSON(http.StatusOK, resDTO)
}
