package cache

import "context"

// Cache
// @Description:
type Cache[T any] interface {
	Get(ctx context.Context, key string) (T, error)                  // 获取缓存
	BatchGet(ctx context.Context, keys []string) map[string]T        // 批量获取缓存
	BatchSet(ctx context.Context, values map[string]T) (bool, error) // 批量设置缓存
	BatchDel(ctx context.Context, keys []string) (bool, []string)    // 批量删除缓存
}