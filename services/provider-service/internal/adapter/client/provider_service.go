package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/gommon/log"
	clientPort "github.com/zelalem-12/bill-aggregation-system_onetab/provider-service/internal/app/client"
	"github.com/zelalem-12/bill-aggregation-system_onetab/provider-service/internal/domain"
	"github.com/zelalem-12/bill-aggregation-system_onetab/provider-service/internal/infrastructure/config"
	"golang.org/x/net/context"
)

type ProviderServiceClient struct {
	cfg        *config.Config
	HTTPClient *http.Client
	Redis      clientPort.CacheServicePort
}

func NewProviderServiceClient(cfg *config.Config, redisClient clientPort.CacheServicePort) clientPort.ProviderServiceClient {
	return &ProviderServiceClient{
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
		Redis: redisClient,
		cfg:   cfg,
	}
}

const (
	maxRetries     = 3                // Maximum number of retry attempts
	initialBackoff = 1 * time.Second  // Initial backoff duration
	backoffFactor  = 2                // Multiplier for backoff time
	rateLimitTTL   = 60 * time.Second // Rate limit window (1 min)
	rateLimitMax   = 10               // Max requests per provider per window
)

func (c *ProviderServiceClient) FetchBillsFromProvider(account *clientPort.LinkedAccount, provider *domain.Provider) ([]*clientPort.ProviderBillResponse, error) {
	url := provider.GetAPIURL()

	// Rate limit check
	if !c.allowRequest(provider) {
		log.Warnf("Rate limit exceeded for provider: %s", provider.GetName())
		return nil, errors.New("rate limit exceeded")
	}

	var bills []*clientPort.ProviderBillResponse
	var err error

	for attempt := 1; attempt <= maxRetries; attempt++ {
		bills, err = c.fetchBills(account, provider, url)
		if err == nil {
			log.Info("Successfully fetched bills", bills)
			return bills, nil
		}

		log.Warnf("Attempt %d/%d failed: %v", attempt, maxRetries, err)
		time.Sleep(time.Duration(attempt) * initialBackoff * backoffFactor)
	}

	return c.getFallbackBills(provider), nil
}

func (c *ProviderServiceClient) allowRequest(provider *domain.Provider) bool {
	ctx := context.Background()
	providerKey := fmt.Sprintf("rate_limit:%s", provider.GetID())

	// Get the current rate limit count
	val, err := c.Redis.GetCache(ctx, providerKey)
	if err != nil && err.Error() != "cache miss" {
		log.Error("Redis error in rate limiting:", err)
		return true
	}

	count := 0
	if val != nil {
		// Ensure type assertion safety
		if storedCount, ok := val.(int); ok {
			count = storedCount
		} else {
			log.Error("Invalid type in Redis cache for rate limiting")
			count = 0
		}
	}

	count++

	if count > rateLimitMax {
		return false
	}

	// Set updated count with TTL
	if err := c.Redis.SetCache(ctx, providerKey, count); err != nil {
		log.Error("Failed to update rate limit cache:", err)
	}

	return true
}

func (c *ProviderServiceClient) fetchBills(account *clientPort.LinkedAccount, provider *domain.Provider, url string) ([]*clientPort.ProviderBillResponse, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	switch strings.ToLower(provider.GetAuthMethod()) {
	case strings.ToLower("OAuth2"):
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", account.AuthToken))
	case strings.ToLower("APIKey"):
		req.Header.Set("X-API-Key", account.AuthToken)
	default:
		return nil, errors.New("unsupported authentication method")
	}

	log.Infof("Fetching bills from provider: %s", url)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch bills, status code: %d", resp.StatusCode)
	}

	var bills []*clientPort.ProviderBillResponse
	if err := json.NewDecoder(resp.Body).Decode(&bills); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return bills, nil
}

func (c *ProviderServiceClient) getFallbackBills(provider *domain.Provider) []*clientPort.ProviderBillResponse {
	log.Warn("Using fallback data for provider:", provider.GetName())

	providerID, err := uuid.Parse(provider.GetID())
	if err != nil {
		log.Error("Failed to parse provider ID:", err)
		return nil
	}

	return []*clientPort.ProviderBillResponse{
		{
			Amount:     100.0,
			DueDate:    time.Now(),
			Status:     "pending",
			ProviderID: providerID,
			PaidDate:   time.Now().AddDate(0, 0, 7),
		},
	}
}
