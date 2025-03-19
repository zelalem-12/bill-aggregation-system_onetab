package handler

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mehdihadeli/go-mediatr"
	"github.com/zelalem-12/bill-aggregation-system_onetab/provider-service/internal/adapter/http/request"
	"github.com/zelalem-12/bill-aggregation-system_onetab/provider-service/internal/adapter/http/response"
	"github.com/zelalem-12/bill-aggregation-system_onetab/provider-service/internal/app/query/providerbyid"
	"github.com/zelalem-12/bill-aggregation-system_onetab/provider-service/internal/app/query/providerbyname"
	"github.com/zelalem-12/bill-aggregation-system_onetab/provider-service/internal/app/query/providers"
)

type ProviderHandler struct {
	Handler
}

func NewProviderHandler() *ProviderHandler {
	return &ProviderHandler{
		Handler: NewHandler(),
	}
}

// GetProviderByID godoc
// @Summary Get a provider by ID
// @Description Fetches provider details using the provider ID.
// @Tags Providers
// @Accept json
// @Produce json
// @Param provider_id path string true "Provider ID"
// @Success 200 {object} response.GetProviderByIDResponse
// @Failure 400 {object} response.ErrorResponse
// @Router /providers/{provider_id} [get]
func (handler *ProviderHandler) GetProviderByIdHandler(c echo.Context) error {

	requestDTO := &request.GetProviderByIDRequest{}
	if err := handler.BindAndValidate(c, requestDTO); err != nil {
		return echo.ErrBadRequest
	}

	query := requestDTO.ToQuery()

	if err := query.Validate(); err != nil {
		return echo.ErrBadRequest
	}

	result, err := mediatr.Send[*providerbyid.GetProviderByIDQuery, *providerbyid.GetProviderByIDQueryResponse](context.Background(), query)
	if err != nil {
		return err
	}

	responseDTO := response.GetProviderByIDResponse{}
	responseDTO.FromQueryResponse(result)

	return c.JSON(http.StatusOK, responseDTO)
}

// GetProviderByName godoc
// @Summary Get a provider by name
// @Description Fetches provider details using the provider name.
// @Tags Providers
// @Accept json
// @Produce json
// @Param provider_name path string true "Provider Name"
// @Success 200 {object} response.GetProviderByNameResponse
// @Failure 400 {object} response.ErrorResponse
// @Router /providers/name/{provider_name} [get]
func (handler *ProviderHandler) GetProviderByNameHandler(c echo.Context) error {

	requestDTO := &request.GetProviderByNameRequest{}
	if err := handler.BindAndValidate(c, requestDTO); err != nil {
		return echo.ErrBadRequest
	}

	query := requestDTO.ToQuery()

	if err := query.Validate(); err != nil {
		return echo.ErrBadRequest
	}

	result, err := mediatr.Send[*providerbyname.GetProviderByNameQuery, *providerbyname.GetProviderByNameQueryResponse](context.Background(), query)
	if err != nil {
		return err
	}

	responseDTO := response.GetProviderByNameResponse{}
	responseDTO.FromQueryResponse(result)

	return c.JSON(http.StatusOK, responseDTO)
}

// GetProvidersHandler godoc
// @Summary Get all providers
// @Description Retrieves a list of all available providers.
// @Tags Providers
// @Accept json
// @Produce json
// @Success 200 {object} response.GetProvidersResponse
// @Failure 400 {object} response.ErrorResponse
// @Router /providers [get]
func (handler *ProviderHandler) GetProvidersHandler(c echo.Context) error {
	query := &providers.GetProvidersQuery{}

	result, err := mediatr.Send[*providers.GetProvidersQuery, *providers.GetProvidersQueryResponse](context.Background(), query)
	if err != nil {
		return err
	}

	providerResponse := response.GetProvidersResponse{}
	providerResponse.FromQueryResponse(result)

	return c.JSON(http.StatusOK, providerResponse)
}
