package client

import (
	"context"

	"github.com/go-redis/redis/v8"
	cachePort "github.com/zelalem-12/bill-aggregation-system_onetab/provider-service/internal/app/client"
)

type CacheService struct {
	redisClient *redis.Client
}

func NewCacheService(redisClient *redis.Client) cachePort.CacheServicePort {
	return &CacheService{
		redisClient: redisClient,
	}
}

func (c *CacheService) SetCache(ctx context.Context, key string, value interface{}) error {
	return c.redisClient.Set(ctx, key, value, 0).Err()
}

func (c *CacheService) GetCache(ctx context.Context, key string) (interface{}, error) {
	val, err := c.redisClient.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	return val, nil
}
