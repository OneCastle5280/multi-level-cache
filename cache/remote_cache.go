package cache

import "context"

type (
	RemoteCache[T any] struct {
		//
		//  CommonCache[T]
		//  @Description: 公共处理逻辑
		//
		CommonCache[T]

		//
		//  cache
		//  @Description: 远程缓存 cache
		//
		cache Cache
	}
)

// NewRemoteCache
//
//	@Description: 创建远程缓存
//	@param mode
//	@return cache.RemoteCache[T]
func NewRemoteCache[T any](loader Loader, config *Config) *RemoteCache[T] {
	return &RemoteCache[T]{
		cache:       config.getRemoteCache(),
		CommonCache: NewCommonCache[T](loader, config.GetRemoteExpire(), config),
	}
}

// BatchGet
//
//	@Description: 批量获取缓存信息
//	@receiver r
//	@param ctx
//	@param keys
//	@return map[string][]byte
//	@return error
func (r *RemoteCache[T]) BatchGet(ctx context.Context, keys []string) (map[string][]byte, error) {
	return r.batchGet(ctx, r.cache, keys)
}

// BatchDel
//
//	@Description: 批量删除 key 缓存
//	@receiver r
//	@param ctx
//	@param keys
//	@return error
func (r *RemoteCache[T]) BatchDel(ctx context.Context, keys []string) error {
	if len(keys) == 0 {
		return nil
	}
	return r.cache.BatchDel(ctx, keys)
}
