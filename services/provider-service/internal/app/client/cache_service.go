package client

import "context"

type CacheServicePort interface {
	SetCache(ctx context.Context, key string, value interface{}) error
	GetCache(ctx context.Context, key string) (interface{}, error)
}
