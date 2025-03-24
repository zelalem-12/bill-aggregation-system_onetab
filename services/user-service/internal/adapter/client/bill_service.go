package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/gommon/log"
	billClientPort "github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/app/client"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/infrastructure/config"
)

type BillServiceClient struct {
	cfg        *config.Config
	HTTPClient *http.Client
}

func NewBillServiceClient(cfg *config.Config) billClientPort.BillServiceClient {
	return &BillServiceClient{
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
		cfg: cfg,
	}
}

func (c *BillServiceClient) RemoveUnlinkedProviderBills(userId, providerId uuid.UUID) (*billClientPort.RemoveBillsResponse, error) {
	url := fmt.Sprintf("%s/internal/bills/%s", c.cfg.BILL_BASE_URL, providerId.String())

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("user_id", userId.String())

	log.Info("Sending request to remove unlinked provider bills", req)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		log.Error("Failed to remove unlinked provider bills", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to remove unlinked provider bills")
	}

	var removeBillsResponse billClientPort.RemoveBillsResponse
	if err := json.NewDecoder(resp.Body).Decode(&removeBillsResponse); err != nil {
		return nil, err
	}

	log.Info("Successfully removed unlinked provider bills", removeBillsResponse)

	return &removeBillsResponse, nil
}
