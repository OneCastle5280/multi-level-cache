package cache

import (
	"context"
	"time"
)

// Cache
// @Description:
type Cache interface {
	Set(ctx context.Context, key string, value any, expire time.Duration) error         // 设置缓存
	BatchSet(ctx context.Context, values map[string][]byte, expire time.Duration) error // 批量设置缓存
	Get(ctx context.Context, key string) ([]byte, error)                                // 获取缓存
	BatchGet(ctx context.Context, keys []string) (map[string][]byte, error)             // 批量获取缓存
	Del(ctx context.Context, key string) error                                          // 删除缓存
	BatchDel(ctx context.Context, keys []string) error                                  // 批量删除缓存
}
