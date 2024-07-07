package mlc

import (
	"context"
	. "mlc/cache"
)

type (
	// DefaultMultiLevelCache 多级缓存默认实现类
	DefaultMultiLevelCache[T any] struct {
		config           *Config         // 缓存配置
		redisCache       *RemoteCache[T] // redis 缓存
		localCache       *LocalCache[T]  // local 缓存
		cacheInitializer *Initializer    // 缓存初始器
		getFromDb        Loader[T]       // 回源 db loader
		unionKey         string          // 缓存唯一标识(全局唯一）
	}
)

func (d DefaultMultiLevelCache[T]) Get(ctx context.Context, key string) *T {
	//TODO implement me
	panic("implement me")
}

func (d DefaultMultiLevelCache[T]) BatchGet(ctx context.Context, keys []string) []*T {
	//TODO implement me
	panic("implement me")
}

func (d DefaultMultiLevelCache[T]) Del(ctx context.Context, key string) bool {
	//TODO implement me
	panic("implement me")
}

func (d DefaultMultiLevelCache[T]) BatchDel(ctx context.Context, keys ...string) bool {
	//TODO implement me
	panic("implement me")
}
