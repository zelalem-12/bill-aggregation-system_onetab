package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/gommon/log"
	userClientPort "github.com/zelalem-12/bill-aggregation-system_onetab/provider-service/internal/app/client"
	"github.com/zelalem-12/bill-aggregation-system_onetab/provider-service/internal/infrastructure/config"
)

type UserServiceClient struct {
	cfg        *config.Config
	HTTPClient *http.Client
}

func NewUserServiceClient(cfg *config.Config) userClientPort.UserServiceClient {
	return &UserServiceClient{
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
		cfg: cfg,
	}
}

func (c *UserServiceClient) GetUsers() ([]*userClientPort.User, error) {

	url := fmt.Sprintf("%s/internal/users", c.cfg.USER_BASE_URL)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	log.Info("Fetching users from user service", req)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		log.Error("Failed to fetch users", err)
		return nil, err
	}
	log.Info("Successfully fetched users from user service", resp)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to fetch users")
	}

	var users []*userClientPort.User
	if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
		return nil, err
	}

	log.Info("Successfully fetched users from user service", users)

	return users, nil
}

func (c *UserServiceClient) GetUserDetail(userID uuid.UUID) (*userClientPort.User, error) {
	url := fmt.Sprintf("%s/internal/users/%s", c.cfg.USER_BASE_URL, userID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	log.Info("Fetching user details from user service", req)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		log.Error("Failed to fetch user details", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Error("Failed to fetch user details, status:", resp.StatusCode)
		return nil, fmt.Errorf("failed to fetch user details, status: %d", resp.StatusCode)
	}

	var user userClientPort.User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		log.Error("Failed to decode user details response", err)
		return nil, err
	}

	log.Info("Successfully fetched user details from user service", user)

	return &user, nil
}
