package cache

import "context"

type (
	// MultiLevelCache
	// @Description: 多级缓存对外 API
	MultiLevelCache[T any] interface {
		// Get 获取单个
		Get(ctx context.Context, key string) (*T, error)
		// BatchGet 批量获取
		BatchGet(ctx context.Context, keys []string) (map[string]*T, error)
		// Del 失效单个缓存
		Del(ctx context.Context, key string) error
		// BatchDel 批量失效缓存
		BatchDel(ctx context.Context, keys ...string) error
	}

	// Loader 回源 func
	Loader func(ctx context.Context, keys []string) (map[string][]byte, error)

	// BatchDeleteLocalCacheFunc 批量删除本地缓存
	BatchDeleteLocalCacheFunc func(ctx context.Context, keys []string) error
)
