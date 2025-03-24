package cache

import (
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/zelalem-12/bill-aggregation-system_onetab/provider-service/internal/infrastructure/config"
)

func InitRedisInMemoryDB(config *config.Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.REDIS_HOST, config.REDIS_PORT),
		Password: config.REDIS_PASSWORD,
		DB:       0, // use default DB
	})
}
