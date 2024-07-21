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
		config        *Config         // 缓存配置
		remoteCache   *RemoteCache[T] // remote 缓存
		localCache    *LocalCache[T]  // local 缓存
		getFromDb     Loader          // 回源 db loader
		serialization Serialization   // 序列化组件
		unionKey      string          // 缓存唯一标识(全局唯一）
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
		cacheValueMap, err = d.remoteCache.BatchGet(ctx, keys)
		break
	default:
		// 默认查询本地缓存
		cacheValueMap, err = d.localCache.BatchGet(ctx, keys)
	}

	if err != nil {
		// todo 查询 err 时异常处理
	}

	// 序列化结果
	for key, value := range cacheValueMap {
		t := new(T)
		marshalErr := d.serialization.Unmarshal(value, t)
		if marshalErr != nil {
			// todo 日志打印
			continue
		}
		result[key] = t
	}
	return result, nil
}

// Del
//
//	@Description: 删除单个缓存
//	@receiver d
//	@param ctx
//	@param key
//	@return error
func (d DefaultMultiLevelCache[T]) Del(ctx context.Context, key string) error {
	if key == EMPTY {
		return nil
	}

	return d.BatchDel(ctx, key)
}

// BatchDel
//
//	@Description: 批量删除缓存
//	@receiver d
//	@param ctx
//	@param keys
//	@return error
func (d DefaultMultiLevelCache[T]) BatchDel(ctx context.Context, keys ...string) error {
	// 清理远端缓存
	err := d.remoteCache.BatchDel(ctx, keys)

	if err != nil {
		// todo 异常重试异步兜底
	}

	// todo 发送本地缓存失效
	return nil
}
