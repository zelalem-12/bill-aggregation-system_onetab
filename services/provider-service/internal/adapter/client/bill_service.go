package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/gommon/log"
	billClientPort "github.com/zelalem-12/bill-aggregation-system_onetab/provider-service/internal/app/client"
	"github.com/zelalem-12/bill-aggregation-system_onetab/provider-service/internal/infrastructure/config"
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

func (c *BillServiceClient) CreateBill(userId uuid.UUID, bill *billClientPort.CreateBillRequestDTO) (*billClientPort.CreateBillResponse, error) {
	url := fmt.Sprintf("%s/internal/bills", c.cfg.BILL_BASE_URL)

	body, err := json.Marshal(bill)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("user_id", userId.String())

	log.Info("Sending request to create a new bill", req)

	resp, err := c.HTTPClient.Do(req)
	if err != nil || resp.StatusCode != http.StatusOK {
		log.Error("Failed to create bill", err)
		return nil, err
	}
	defer resp.Body.Close()

	var createBillResponse billClientPort.CreateBillResponse
	if err := json.NewDecoder(resp.Body).Decode(&createBillResponse); err != nil {
		return nil, err
	}

	log.Info("Successfully created bill", createBillResponse)

	return &createBillResponse, nil
}
