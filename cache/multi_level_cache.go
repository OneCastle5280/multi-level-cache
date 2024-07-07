package cache

import "context"

type (
	// MultiLevelCache
	// @Description: 多级缓存对外 API
	MultiLevelCache[T any] interface {
		// Get 获取单个
		Get(ctx context.Context, key string) *T
		// BatchGet 批量获取
		BatchGet(ctx context.Context, keys []string) []*T
		// Del 失效单个缓存
		Del(ctx context.Context, key string) bool
		// BatchDel 批量失效缓存
		BatchDel(ctx context.Context, keys ...string) bool
	}

	// Loader 回源 func
	Loader func(ctx context.Context, keys []string) (map[string][]byte, error)
)
