package mlc

import (
	"context"
	. "mlc/cache"
)

const (
	EMPTY = ""
)

type (
	// DefaultMultiLevelCache 多级缓存默认实现类
	DefaultMultiLevelCache[T any] struct {
		config      *Config         // 缓存配置
		remoteCache *RemoteCache[T] // remote 缓存
		localCache  *LocalCache[T]  // local 缓存
		getFromDb   Loader          // 回源 db loader
		coder       Serialization   // 序列化组件
		unionKey    string          // 缓存唯一标识(全局唯一）
	}
)

func (d DefaultMultiLevelCache[T]) Get(ctx context.Context, key string) (*T, error) {
	if key == EMPTY {
		return nil, nil
	}

	values, err := d.BatchGet(ctx, []string{key})
	if err != nil {
		return nil, err
	}

	if len(values) == 0 {
		return nil, nil
	}

	return values[key], nil
}

func (d DefaultMultiLevelCache[T]) BatchGet(ctx context.Context, keys []string) (map[string]*T, error) {
	result := make(map[string]*T, len(keys))

	cacheValueMap := make(map[string][]byte, len(keys))
	if len(keys) == 0 {
		return result, nil
	}

	var err error
	// 区分模式查询数据
	switch d.config.GetMode() {
	case REMOTE:
		cacheValueMap, err = d.remoteCache.Cache.BatchGet(ctx, keys)
		break
	default:
		// 默认查询本地缓存
		cacheValueMap, err = d.localCache.Cache.BatchGet(ctx, keys)
	}

	if err != nil {
		// todo 查询 err 时异常处理
	}

	// 序列化结果
	for key, value := range cacheValueMap {
		t := new(T)
		marshalErr := d.coder.Unmarshal(value, t)
		if marshalErr != nil {
			// todo 日志打印
			continue
		}
		result[key] = t
	}
	return result, nil
}

func (d DefaultMultiLevelCache[T]) Del(ctx context.Context, key string) error {
	//TODO implement me
	panic("implement me")
}

func (d DefaultMultiLevelCache[T]) BatchDel(ctx context.Context, keys ...string) error {
	//TODO implement me
	panic("implement me")
}
