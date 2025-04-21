package cache

import (
	"context"
)

type Store interface {
	Set(ctx context.Context, key string, value interface{}, expiration int64) error
	Get(ctx context.Context, key string) (interface{}, error)
	Delete(ctx context.Context, key string) error
	SetJSON(ctx context.Context, key string, value interface{}, expiration int64) error
	GetJSON(ctx context.Context, key string, out interface{}) (bool, error)

	// Check if the key exists in the Redis set
	SRem(ctx context.Context, key string, members ...string) error

	SAdd(ctx context.Context, key string, members ...string) error

	SIsMember(ctx context.Context, setKey, member string) (bool, error)
}
